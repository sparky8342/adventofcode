#!/usr/bin/perl
use strict;
use warnings;

open my $fh, '<', 'input.txt';
my $line = <$fh>;
close $fh;
$line =~ /(\d+).*?(\d+)/;
my ($target_row, $target_col) = ($1,$2);

my ($row, $col) = (6, 6);
my $num = 27995004; 

while ($row != $target_row || $col != $target_col) {
	$col++;
	$row--;
	if ($row == 0) {
		$row = $col;
		$col = 1;
	}

	$num *= 252533;
	$num %= 33554393;
}
print "$num\n";
