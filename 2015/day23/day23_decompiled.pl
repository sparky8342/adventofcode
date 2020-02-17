#!/usr/bin/perl
use strict;
use warnings;

# decompiled by hand, the code is using the sequence defined by the Collatz conjecture

my @values = (26623,31911);
for my $a (@values) {
	my $b = 0;

	while ($a != 1) {
		$b++;
		if ($a % 2 == 0) {
			$a /= 2;
		}
		else {
			$a = $a * 3 + 1;
		}
	}
	print "$b\n";
}
