#!/usr/bin/perl
use strict;
use warnings;

use Algorithm::Combinatorics qw(permutations);

open my $fh, '<', 'input.txt';
my @data = <$fh>;
close $fh;

my %distances;

foreach my $line (@data) {
	my ($city1, $city2, $distance) = $line =~ /^(.*?) to (.*?) = (\d+)/;
	$distances{$city1}{$city2} = $distance;
	$distances{$city2}{$city1} = $distance;
}

my @cities = keys %distances;

my $shortest = 999999;
my $longest = 0;
my $iter = permutations(\@cities);
while (my $p = $iter->next) {
	my $distance = 0;
	for (my $i = 0; $i < @$p - 1; $i++) {
		$distance += $distances{$p->[$i]}{$p->[$i+1]};
	}
	if ($distance < $shortest) {
		$shortest = $distance;
	}
	if ($distance > $longest) {
		$longest = $distance;
	}
}

print "$shortest\n$longest\n";
