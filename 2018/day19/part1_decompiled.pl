#!/usr/bin/perl
use strict;
use warnings;

my @r = (0,0,0,0,0,0);

$r[2] = 976;

$r[1] = 1;
$r[5] = 1;

while (1) {
	$r[3] = $r[1] * $r[5];

	if ($r[3] == $r[2]) {
		$r[0] = $r[0] + $r[1];
	}

	$r[5]++;

	if ($r[5] > $r[2]) {
		$r[1]++;
		if ($r[1] > $r[2]) {
			last;
		}
		else {
			$r[5] = 1;
		}
	}
}

print $r[0] . "\n";
