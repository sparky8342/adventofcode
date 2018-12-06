#!/usr/bin/perl
use strict;
use warnings;
use Coordinates qw(largest_area);

open my $fh, '<', 'input.txt';
my @points = <$fh>;
close $fh;
chomp($_) foreach @points;

print largest_area(@points) . "\n";
