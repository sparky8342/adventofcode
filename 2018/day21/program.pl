#!/usr/bin/perl
use strict;
use warnings;

# set a value that won't cause the program to end normally
my $in = 1;

my ($x,$y,$c,$d);

$x = 0;

# used to keep track of possible answers
my %possible = ();

# count of instructions executed (roughly)
my $ins_count = 0;

L1:
while (1) {

	$y = $x | 65536;
	$x = 16123384;

	$ins_count += 2;

	L2:
	while (1) {

		$c = $y & 255;
		$x = $x + $c;
		$x = $x & 16777215;
		$x = $x * 65899;
		$x = $x & 16777215;

		$ins_count += 5;

		if ($y < 256) {

			# if we see a possible answer repeat, stop and find the answers
			# that took the fewest and most instructions to reach
			if (exists($possible{$x})) {
				my %h;
				foreach my $k (keys %possible) {
					$h{$possible{$k}} = $k;
				}

				my @sorted = sort { $b <=> $a } keys %h;
				print "part1: " . $h{$sorted[-1]} . "\n";
				print "part2: " . $h{$sorted[0]} . "\n";
				exit;
			}

			# record how many instructions to reach this answer
			$possible{$x} = $ins_count;

			if ($in == $x) {
				print "$x $y $c $d\n";
				exit;
			}
			else {
				$ins_count += 2;
				next L1;
			}	
		}
		else {
			$c = 0;

			$ins_count++;

			while (1) {
				$d = $c + 1;
				$d *= 256;

				$ins_count += 4;

				#print "$x $y $c $d\n";

				if ($d > $y) {
					$y = $c;
					next L2;
				}

				$c++;

			}
		}
	}	
}
