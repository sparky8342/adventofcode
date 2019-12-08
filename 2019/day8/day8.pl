#!/usr/bin/perl
use strict;
use warnings;

open my $fh, '<', 'input.txt';
chomp(my $line = <$fh>);
close $fh;

my @digits = split(//, $line);

my @grid;

my $layer_size = 25 * 6;
my $zeroes = 99999;
my $result;
while(@digits) {
	my @layer = splice(@digits, 0, $layer_size);

	my $pos = 0;
	for (my $y = 5; $y >= 0; $y--) {
		for (my $x = 0; $x < 25; $x++) {
			if (!defined($grid[$y][$x]) || $grid[$y][$x] == 2) {
				$grid[$y][$x] = $layer[$pos];
			}
			$pos++;
		}
	}

	my @digit_count;
	for my $digit (0..2) {
		$digit_count[$digit]++ foreach grep { $_ == $digit} @layer;
	}
	if ($digit_count[0] < $zeroes) {
		$zeroes = $digit_count[0];
		$result = $digit_count[1] * $digit_count[2];
	}
}
print "$result\n";

for (my $y = 5; $y >= 0; $y--) {
	for (my $x = 0; $x < 25; $x++) {
		print $grid[$y][$x] == 1 ? '#' : ' ';
	}
	print "\n";
}	
