#!/usr/bin/perl
use strict;
use warnings;
use Storable qw(dclone);
use Math::Round;
use Test::More tests => 5;
use Math::Trig;

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
	my ($grid) = @_;

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
	return $asteroids[0];
}

sub zap {
	my ($grid, $asteroid) = @_;

	my $height = @$grid;
	my $width = @{$grid->[0]};

	my %angles;
	for (my $y = 0; $y < @$grid; $y++) {
		for (my $x = 0; $x < @{$grid->[0]}; $x++) {
			next if $x == $asteroid->{x} && $y == $asteroid->{y};
			if ($grid->[$y][$x] eq '#') {
				my $target = {
					x => $x,
					y => $y,
					sq_dist => ($x - $asteroid->{x}) ** 2 + ($y - $asteroid->{y}) ** 2
				};
				my $angle = 180 - rad2deg(atan2($x - $asteroid->{x}, $y - $asteroid->{y}));
				push @{$angles{$angle}}, $target;
			}
		}
	}

	foreach my $angle (keys %angles) {
		@{$angles{$angle}} = sort { $a->{sq_dist} <=> $b->{sq_dist} } @{$angles{$angle}};
	}

	my $shots = 0;
	while (1) {
		foreach my $angle (sort { $a <=> $b } keys %angles) {
			#print "$angle\n";
			if (@{$angles{$angle}}) {
				my $asteroid = shift @{$angles{$angle}};
				
				$shots++;
				$grid->[$asteroid->{y}][$asteroid->{x}] = '.';
	
				if ($shots == 200) {
					return $asteroid->{x} * 100 + $asteroid->{y};
				}
			}
			else {
				delete($angles{$angle});
			}
		}	
	}
}

my @tests;
my $data = <<GRID;
.#..#
.....
#####
....#
...##
GRID
my $grid = [map { [split(//, $_)] } split(/\n/, $data)];
push @tests, { grid => $grid, result => 8 };
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
$grid = [map { [split(//, $_)] } split(/\n/, $data)];
push @tests, { grid => $grid, result => 33 };
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
$grid = [map { [split(//, $_)] } split(/\n/, $data)];
push @tests, { grid => $grid, result => 35 };
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
$grid = [map { [split(//, $_)] } split(/\n/, $data)];
push @tests, { grid => $grid, result => 41 };
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
$grid = [map { [split(//, $_)] } split(/\n/, $data)];
push @tests, { grid => $grid, result => 210 };

foreach my $test (@tests) {
	my $asteroid = best_location($test->{grid});
	is($asteroid->{view_count},$test->{result});
}


my $d = <<GRID;
.#....#####...#..
##...##.#####..##
##...#...#.#####.
..#.....X...###..
..#.#.....#....##
GRID

$grid = [map { [split(//, $_)] } split(/\n/, $d)];
my $a = { x => 8, y => 3 };
#print zap($grid, $a) . "\n";
#exit;

open my $fh, '<', 'input.txt';
chomp(my @data = <$fh>);
close $fh;

$grid = [map { [split(//, $_)] } @data];
my $asteroid = best_location($grid);
print $asteroid->{view_count} . "\n";

print zap($grid, $asteroid) . "\n";
