package Cave;
use strict;
use warnings;

use parent "Exporter";
our @EXPORT_OK = qw(risk_level);

sub risk_level {
	my @data = @_;
	my $depth = (split/ /,$data[0])[1];
	my ($target_x,$target_y) = split/,/,(split/ /,$data[1])[1];

	my $geo_grid = [];
	my $ero_grid = [];
	my $risk = 0;
	$geo_grid->[0][0] = 0;
	$ero_grid->[0][0] = $depth % 20183;
	$risk += $ero_grid->[0][0] % 3;
	$geo_grid->[$target_x][$target_y] = 0;
	$ero_grid->[$target_x][$target_y] = $depth % 20183;
	$risk += $ero_grid->[$target_x][$target_y] % 3;

	for my $x (1..$target_x) {
		$geo_grid->[$x][0] = $x * 16807;
		$ero_grid->[$x][0] = ($geo_grid->[$x][0] + $depth) % 20183;
		$risk += $ero_grid->[$x][0] % 3;
	}

	for my $y (1..$target_y) {
		$geo_grid->[0][$y] = $y * 48271;
		$ero_grid->[0][$y] = ($geo_grid->[0][$y] + $depth) % 20183;
		$risk += $ero_grid->[0][$y] % 3;
	}	

	for my $x (1..$target_x) {
		for my $y (1..$target_y) {
			last if $x == $target_x && $y == $target_y;
			$geo_grid->[$x][$y] = $ero_grid->[$x-1][$y] * $ero_grid->[$x][$y-1]; 
			$ero_grid->[$x][$y] = ($geo_grid->[$x][$y] + $depth) % 20183;
			$risk += $ero_grid->[$x][$y] % 3;
		}
	}

	return $risk;
}
