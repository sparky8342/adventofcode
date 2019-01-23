#!/usr/bin/perl
use strict;
use warnings;

my @r = (0,0,0,0,0,0);
$r[0] = 12980435;

$r[2] = 123;
$r[2] = $r[2] & 456;

if ($r[2] == 72) {

	$r[2] = 0;

	L3:

	$r[5] = $r[2] | 65536;
	$r[2] = 16123384;

	L2:

	$r[3] = $r[5] & 255;
	$r[2] = $r[2] + $r[3];
	$r[2] = $r[2] & 16777215;
	$r[2] = $r[2] * 65899;
	$r[2] = $r[2] & 16777215;


	if ($r[5] < 256) {
		# line 16
		if ($r[0] == $r[2]) {
			# end
			print join(',',@r) . "\n";
			exit;
		}
		else {
			goto L3;		
		}	
	}

	else {
		$r[3] = 0;

		L1:

		$r[1] = $r[3] + 1;
		$r[1] *= 256;

		print join(',',@r) . "\n";

		if ($r[1] > $r[5]) {
			$r[5] = $r[3];
			goto L2;
		}
		else {
			$r[3]++;
			goto L1;
		}	
	}


}
else {
	# line 4
}




