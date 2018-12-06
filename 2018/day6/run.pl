#!/usr/bin/perl
use strict;
use warnings;
use Coordinates qw(find_areas);

open my $fh, '<', 'input.txt';
my @points = <$fh>;
close $fh;
chomp($_) foreach @points;

my ($area,$safe_area) = find_areas(10000,@points);
print "$area\n$safe_area\n";
