package Fabric;
use strict;
use warnings;

use parent "Exporter";
our @EXPORT_OK = qw(overlap);

use constant SIZE => 999;

sub overlap {
	my @claims = @_;

	my @grid;
	foreach my $y (0..&SIZE) {
		foreach my $x (0..&SIZE) {
			$grid[$y][$x] = '.';
		}
	}

	my %claims_found;

	foreach my $claim (@claims) {
		$claim =~ /^#(\d+) @ (\d+),(\d+): (\d+)x(\d+)$/;
		my ($id,$x,$y,$w,$h) = ($1,$2,$3,$4,$5);

		$claims_found{$id} = 0;

		for (my $ypos = $y; $ypos < $y + $h; $ypos++) {
			for (my $xpos = $x; $xpos < $x + $w; $xpos++) {	
				if ($grid[$ypos][$xpos] ne '.') {
					# mark overlapping claims	
					$claims_found{$grid[$ypos][$xpos]} = 1;
					$claims_found{$id} = 1;
					$grid[$ypos][$xpos] = 'X';
				}
				else {
					$grid[$ypos][$xpos] = $id;
				}
			}
		}
	}

	my $overlap = 0;

	foreach my $y (0..&SIZE) {
		foreach my $x (0..&SIZE) {
			if ($grid[$y][$x] eq 'X') {
				$overlap++;
			}
		}
	}

	my $non_overlap_id;
	foreach my $id (keys %claims_found) {
		if ($claims_found{$id} == 0) {
			$non_overlap_id = $id;
			last;
		}
	}

	return ($overlap,$non_overlap_id);
}
