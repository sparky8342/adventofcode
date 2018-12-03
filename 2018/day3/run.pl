#!/usr/bin/perl
use strict;
use warnings;
use Fabric qw(overlap);

open my $fh, '<', 'input.txt';
my @claims = <$fh>;
close $fh;
chomp($_) foreach @claims;

print overlap(@claims) . "\n";
