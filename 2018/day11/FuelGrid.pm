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

	for my $size ($min_size..$max_size) {
		my $x = 1;
		my $y = 1;
		my $power = 0;

		for my $dx ($x..$x + $size - 1) {
			for my $dy ($y..$y + $size - 1) {
				$power += $grid[$dx][$dy];
			}
		}
		max_check($power,$x,$y,$size);

		while ($y < 300 - $size) {
			# go right
			while ($x <= 300 - $size) {
				for (my $dy = $y; $dy < $y + $size; $dy++) {
					$power -= $grid[$x][$dy];
				}
				$x++;
				for (my $dy = $y; $dy < $y + $size; $dy++) {
					$power += $grid[$x+$size-1][$dy];
				}
				max_check($power,$x,$y,$size);
			}

			# down 1
			for (my $dx = $x; $dx < $x + $size; $dx++) {
				$power -= $grid[$dx][$y]
			}
			$y++;
			for (my $dx = $x; $dx < $x + $size; $dx++) {
				$power += $grid[$dx][$y+$size-1]
			}

			max_check($power,$x,$y,$size);

			# go left
			while ($x > 1) {
				for (my $dy = $y; $dy < $y + $size; $dy++) {
					$power -= $grid[$x+$size-1][$dy];
				}
				$x--;
				for (my $dy = $y; $dy < $y + $size; $dy++) {
					$power += $grid[$x][$dy];
				}
				max_check($power,$x,$y,$size);
			}

			# down 1
			for (my $dx = $x; $dx < $x + $size; $dx++) {
				$power -= $grid[$dx][$y]
			}
			$y++;
			for (my $dx = $x; $dx < $x + $size; $dx++) {
				$power += $grid[$dx][$y+$size-1]
			}
			max_check($power,$x,$y,$size);
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
