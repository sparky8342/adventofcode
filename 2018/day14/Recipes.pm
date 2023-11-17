package Recipes;
use strict;
use warnings;

use parent "Exporter";
our @EXPORT_OK = qw(make find);

sub make {
	my ($n) = @_;

	my @r = (3,7);
	my $e1 = 0;
	my $e2 = 1;

	for (1..$n + 9) {
		my $new = $r[$e1] + $r[$e2];
		if ($new > 9) {
			push @r, 1, $new % 10;
		}
		else {
			push @r, $new;
		}
		$e1 = ($e1 + $r[$e1] + 1) % @r;
		$e2 = ($e2 + $r[$e2] + 1) % @r;
	}

	return join('',@r[$n..$n+9]);
}

sub find {
	my ($n) = @_;
	my $len = length($n);

	my @r = (3,7);
	my $e1 = 0;
	my $e2 = 1;

	while (1) {
		my $new = $r[$e1] + $r[$e2];
		if ($new > 9) {
			push @r, 1, $new % 10;
		}
		else {
			push @r, $new;
		}
		$e1 = ($e1 + $r[$e1] + 1) % @r;
		$e2 = ($e2 + $r[$e2] + 1) % @r;
		if (join('',@r[@r-$len..@r-1]) eq $n) {
			return @r-$len;
		}
		if (join('',@r[@r-$len-1..@r-2]) eq $n) {
			return @r-$len-1;
		}
	}
}
