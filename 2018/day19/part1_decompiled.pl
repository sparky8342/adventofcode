#!/usr/bin/perl
use strict;
use warnings;

my $a = 0;
my $c = 976;

my $b = 1;
my $f = 1;

while (1) {
	if ($b * $f == $c) {
		$a = $a + $b;
	}

	$f++;

	if ($f > $c) {
		$b++;
		if ($b > $c) {
			last;
		}
		else {
			$f = 1;
		}
	}
}

print $a . "\n";
