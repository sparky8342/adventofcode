#!/usr/bin/perl
use strict;
use warnings;

use constant SIZE => 100;
use constant ROUNDS => 100;

open my $fh, '<', 'input.txt';
chomp(my @data = <$fh>);
close $fh;

my @source_grid = ();
foreach my $line (@data) {
	push @source_grid, [split//, $line];
}

my @grid = @source_grid;
my $lights = iterate();
print "$lights\n";

@grid = @source_grid;
$lights = iterate(1);
print "$lights\n";

sub iterate {
	my ($corners_on) = @_;
	if ($corners_on) {
		$grid[0][0] = '#';
		$grid[0][SIZE - 1] = '#';
		$grid[SIZE - 1][0] = '#';
		$grid[SIZE - 1][SIZE - 1] = '#';
	}
	my $lights = 0;
	for (my $i = 1; $i <= ROUNDS; $i++) {
		my @new_grid;
		for (my $y = 0; $y < SIZE; $y++) {
			for (my $x = 0; $x < SIZE; $x++) {
				my $n = neighbours($x, $y);
				if ($corners_on && (
					($x == 0 && $y == 0)
					||
					($x == 0 && $y == SIZE - 1)
					||
					($x == SIZE - 1 && $y == 0)
					||
					($x == SIZE - 1 && $y == SIZE - 1)
				)) {
					$new_grid[$y][$x] = $grid[$y][$x];
				}
				elsif ($grid[$y][$x] eq '#' && ($n != 2 && $n != 3)) {
					$new_grid[$y][$x] = '.'
				}
				elsif ($grid[$y][$x] eq '.' && $n == 3) {
					$new_grid[$y][$x] = '#';
				}
				else {
					$new_grid[$y][$x] = $grid[$y][$x];
				}
				if ($i == ROUNDS && $new_grid[$y][$x] eq '#') {
					$lights++;
				}
			}
		}
		@grid = @new_grid;
	}
	return $lights;
}

sub neighbours {
	my ($x, $y) = @_;
	my $n = 0;

	for (my $dx = -1; $dx <= 1; $dx++) {
		for (my $dy = -1; $dy <= 1; $dy++) {
			next if $dx == 0 && $dy == 0;
			my $nx = $x + $dx;
			my $ny = $y + $dy;
			if ($nx >= 0 && $nx < SIZE && $ny >= 0 && $ny < SIZE) {
				if ($grid[$ny][$nx] eq '#') {
					$n++;
				}
			}
		}
	}
	return $n;
}
