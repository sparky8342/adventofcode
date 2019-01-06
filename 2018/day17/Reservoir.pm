package Reservoir;
use strict;
use warnings;

use parent "Exporter";
our @EXPORT_OK = qw(run_water);

my ($min_x,$min_y,$max_x,$max_y);

sub run_water {
	my @data = @_;

	my @grid;
	$grid[0][500] = '+';
	foreach my $line (@data) {
		if ($line =~ /x=(\d+), y=(\d+)\.\.(\d+)/) {
			my ($x,$y1,$y2) = ($1,$2,$3);
			$min_x = $x if !defined($min_x) || $x < $min_x;
			$max_x = $x if !defined($max_x) || $x > $max_x;
			$min_y = $y1 if !defined($min_y) || $y1 < $min_y;
			$max_y = $y2 if !defined($max_y) || $y2 > $max_y;
			for my $y ($y1..$y2) {
				$grid[$y][$x] = '#';
			}
		}
		elsif ($line =~ /y=(\d+), x=(\d+)\.\.(\d+)/) {
			my ($y,$x1,$x2) = ($1,$2,$3);
			$min_x = $x1 if !defined($min_x) || $x1 < $min_x;
			$max_x = $x2 if !defined($max_x) || $x2 > $max_x;
			$max_y = $y if !defined($max_y) || $y > $max_y;
			for my $x ($x1..$x2) {
				$grid[$y][$x] = '#';
			}
		}
	}

	for my $y (0..$max_y+1) {
		for my $x ($min_x-1..$max_x+1) {
			if (!defined($grid[$y][$x])) {
				$grid[$y][$x] = '.';
			}
		}
	}

	pp($min_x,$max_x,$min_y,$max_y,\@grid);

	$grid[1][500] = '|';
	my @running_water = ([1,500]);

	while (@running_water) {
		my $change = 0;

		# analyze each space of running water to see
		# if more should be added, or it should change
		# to standing water

		my $space = pop(@running_water);
		my ($y,$x) = @$space;
		next if $y == $max_y;
		next if $grid[$y][$x] eq '~';

		if ($grid[$y+1][$x] eq '.') {
			$grid[$y+1][$x] = '|';
			push @running_water,[$y+1,$x];
		}

		if ($grid[$y+1][$x] =~ /^[#~]$/) {
			if ($grid[$y][$x-1] eq '.') {
				$grid[$y][$x-1] = '|';
				push @running_water,[$y,$x-1];
			}
			if ($grid[$y][$x+1] eq '.') {
				$grid[$y][$x+1] = '|';
				push @running_water,[$y,$x+1];
			}

			# check for standing water
			my $dir = 0;
			if ($grid[$y][$x-1] eq '#') {
				$dir = 1;
			}			
			elsif ($grid[$y][$x+1] eq '#') {
				$dir = -1;
			}

			if ($dir != 0) {
				my $dx = $x;
				while ($grid[$y][$dx+$dir] eq '|' && $grid[$y+1][$dx+$dir] =~ /^[#~]$/) {
					$dx += $dir;
				}
				if ($grid[$y][$dx+$dir] eq '#') {
					my ($s,$e);
					if ($dir == 1) {
						$s = $x;
						$e = $dx;
					}
					elsif ($dir == -1) {
						$s = $dx;
						$e = $x;
					}
					for (my $cx = $s; $cx <= $e; $cx++) {
						$grid[$y][$cx] = '~';
						if ($grid[$y-1][$cx] eq '|') {
							push @running_water,[$y-1,$cx];
						}
					}
				}
			}
		}
	}

	pp($min_x-1,$max_x+1,$min_y,$max_y,\@grid);

	my $standing = 0;
	my $running = 0;
	for my $y ($min_y..$max_y) {
		for my $x ($min_x-1..$max_x+1) {
			$standing++ if $grid[$y][$x] eq '~';
			$running++ if $grid[$y][$x] eq '|';
		}
	}

	return ($standing + $running, $standing);
}

sub pp {
	my ($min_x,$max_x,$min_y,$max_y,$grid) = @_;

	for my $y (0..$max_y) {
		for my $x ($min_x..$max_x) {
			print $grid->[$y][$x];
		}
		print "\n";
	}
	print "\n";
}	

1;
