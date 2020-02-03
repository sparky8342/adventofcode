#!/usr/bin/perl
use strict;
use warnings;

use Math::Prime::Util qw(fordivisors);

open my $fh, '<', 'input.txt';
my $presents = <$fh>;
close $fh;

$presents /= 10;

my $house = 1;
while (1) {
	my $sum = 0;
	fordivisors { $sum += $_ } $house;
	if ($sum >= $presents) {
		print "$house\n";
		last;
	}
	$house++;
}

$presents = $presents * 10 / 11;
$house = 1;
while (1) {
	my $sum = 0;
	fordivisors { if ($house / $_ <= 50) { $sum += $_ } } $house;
	if ($sum >= $presents) {
		print "$house\n";
		last;
	}
	$house++;
}
