#!/usr/bin/perl
use strict;
use warnings;

my $grid = [];
open my $fh, '<', 'input.txt';
while(my $line = <$fh>) {
	chomp($line);
	push @$grid, [split//,$line];
}
close $fh;

my %seen = ();
while (1) {
	my $flat = join ('', map { join('',@$_) } @$grid);
	if (exists($seen{$flat})) {
		print evaluate($grid) . "\n";
		last;
	}
	$seen{$flat} = 1;
	$grid = iterate($grid);
}

sub evaluate {
	my ($grid) = @_;
	my $total = 0;
	my $s = 1;
	for my $y (0..4) {
		for my $x (0..4) {
			if ($grid->[$y][$x] eq '#') {
				$total += $s;
			}
			$s *= 2;
		}
	}
	return $total;
}

sub iterate {
	my ($grid) = @_;
	my $new_grid = [];
	for my $y (0..4) {
		for my $x (0..4) {
			my $n = 0;
			if ($y > 0 && $grid->[$y - 1][$x] eq '#') { $n++ }
			if ($y < 4 && $grid->[$y + 1][$x] eq '#') { $n++ }
			if ($x > 0 && $grid->[$y][$x - 1] eq '#') { $n++ }
			if ($x < 4 && $grid->[$y][$x + 1] eq '#') { $n++ }
	
			if ($grid->[$y][$x] eq '#') {
				if ($n != 1) {
					$new_grid->[$y][$x] = '.';
				}
				else {
					$new_grid->[$y][$x] = '#';
				}
			}
			elsif ($grid->[$y][$x] eq '.') {
				if ($n == 1 || $n == 2) {
					$new_grid->[$y][$x] = '#';
				}
				else {
					$new_grid->[$y][$x] = '.';
				}
			}
		}
	}
	return $new_grid;
}
