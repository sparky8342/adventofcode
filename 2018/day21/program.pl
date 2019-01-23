#!/usr/bin/perl
use strict;
use warnings;

my $in = 12980435;

my ($a,$b,$c,$d);

$a = 0;

L1:
while (1) {

	$b = $a | 65536;
	$a = 16123384;

	L2:
	while (1) {

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
				next L1;
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
					next L2;
				}

				$c++;

			}
		}
	}	
}
