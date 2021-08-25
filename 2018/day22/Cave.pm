package Cave;
use strict;
use warnings;

use parent "Exporter";
our @EXPORT_OK = qw(risk_level search);

use constant NOTHING => 1;
use constant TORCH => 2;
use constant CLIMBING_GEAR => 3;

use constant ADD_ON => 50;

sub risk_level {
	my @data = @_;
	my $depth = (split/ /,$data[0])[1];
	my ($target_x,$target_y) = split/,/,(split/ /,$data[1])[1];

	my $geo_grid = [];
	my $ero_grid = [];
	my $type_grid = [];
	my $risk = 0;
	$geo_grid->[0][0] = 0;
	$ero_grid->[0][0] = $depth % 20183;
	$type_grid->[0][0] = $ero_grid->[0][0] % 3;
	$risk += $type_grid->[0][0];
	$geo_grid->[$target_x][$target_y] = 0;
	$ero_grid->[$target_x][$target_y] = $depth % 20183;
	$type_grid->[$target_x][$target_y] = $ero_grid->[$target_x][$target_y] % 3;
	$risk += $type_grid->[$target_x][$target_y];

	for my $x (1..$target_x + ADD_ON) {
		$geo_grid->[$x][0] = $x * 16807;
		$ero_grid->[$x][0] = ($geo_grid->[$x][0] + $depth) % 20183;
		$type_grid->[$x][0] = $ero_grid->[$x][0] % 3;
		$risk += $type_grid->[$x][0] if $x <= $target_x;
	}

	for my $y (1..$target_y + ADD_ON) {
		$geo_grid->[0][$y] = $y * 48271;
		$ero_grid->[0][$y] = ($geo_grid->[0][$y] + $depth) % 20183;
		$type_grid->[0][$y] = $ero_grid->[0][$y] % 3;
		$risk += $type_grid->[0][$y] if $y <= $target_y;
	}	

	for my $x (1..$target_x + ADD_ON) {
		for my $y (1..$target_y + ADD_ON) {
			next if $x == $target_x && $y == $target_y;
			$geo_grid->[$x][$y] = $ero_grid->[$x-1][$y] * $ero_grid->[$x][$y-1]; 
			$ero_grid->[$x][$y] = ($geo_grid->[$x][$y] + $depth) % 20183;
			$type_grid->[$x][$y] = $ero_grid->[$x][$y] % 3;
			$risk += $type_grid->[$x][$y] if $x <= $target_x && $y <= $target_y;
		}
	}

	return ($risk, $type_grid, $target_x, $target_y);
}

sub search {
	my ($grid, $target_x, $target_y) = @_;

	my $width = scalar(@$grid);
	my $height = scalar(@{$grid->[0]});

	my @dirs = (
		{ dx =>  0, dy =>  1 },
		{ dx =>  0, dy => -1 },
		{ dx =>  1, dy =>  0 },
		{ dx => -1, dy =>  0 }
	);

	my $start = { x => 0, y => 0, equipped => TORCH, time => 0 };

	my @queue = ($start);
	my %visited;

	while (@queue > 0) {
		my $state = shift @queue;

		if ($state->{x} == $target_x && $state->{y} == $target_y) {
			return $state->{time};
		}

		if (exists($visited{$state->{x}, $state->{y}, $state->{equipped}})) {
			next;
		}
		$visited{$state->{x}, $state->{y}, $state->{equipped}} = 1;

		if ($state->{x} < 0 || $state->{x} == $width || $state->{y} < 0 || $state->{y} == $height) {
			next;
		}

		if ($state->{equipped} == NOTHING && $grid->[$state->{x}][$state->{y}] == 0) {
			next;
		}
		elsif ($state->{equipped} == TORCH && $grid->[$state->{x}][$state->{y}] == 1) {
			next;
		}
		elsif ($state->{equipped} == CLIMBING_GEAR && $grid->[$state->{x}][$state->{y}] == 2) {
			next;
		}

		foreach my $equipment (NOTHING, TORCH, CLIMBING_GEAR) {
			if ($state->{equipped} eq $equipment) {
				next;
			}
			my $new_state = {
				x => $state->{x},
				y => $state->{y},
				equipped => $equipment,
				time => $state->{time} + 7
			};
			push @queue, $new_state;
		}
		foreach my $dir (@dirs) {
			my $new_state = {
				x => $state->{x} + $dir->{dx},
				y => $state->{y} + $dir->{dy},
				equipped => $state->{equipped},
				time => $state->{time} + 1
			};
			push @queue, $new_state;
		}

		@queue = sort { $a->{time} <=> $b->{time} } @queue;
	}
}
