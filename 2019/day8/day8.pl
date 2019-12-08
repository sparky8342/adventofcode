#!/usr/bin/perl
use strict;
use warnings;

use constant WIDTH => 25;
use constant HEIGHT => 6;
use constant SIZE => HEIGHT * WIDTH;

use constant BLACK => 0;
use constant WHITE => 1;
use constant TRANSPARENT => 2;

open my $fh, '<', 'input.txt';
chomp(my $line = <$fh>);
close $fh;
my @digits = split(//, $line);

my @grid;

my $zeroes = 99999;
my $result;
while(@digits) {
	my @layer = splice(@digits, 0, SIZE);

	my @digit_count;

	my $pos = 0;
	for (my $y = HEIGHT - 1; $y >= 0; $y--) {
		for (my $x = 0; $x < WIDTH; $x++) {
			$digit_count[$layer[$pos]]++;
			if (!defined($grid[$y][$x]) || $grid[$y][$x] == TRANSPARENT) {
				$grid[$y][$x] = $layer[$pos];
			}
			$pos++;
		}
	}

	if ($digit_count[BLACK] < $zeroes) {
		$zeroes = $digit_count[BLACK];
		$result = $digit_count[WHITE] * $digit_count[TRANSPARENT];
	}
}
print "$result\n";

for (my $y = HEIGHT - 1; $y >= 0; $y--) {
	for (my $x = 0; $x < WIDTH; $x++) {
		print $grid[$y][$x] == WHITE ? '#' : ' ';
	}
	print "\n";
}	
