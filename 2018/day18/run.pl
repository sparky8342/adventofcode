#!/usr/bin/perl
use strict;
use warnings;
use Acres qw(iterate);

open my $fh, '<', 'input.txt';
my @data = <$fh>;
close $fh;
chomp($_) foreach @data;

print iterate(10,@data) . "\n";
print iterate(1000000000,@data) . "\n";
