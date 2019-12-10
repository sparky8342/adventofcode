#!/usr/bin/perl
use strict;
use warnings;
use Storable qw(dclone);
use Test::More tests => 5;

sub view_count {
	my ($grid,$asteroid) = @_;

	$grid = dclone($grid);

	my $width = @$grid;
	my $height = @{$grid->[0]};

	for (my $dx = -$width; $dx < $width; $dx++) {
		for (my $dy = -$height; $dy < $height; $dy++) {
			next if $dx == 0 && $dy == 0;
			my $newx = $asteroid->{x} + $dx;
			my $newy = $asteroid->{y} + $dy;
			my $found = 0;
			while ($newx >= 0 && $newx < $width && $newy >= 0 && $newy < $height) {
				if ($grid->[$newy][$newx] eq '#') {
					if ($found == 0) {
						$found = 1;
					}
					else {
						$grid->[$newy][$newx] = '.';
					}
				}
				$newx += $dx;
				$newy += $dy;
			}
		}
	}
	
	my $seen = 0;
	foreach my $row (@$grid) {
		foreach my $space (@$row) {
			if ($space eq '#') {
				$seen++;
			}
		}
	}

	return $seen - 1;
}

sub best_location {
	my ($data) = @_;
	my $grid = [map { [split(//, $_)] } split(/\n/, $data)];
	my @asteroids;
	for (my $y = 0; $y < @$grid; $y++) {
		for (my $x = 0; $x < @{$grid->[0]}; $x++) {
			if ($grid->[$y][$x] eq '#') {
				my $asteroid = { x => $x, y => $y };
				$asteroid->{view_count} = view_count($grid, $asteroid);
				push @asteroids, $asteroid;
			}
		}
	}
	@asteroids = sort { $b->{view_count} <=> $a->{view_count} } @asteroids;
	return $asteroids[0]->{view_count};
}

my @tests;
my $data = <<GRID;
.#..#
.....
#####
....#
...##
GRID
push @tests, { data => $data, result => 8 };
$data = <<GRID;
......#.#.
#..#.#....
..#######.
.#.#.###..
.#..#.....
..#....#.#
#..#....#.
.##.#..###
##...#..#.
.#....####
GRID
push @tests, { data => $data, result => 33 };
$data = <<GRID;
#.#...#.#.
.###....#.
.#....#...
##.#.#.#.#
....#.#.#.
.##..###.#
..#...##..
..##....##
......#...
.####.###.
GRID
push @tests, { data => $data, result => 35 };
$data = <<GRID;
.#..#..###
####.###.#
....###.#.
..###.##.#
##.##.#.#.
....###..#
..#.#..#.#
#..#.#.###
.##...##.#
.....#.#..
GRID
push @tests, { data => $data, result => 41 };
$data = <<GRID;
.#..##.###...#######
##.############..##.
.#.######.########.#
.###.#######.####.#.
#####.##.#.##.###.##
..#####..#.#########
####################
#.####....###.#.#.##
##.#################
#####.##.###..####..
..######..##.#######
####.##.####...##..#
.#####..#.######.###
##...#.##########...
#.##########.#######
.####.#.###.###.#.##
....##.##.###..#####
.#.#.###########.###
#.#.#.#####.####.###
###.##.####.##.#..##
GRID
push @tests, { data => $data, result => 210 };

foreach my $test (@tests) {
	is(best_location($test->{data}),$test->{result});
}

open my $fh, '<', 'input.txt';
my @data = <$fh>;
close $fh;

print best_location(join('',@data)) . "\n";
