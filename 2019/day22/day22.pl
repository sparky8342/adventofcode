#!/usr/bin/perl
use strict;
use warnings;

use constant MAX_CARD => 10006;

open my $fh, '<', 'input.txt';
chomp(my @ops = <$fh>);
close $fh;

my $pos = 2019;

foreach my $op (@ops) {
	if ($op eq 'deal into new stack') {
		$pos = MAX_CARD - $pos;
	}
	elsif ($op =~ /^cut (.*)$/) {
		my $cut = $1;
		if ($cut > 0) {
			if ($pos < $cut) {
				$pos += MAX_CARD - $cut + 1;
			}
			else {
				$pos -= $cut;
			}
		}
		else {
			$cut = MAX_CARD + $cut;
			if ($pos < $cut) {
				$pos += MAX_CARD - $cut;
			}
			else {
				$pos -= $cut + 1;
			}
		}
	}
	elsif ($op =~ /^deal with increment (\d+)$/) {
		my $inc = $1;
		$pos = ($pos * $inc) % (MAX_CARD + 1);
	}
}

print "$pos\n";
