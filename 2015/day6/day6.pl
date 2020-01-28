#!/usr/bin/perl
use strict;
use warnings;

open my $fh, '<', 'input.txt';
my @data = <$fh>;
close $fh;

my @grid;
my @bright_grid;
for (my $y = 0; $y < 1000; $y++) {
	for (my $x = 0; $x < 1000; $x++) {
		$grid[$x][$y] = 0;
		$bright_grid[$x][$y] = 0;
	}
}

foreach my $line (@data) {
	my ($op, $x1, $y1, $x2, $y2) = $line =~ /^(turn on|turn off|toggle) (\d+),(\d+)+ through (\d+),(\d+)$/;
	for (my $x = $x1; $x <= $x2; $x++) {
		for (my $y = $y1; $y <= $y2; $y++) {
			if ($op eq 'turn off') {
				$grid[$x][$y] = 0;
				$bright_grid[$x][$y]-- if $bright_grid[$x][$y] > 0;
			}
			elsif ($op eq 'turn on') {
				$grid[$x][$y] = 1;
				$bright_grid[$x][$y]++;
			}
			elsif ($op eq 'toggle') {
				$grid[$x][$y] = $grid[$x][$y] == 1 ? 0 : 1;
				$bright_grid[$x][$y] += 2;
			}
		}
	}
}

my $lit = 0;
my $brightness = 0;
for (my $y = 0; $y < 1000; $y++) {
	for (my $x = 0; $x < 1000; $x++) {
		$lit += $grid[$x][$y];
		$brightness += $bright_grid[$x][$y];
	}
}
print "$lit\n$brightness\n";
