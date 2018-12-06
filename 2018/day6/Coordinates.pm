package Coordinates;
use strict;
use warnings;

use parent "Exporter";
our @EXPORT_OK = qw(largest_area);

sub largest_area {
	my @data = @_;

	my @points;
	foreach my $entry (@data) {
		my ($x,$y) = split(', ',$entry);
		push @points, [$x,$y];
	}

	my ($startx,$starty,$endx,$endy);
	foreach my $point (@points) {
		my ($x,$y) = @$point;
		$startx = $x if !defined($startx) || $x < $startx;
		$starty = $y if !defined($starty) || $y < $starty;
		$endx = $x if !defined($endx) || $x > $endx;
		$endy = $y if !defined($endy) || $y > $endy;
	}

	my %area;
	my %ignore;
	foreach my $x ($startx..$endx) {
		foreach my $y ($startx..$endy) {
			my ($smallest_dist,$smallest_point);

			my %distances;

			for (my $i = 0; $i < @points; $i++) {
				my ($pointx,$pointy) = @{$points[$i]};
				my $dist = abs($x - $pointx) + abs($y - $pointy);
				$distances{$i} = $dist;
			}

			my @closest_points = sort { $distances{$a} <=> $distances{$b} } keys %distances;

			if ($distances{$closest_points[0]} != $distances{$closest_points[1]}) {
				if ($x == $startx || $x == $endx || $y == $starty || $y == $endy) {
					$ignore{$closest_points[0]} = 1;
				}
				else {
					$area{$closest_points[0]}++;
				}
			}
		}
	}

	delete($area{$_}) foreach keys %ignore;

	my $l = (sort { $b <=> $a } values %area)[0];
	return $l;
}
