package FuelGrid;
use strict;
use warnings;

use parent "Exporter";
our @EXPORT_OK = qw(find_best_square);

$| = 1;

my $best_power;
my $best_x;
my $best_y;
my $best_size;

sub find_best_square {
	my ($serial,$min_size,$max_size) = @_;

	$best_power = 0;

	my @grid;
	for my $x (1..300) {
		my $rack_id = $x + 10;
		for my $y (1..300) {
			my $pl = $rack_id * $y;
			$pl += $serial;
			$pl *= $rack_id;
			$pl = int($pl/100);
			$pl = $pl % 10;	
			$pl -= 5;
			$grid[$x][$y] = $pl;
		}
	}
	$grid[0][0] = 0;
	for (my $i = 1; $i <= 300; $i++) {
		$grid[$i][0] = 0;
		$grid[0][$i] = 0;
	}

	# partial sums
	for (my $i = 2; $i <= 300; $i++) {
		$grid[$i][1] += $grid[$i-1][1];
		$grid[1][$i] += $grid[1][$i-1];
	}
	for my $x (2..300) {
		for my $y (2..300) {
			$grid[$x][$y] = $grid[$x][$y] + $grid[$x-1][$y] + $grid[$x][$y-1] - $grid[$x-1][$y-1];
		}
	}

	# find best power
	for my $size ($min_size..$max_size) {
		for (my $x = $size; $x <= 300; $x++) {
			for (my $y = $size; $y <= 300; $y++) {
				my $power = $grid[$x][$y] - $grid[$x-$size][$y] - $grid[$x][$y-$size] + $grid[$x-$size][$y-$size];
				max_check($power, $x-$size + 1, $y-$size + 1, $size);
			}
		}
	}

	if ($min_size == $max_size) {
		return "$best_x,$best_y";
	}
	else {
		return "$best_x,$best_y,$best_size";
	}
}


sub max_check {
	my ($power,$x,$y,$size) = @_;
	if ($power > $best_power) {
		$best_power = $power;
		$best_x = $x;
		$best_y = $y;
		$best_size = $size;
	}
}

1;
