#!/usr/bin/perl
use strict;
use warnings;
use Math::BigInt;
use Parallel::ForkManager;
use DBI;
use DBD::SQLite;
use Data::Dumper;

use constant DBFILE => 'day23.db';
use constant EXTRA_DBFILES => ['day23.db-wal', 'day23.db-shm'];

sub db_connect {
	my $dbh = DBI->connect("dbi:SQLite:dbname=" . DBFILE,"","");
	$dbh->do("PRAGMA journal_mode=WAL");
	return $dbh;
}

my $dbh = db_connect();
$dbh->do("DROP TABLE IF EXISTS stream");
$dbh->do("CREATE TABLE stream (id INTEGER PRIMARY KEY, program_id INTEGER, value INTEGER)");

my $qh = $dbh->prepare("INSERT INTO stream (program_id, value) values (?, ?)");
for (0..49) {
	$qh->execute($_, $_);
}
undef($dbh);

my $pos = 0;
my $rb = 0;
my @output;
my $id;

sub run_program {
	my ($program) = @_;

	my $ins = $program->[$pos];

	my $opcode = sprintf("%d", substr($ins, length($ins) - 2, 2));

	if ($opcode == 99) {
		return -9;
	}

	my ($mode3, $mode2, $mode1) = split(//, sprintf("%03d", substr($ins, 0, length($ins) - 2) || "000"));

	my ($a1, $a2, $a3) = @$program[$pos+1..$pos+3];
	if ($opcode != 3) {
		if ($mode1 == 0) {
			$a1 = $program->[$a1] || 0;
		}
		elsif ($mode1 == 2) {
			$a1 = $program->[$a1 + $rb] || 0;
		}
	}
	if ($mode2 == 0) {
		$a2 = $program->[$a2] || 0;
	}
	elsif ($mode2 == 2) {
		$a2 = $program->[$a2 + $rb] || 0;
	}
	if ($mode3 == 2) {
		$a3 += $rb;
	}

	if ($opcode == 1) {
		$program->[$a3] = $a1 + $a2;
		$pos += 4;
	}
	elsif ($opcode == 2) {
		$program->[$a3] = $a1 * $a2;
		$pos += 4;
	}
	elsif ($opcode == 3) {
		$a1 += $rb if $mode1 == 2;
		my $qh = $dbh->prepare("SELECT id, value FROM stream where program_id = $id ORDER BY id LIMIT 1");
		$qh->execute();
		my $row = $qh->fetchrow_hashref();
		if ($row->{id}) {
			$dbh->do("DELETE FROM stream where id = " . $row->{id});
			# message to end program
			if ($row->{value} == -1) {
				return -9;
			}
			$program->[$a1] = $row->{value};
		}
		else {
			$program->[$a1] = -1;
			sleep 1;
		}
		$pos += 2;
	}
	elsif ($opcode == 4) {
		$pos += 2;
		push @output, $a1;
		if (scalar(@output) == 3) {
			my $qh = $dbh->prepare("INSERT INTO stream (program_id, value) values (?, ?)");
			$qh->execute($output[0], $output[1]);
			$qh->execute($output[0], $output[2]);
			@output = ();
		}
	}
	elsif ($opcode == 5) {
		$pos = $a1 != 0 ? $a2 : $pos + 3;
	}
	elsif ($opcode == 6) {
		$pos = $a1 == 0 ? $a2 : $pos + 3;
	}
	elsif ($opcode == 7) {
		$program->[$a3] = $a1 < $a2 ? 1 : 0;
		$pos += 4;
	}
	elsif ($opcode == 8) {
		$program->[$a3] = $a1 == $a2 ? 1 : 0;
		$pos += 4;
	}
	elsif ($opcode == 9) {
		$rb += $a1;
		$pos += 2;
	}
	
	return undef;
}

open my $fh, '<', 'input.txt';
chomp(my $line = <$fh>);
close $fh;
my @program = map { Math::BigInt->new($_) } split(/,/,$line);

my $pm = Parallel::ForkManager->new(50);

for (my $i = 0; $i < 50; $i++) {
	$pm->start and next;

	$dbh = db_connect();

	$id = $i;
	while (1) {
		my $result = run_program(\@program);
		last if $result && $result == -9;
	}

	$pm->finish();
}

$dbh = db_connect();
my $first = 1;
my $last_y = 0;
while (1) {
	my $qh = $dbh->prepare("SELECT count(*) C FROM stream WHERE program_id != 255");
	$qh->execute();
	my $row = $qh->fetchrow_hashref();
	if ($row->{C} == 0) {
		my $qh = $dbh->prepare("SELECT value FROM stream WHERE program_id = 255 ORDER BY id desc LIMIT 2");
		$qh->execute();
		my @vals;
		while (my $row = $qh->fetchrow_hashref()) {
			unshift @vals, $row->{value};
		}
		$qh = $dbh->prepare("INSERT INTO stream (program_id, value) values (?, ?)");
		$qh->execute(0, $vals[0]);
		$qh->execute(0, $vals[1]);
		if ($first == 1) {
			# part 1
			print $vals[1] . "\n";
			$first = 0;
		}
		if ($last_y == $vals[1]) {
			# part 2
			print $last_y . "\n";
			last;
		}
		$last_y = $vals[1];
	}
	sleep 1;
}

# send message to end
$qh = $dbh->prepare("INSERT INTO stream (program_id, value) values (?, ?)");
for my $id (0..49) {
	$qh->execute($id, -1);
}

$pm->wait_all_children();
unlink($_) foreach (DBFILE, @{&EXTRA_DBFILES});
