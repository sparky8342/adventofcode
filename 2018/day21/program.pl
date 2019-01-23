#!/usr/bin/perl
use strict;
use warnings;

my $in = 12980435;

my ($a,$b,$c,$d);

$a = 123;
$a = $a & 456;

if ($a == 72) {

	$a = 0;

	L3:

	$b = $a | 65536;
	$a = 16123384;

	L2:

	$c = $b & 255;
	$a = $a + $c;
	$a = $a & 16777215;
	$a = $a * 65899;
	$a = $a & 16777215;


	if ($b < 256) {
		# line 16
		if ($in == $a) {
			# end
			print "$a $b $c $d\n";
			exit;
		}
		else {
			goto L3;		
		}	
	}

	else {
		$c = 0;

		while (1) {
			$d = $c + 1;
			$d *= 256;

			print "$a $b $c $d\n";

			if ($d > $b) {
				$b = $c;
				goto L2;
			}

			$c++;
		}	
	}


}
else {
	# line 4
}




