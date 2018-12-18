package Acres;
use strict;
use warnings;

use parent "Exporter";
our @EXPORT_OK = qw(iterate);

sub iterate {
	my ($amount,@data) = @_;

	my @grid;
	foreach my $line (@data) {
		push @grid, [split//,$line];
	}

	pp(\@grid);

	my $x_size = @{$grid[0]};
	my $y_size = @grid;

	my %cache = (join('', map { join('',@$_) } @grid) => 0);

	for my $iter (1..$amount) {
		my @newgrid;

		for (my $y = 0; $y < $y_size; $y++) {
			for (my $x = 0; $x < $x_size; $x++) {
				my %counts = ('|' => 0, '#' => 0);
				for (my $dx = -1; $dx <= 1; $dx++) {
					for (my $dy = -1; $dy <= 1; $dy++) {
						next if $dx == 0 && $dy == 0;
						my $ax = $x + $dx;
						next if $ax < 0 || $ax >= $x_size;
						my $ay = $y + $dy;
						next if $ay < 0 || $ay >= $y_size; 
						#print "$ay $ay\n";
						$counts{$grid[$ay][$ax]}++;
					}
				}
				my $char;
				if ($grid[$y][$x] eq '.') {
					if ($counts{'|'} >= 3) {
						$char = '|';
					}
					else {
						$char = '.';
					}
				}
				if ($grid[$y][$x] eq '|') {
					if ($counts{'#'} >= 3) {
						$char = '#';
					}
					else {
						$char = '|';
					}
				}
				if ($grid[$y][$x] eq '#') {
					if ($counts{'#'} >= 1 && $counts{'|'} >= 1) {
						$char = '#';
					}
					else {
						$char = '.';
					}
				}
				$newgrid[$y][$x] = $char;
			}
		}
		@grid = @newgrid;

		my $squash = join('', map { join('',@$_) } @grid);
		if (exists($cache{$squash})) {
			my $cycle = $iter - $cache{$squash};
			if (($amount - $iter) % $cycle == 0) {
				last;
			}
		}

		$cache{$squash} = $iter;
		pp(\@grid);
	}

	my %counts;	
	for (my $y = 0; $y < $y_size; $y++) {
		for (my $x = 0; $x < $x_size; $x++) {
			$counts{$grid[$y][$x]}++;
		}
	}
	return $counts{'|'} * $counts{'#'};
}

sub pp {
	my ($grid) = @_;
	foreach my $line (@$grid) {
		print join('',@$line) . "\n";
	}
	print "\n";
}
