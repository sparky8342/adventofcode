#!/usr/bin/perl
use strict;
use warnings;
use Nanobots qw(in_range most_in_range);

open my $fh, '<', 'input.txt';
my @data = <$fh>;
close $fh;
chomp($_) foreach @data;

print in_range(@data) . "\n";

most_in_range(@data);
