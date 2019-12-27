#!/usr/bin/perl
use strict;
use warnings;
use Math::BigInt;

my $pos = 0;
my $rb = 0;

sub run_program {
	my ($program, $input) = @_;

	my @output;

	while (1) {
		my $ins = $program->[$pos];

                my $opcode = sprintf("%d", substr($ins, length($ins) - 2, 2));

		if ($opcode == 99) {
			last;
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
			$program->[$a1] = shift(@$input);
			$pos += 2;
		}
		elsif ($opcode == 4) {
			$pos += 2;
			return $a1;
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
	}
	return -9;
}

open my $fh, '<', 'input.txt';
chomp(my $line = <$fh>);
close $fh;
my @source_program = map { Math::BigInt->new($_) } split(/,/,$line);

# part 1
my @program;

my $in_beam = 0;
for (my $x = 0; $x < 50; $x++) {
	for (my $y = 0; $y < 50; $y++) {
		$in_beam++ if inside($x, $y);
	}
}
print "$in_beam\n";

# part 2
my $size = 99;

# skip over section with empty spaces
my $x = 10;
my $y = 0;
while (!inside($x,$y)) {
	$y++;
}

# walk down top of beam, check bottom left
# corner of possible square
while (1) {
	while (inside($x, $y)) {
		$x++;
	}
	$x--;

	my $lx = $x - $size;

	if ($lx > 0 && inside($lx, $y + $size)) {
		my $ans = $lx * 10000 + $y;
		print "$ans\n";
		last;
	}
	$y++;
}

sub inside {
	my ($x, $y) = @_;
	$pos = 0; $rb = 0;
	@program = @source_program;
	return run_program(\@program, [$x, $y]);
}
