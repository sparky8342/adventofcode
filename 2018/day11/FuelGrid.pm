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

	for my $size (1..300) {
		print "$size\n";
		for my $x (1..300) {
			for my $y (1..300) {
				#print "$x $y\n";
				my $power = 0;
				my $xlimit = $size;
				if (300 - $x < $xlimit) {
					$xlimit = 300 - $x;
				}
				for my $dx (0..$xlimit) {
					my $ylimit = $size;
					if (300 - $y < $ylimit) {
						$ylimit = 300 - $y;
					}
					for my $dy (0..$ylimit) {
						$power += $grid[$x+$dx][$y+$dy];
					}
				}

				if ($power > $max_power) {
					print "$power : $x,$y,$size\n";
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

