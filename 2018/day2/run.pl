#!/usr/bin/perl
use strict;
use warnings;
use Inventory qw(checksum find_similar);

open my $fh, '<', 'input.txt';
my @boxes = <$fh>;
close $fh;

print checksum(@boxes) . "\n";
print find_similar(@boxes) . "\n";
