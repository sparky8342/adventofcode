package Recipes;
use strict;
use warnings;

use parent "Exporter";
our @EXPORT_OK = qw(make);

sub make {
	my ($n) = @_;

	my @r = (3,7);
	my $e1 = 0;
	my $e2 = 1;

	for (1..$n + 10) {
		push @r, split//,$r[$e1] + $r[$e2];
		$e1 = ($e1 + $r[$e1] + 1) % @r;
		$e2 = ($e2 + $r[$e2] + 1) % @r;
		#print "$e1 $e2: " . join(' ',@r) . "\n";
	}

	return join('',@r[$n..$n+9]);
}
