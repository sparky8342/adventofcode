package FuelGrid;
use strict;
use warnings;

use parent "Exporter";
our @EXPORT_OK = qw(find_best_square);

$| = 1;

sub find_best_square {
	my ($serial) = @_;
 
	my $max_power = 0;
	my $max_x;
	my $max_y;
	my $max_size;

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

	for my $x (1..298) {
		for my $y (1..298) {
			print "$x $y\n";
			#my $size_limit = 300 - ($x > $y ? $x : $y);
			#foreach my $size (1..$size_limit) {
			foreach my $size (2..2) {
				my $power = 0;
				for my $dx (0..$size) {
					for my $dy (0..$size) {
						$power += $grid[$x+$dx][$y+$dy];
					}
				}

				if ($power > $max_power) {
					print "$max_power\n";
					$max_power = $power;
					$max_x = $x;
					$max_y = $y;
					$max_size = $size;
				}
			}
		}
	}

	return "$max_x,$max_y";
}

