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

	foreach my $claim (@claims) {
		# #1 @ 1,3: 4x4	

		$claim =~ /^#(\d+) @ (\d+),(\d+): (\d+)x(\d+)$/;
		my ($id,$x,$y,$w,$h) = ($1,$2,$3,$4,$5);

		for (my $ypos = $y; $ypos < $y + $h; $ypos++) {
			for (my $xpos = $x; $xpos < $x + $w; $xpos++) {	
				if ($grid[$ypos][$xpos] ne '.') {
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
		#print "\n";
	}

	return $overlap;
}
