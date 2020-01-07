#!/usr/bin/perl
use strict;
use warnings;
use Math::BigInt;
use Parallel::ForkManager;
use DBI;
use DBD::SQLite;
use Data::Dumper;

sub db_connect {
	my $dbh = DBI->connect("dbi:SQLite:dbname=day23.db","","");
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
			$program->[$a1] = $row->{value};
		}
		else {
			$program->[$a1] = -1;
		}
		$pos += 2;
	}
	elsif ($opcode == 4) {
		$pos += 2;
		push @output, $a1;
		if (scalar(@output) == 3) {
			if ($output[0] == 255) {
				return $output[2];
			}
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

# part 1

my $pm = Parallel::ForkManager->new(50);

$pm->run_on_finish(
	sub {
		my ($pid, $exit_code, $ident, $exit_signal, $core_dump, $data_structure_reference) = @_;

		if (defined($data_structure_reference)) {
			my $string = ${$data_structure_reference};
			print "$string\n";
		}
	}
);

for (my $i = 0; $i < 50; $i++) {
	$pm->start and next;

	$dbh = db_connect();

	$id = $i;
	while (1) {
		my $result = run_program(\@program);
		if ($result) {
			$pm->finish(0, \$result);
		}
	}

	$pm->finish();
}

$pm->wait_all_children();
