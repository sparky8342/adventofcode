#!/usr/bin/perl
use strict;
use warnings;
use Math::BigInt;
use Term::ANSIScreen qw(cls locate);

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
			$program->[$a1] = $input;
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
my @source_program = split(/,/,$line);

# part 1
my @program = map { Math::BigInt->new($_) } @source_program;

my %grid = ();

my ($x, $y) = (0, 0);
my ($hx, $hy) = (0, 0);

while (1) {
	my $output = run_program(\@program);
	if ($output == -9) {
		$hy = $y - 2;
		last;
	}
	if ($output == 10) {
		$hx = $x - 1 unless $hx;
		$y++;
		$x = 0;
	}
	else {
		my $char = chr($output);
		$grid{$x}{$y} = $char;
		$x++;
	}
}
print_grid();

my $sum = 0;
for my $y (0..$hy) {
	for my $x (0..$hx) {
		if ($grid{$x}{$y} eq '#') {
			if (($grid{$x-1}{$y} || '') eq '#'
			&&  ($grid{$x+1}{$y} || '') eq '#'
			&&  ($grid{$x}{$y-1} || '') eq '#'
			&&  ($grid{$x}{$y+1} || '') eq '#') {
				$sum += $x * $y;
			}
		}
	}
}
print "$sum\n";

sub print_grid {
	#locate(1,1);
	for my $y (0..$hy) {
		for my $x (0..$hx) {
			print $grid{$x}{$y} || ' ';
		}
		print "\n";
	}
}
