#!/usr/bin/perl
use strict;
use warnings;

open my $fh, "<input.txt";
my @data = <$fh>;
close $fh;
chomp($_) foreach @data;

my %words = ('one' => 1, 'two' => 2, 'three' => 3, 'four' => 4, 'five' => 5, 'six' => 6, 'seven' => 7, 'eight' => 8, 'nine' => 9);

my $total = 0;
my $total2 = 0;
foreach my $line (@data) {
	my ($first) = $line =~ /(\d)/;
	my ($last) = $line =~ /.*(\d)/;
	$total += $first * 10 + $last;

	($first) = $line =~ /(\d|one|two|three|four|five|six|seven|eight|nine)/;
	$first = $words{$first} if length($first) > 1;
	($last) = $line =~ /.*(\d|one|two|three|four|five|six|seven|eight|nine)/;
	$last = $words{$last} if length($last) > 1;
	$total2 += $first * 10 + $last;
}
print "$total\n$total2\n";
