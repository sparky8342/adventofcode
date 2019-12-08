#!/usr/bin/perl
use strict;
use warnings;

open my $fh, '<', 'input.txt';
chomp(my $line = <$fh>);
close $fh;

my @digits = split(//, $line);

my $layer_size = 25 * 6;
my $zeroes = 99999;
my $result;
while(@digits) {
	my @layer = splice(@digits, 0, $layer_size);
	my @digit_count;
	for my $digit (0..9) {
		$digit_count[$digit]++ foreach grep { $_ == $digit} @layer;
	}
	if ($digit_count[0] < $zeroes) {
		$zeroes = $digit_count[0];
		$result = $digit_count[1] * $digit_count[2];
	}
}
print "$result\n";
