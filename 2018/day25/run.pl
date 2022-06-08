#!/usr/bin/perl
use strict;
use warnings;
use Points qw(count_groups);

open my $fh, '<', 'input.txt';
my @points = <$fh>;
close $fh;
chomp($_) foreach @points;

print count_groups(@points) . "\n";
