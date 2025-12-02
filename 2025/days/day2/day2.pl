#!/usr/bin/perl
use strict;
use warnings;

my $fh;
open $fh, "../../inputs/day2.txt";
my $data = <$fh>;
close $fh;

my $part1 = 0;
my $part2 = 0;
foreach my $range (split(/,/, $data)) {
	my ($start, $end) = split(/-/, $range);
	for (my $i = $start; $i <= $end; $i++) {
		if ($i =~ /^(.+)\1$/) {
			$part1 += $i;
		}
		if ($i =~ /^(.+)\1{1,}$/) {
			$part2 += $i;
		}

	}
}
print "$part1 $part2\n";
