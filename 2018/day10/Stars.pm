package Stars;
use strict;
use warnings;

use parent "Exporter";
our @EXPORT_OK = qw(animate);

sub animate {
	my ($iterations,@data) = @_;

	my @points;
	foreach my $row (@data) {
		$row =~ /((?:-)?\d+).*?((?:-)?\d+).*?((?:-)?\d+).*?((?:-)?\d+)/;
		my ($x,$y,$dx,$dy) = ($1,$2,$3,$4);
		push @points, { x => $x, y => $y, dx => $dx, dy => $dy };
	}

	for (1..$iterations) {
		my ($startx,$endx,$starty,$endy);
		my %grid;
		foreach my $point (@points) {
			$point->{x} += $point->{dx};
			$point->{y} += $point->{dy};
			if (!defined($startx) || $point->{x} < $startx) {
				$startx = $point->{x};
			}
			if (!defined($endx) || $point->{x} > $endx) {
				$endx = $point->{x};
			}
			if (!defined($starty) || $point->{y} < $starty) {
				$starty = $point->{y};
			}
			if (!defined($endy) || $point->{y} > $endy) {
				$endy = $point->{y};
			}

			$grid{$point->{x}}{$point->{y}} = 1;
		}

		if ($endx - $startx < 65) {
			print "Creating candidate $_.stars\n";
			open my $fh, '>', "$_.stars";
			for (my $y = $starty; $y <= $endy; $y++) {
				for (my $x = $startx; $x <= $endx; $x++) {
					if (exists($grid{$x}{$y})) {
						print $fh '#';
					}
					else {
						print $fh ' ';
					}
				}
				print $fh "\n";
			}
			close $fh;
		}
	}
}
