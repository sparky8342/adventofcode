#!/usr/bin/perl
use strict;
use warnings;

open my $fh, '<', 'input.txt';
my @data = <$fh>;
close $fh;

# part 1
my $nice = 0;
foreach my $line (@data) {
	if (
		$line =~ /[aeiou].*[aeiou].*[aeiou]/
		&&
		$line =~ /(.)\1/
		&&
		$line !~ /(ab|cd|pq|xy)/
	) {
		$nice++;
	}
}
print "$nice\n";

# part 2
$nice = 0;
foreach my $line (@data) {
	if (
		$line =~ /(..).*\1/
		&&
		$line =~ /(.)[^\1]\1/
	) {
		$nice++;
	}
}	
print "$nice\n";
