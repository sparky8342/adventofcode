#!/usr/bin/perl
use strict;
use warnings;
use Fabric qw(overlap);

open my $fh, '<', 'input.txt';
my @claims = <$fh>;
close $fh;
chomp($_) foreach @claims;

my ($overlap_size,$non_overlap_id) = overlap(@claims);
print "$overlap_size\n$non_overlap_id\n";
