package FuelGrid;
use strict;
use warnings;

use parent "Exporter";
our @EXPORT_OK = qw(find_best_square);

sub find_best_square {
	my ($serial) = @_;
 
	my $max_power = 0;
	my $max_x;
	my $max_y;

	for my $x (1..298) {
		for my $y (1..298) {
			my $power = 0;
			for my $dx (0..2) {
				my $rack_id = ($x + $dx) + 10;
				for my $dy (0..2) {
					my $pl = $rack_id * ($y + $dy);
					$pl += $serial;
					$pl *= $rack_id;
					$pl = int($pl/100);
					$pl = $pl % 10;	
					$pl -= 5;
					$power += $pl;
				}
			}
			if ($power > $max_power) {
				print "$max_power\n";
				$max_power = $power;
				$max_x = $x;
				$max_y = $y;
			}
		}
	}

	return "$max_x,$max_y";
}

