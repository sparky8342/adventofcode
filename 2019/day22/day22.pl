#!/usr/bin/perl
use strict;
use warnings;
use bigint;

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


my $offset = 0;
my $increment = 1;
my $max_card = 119315717514047;
my $loops = 101741582076661;

foreach my $op (@ops) {
	if ($op eq 'deal into new stack') {
		$increment = ($increment * -1) % $max_card;
		$offset = ($offset + $increment) % $max_card;
	} elsif ($op =~ /^cut (.*)$/) {
		$offset += $increment * $1;
	} elsif ($op =~ /^deal with increment (\d+)$/) {
		my $deal = $1;
		$increment *= modpow($deal, $max_card - 2, $max_card);
	}
}

my $inc = modpow($increment, $loops, $max_card);
$offset = $offset * (1 - $inc) * modpow((1 - $increment) % $max_card, $max_card - 2, $max_card);
$offset %= $max_card;
my $card = ($offset + 2020 * $inc) % $max_card;
print "$card\n";

sub modpow {
	my ($n, $pow, $mod) = @_;

	if ($pow == 0) {
		return 1;
	} elsif ($pow % 2 == 0) {
		return modpow(($n * $n) % $mod, $pow / 2, $mod);
	} else {
		return ($n * modpow($n, $pow - 1, $mod)) % $mod;
	}
}
