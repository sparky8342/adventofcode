#!/usr/bin/perl
use strict;
use warnings;

my $grid = [];
open my $fh, '<', 'input.txt';
while(my $line = <$fh>) {
	chomp($line);
	push @$grid, [split//,$line];
}
close $fh;
my $orig_grid = [@$grid];

my %seen = ();
while (1) {
	my $flat = join ('', map { join('',@$_) } @$grid);
	if (exists($seen{$flat})) {
		print evaluate($grid) . "\n";
		last;
	}
	$seen{$flat} = 1;
	$grid = iterate($grid);
}

my $grids = [$orig_grid];
%seen = ();
for (1..200) {
	if (!is_empty($grids->[0])) {
		unshift @$grids, new_grid();
	}
	if (!is_empty($grids->[-1])) {
		push @$grids, new_grid();
	}

	$grids = iterate2($grids);
}
print evaluate2($grids) . "\n";


sub evaluate {
	my ($grid) = @_;
	my $total = 0;
	my $s = 1;
	for my $y (0..4) {
		for my $x (0..4) {
			if ($grid->[$y][$x] eq '#') {
				$total += $s;
			}
			$s *= 2;
		}
	}
	return $total;
}

sub iterate {
	my ($grid) = @_;
	my $new_grid = [];
	for my $y (0..4) {
		for my $x (0..4) {
			my $n = 0;
			if ($y > 0 && $grid->[$y - 1][$x] eq '#') { $n++ }
			if ($y < 4 && $grid->[$y + 1][$x] eq '#') { $n++ }
			if ($x > 0 && $grid->[$y][$x - 1] eq '#') { $n++ }
			if ($x < 4 && $grid->[$y][$x + 1] eq '#') { $n++ }
	
			if ($grid->[$y][$x] eq '#') {
				if ($n != 1) {
					$new_grid->[$y][$x] = '.';
				}
				else {
					$new_grid->[$y][$x] = '#';
				}
			}
			elsif ($grid->[$y][$x] eq '.') {
				if ($n == 1 || $n == 2) {
					$new_grid->[$y][$x] = '#';
				}
				else {
					$new_grid->[$y][$x] = '.';
				}
			}
		}
	}
	return $new_grid;
}

sub is_empty {
	my ($grid) = @_;

	for my $y (0..4) {
		for my $x (0..4) {
			if ($grid->[$y][$x] eq '#') {
				return 0;
			}
		}
	}
	return 1;
}

sub new_grid {
	my $grid = [];
	for my $y (0..4) {
		push @$grid, [".", ".", ".", ".", "."];
	}
	return $grid;
}

sub iterate2 {
	my ($grids) = @_;
	my $new_grids = [];

	for my $grid_no (0..@$grids-1) {
		my $grid = $grids->[$grid_no];
		my $new_grid = new_grid();
		for my $y (0..4) {
			for my $x (0..4) {
				if ($x == 2 && $y == 2) {
					next;
				}

				my $n = 0;

				if ($y > 0 && $grid->[$y - 1][$x] eq '#') { $n++ }
				if ($y < 4 && $grid->[$y + 1][$x] eq '#') { $n++ }
				if ($x > 0 && $grid->[$y][$x - 1] eq '#') { $n++ }
				if ($x < 4 && $grid->[$y][$x + 1] eq '#') { $n++ }

				if ($grid_no > 0) {
					if ($x == 0 && $grids->[$grid_no-1]->[2][1] eq '#') { $n++ }
					if ($x == 4 && $grids->[$grid_no-1]->[2][3] eq '#') { $n++ }
					if ($y == 0 && $grids->[$grid_no-1]->[1][2] eq '#') { $n++ }
					if ($y == 4 && $grids->[$grid_no-1]->[3][2] eq '#') { $n++ }
				}
				if ($grid_no < @$grids - 1) {
					if ($y == 2 && $x == 1) {
						for my $in_y (0..4) {
							if ($grids->[$grid_no+1]->[$in_y]->[0] eq '#') {
								$n++;
							}
						}
					}
					if ($y == 2 && $x == 3) {
						for my $in_y (0..4) {
							if ($grids->[$grid_no+1]->[$in_y]->[4] eq '#') {
								$n++;
							}
						}
					}
					if ($y == 1 && $x == 2) {
						for my $in_x (0..4) {
							if ($grids->[$grid_no+1]->[0]->[$in_x] eq '#') {
								$n++;
							}
						}
					}
					if ($y == 3 && $x == 2) {
						for my $in_x (0..4) {
							if ($grids->[$grid_no+1]->[4]->[$in_x] eq '#') {
								$n++;
							}
						}
					}
				}

				if ($grid->[$y][$x] eq '#') {
					if ($n != 1) {
						$new_grid->[$y][$x] = '.';
					}
					else {
						$new_grid->[$y][$x] = '#';
					}
				}
				elsif ($grid->[$y][$x] eq '.') {
					if ($n == 1 || $n == 2) {
						$new_grid->[$y][$x] = '#';
					}
					else {
						$new_grid->[$y][$x] = '.';
					}
				}
			}
		}
		push @$new_grids, $new_grid;
	}

	return $new_grids;
}

sub evaluate2 {
	my ($grids) = @_;
	my $bugs = 0;
	for my $grid (@$grids) {
		for my $y (0..4) {
			for my $x (0..4) {
				if ($grid->[$y]->[$x] eq '#') {
					$bugs++
				}
			}
		}
	}
	return $bugs;
}
