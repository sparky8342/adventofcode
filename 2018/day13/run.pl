#!/usr/bin/perl
use strict;
use warnings;
use MineCarts qw(move);

open my $fh, '<', 'input.txt';
my @data = <$fh>;
close $fh;
chomp($_) foreach @data;

print move(1,@data) . "\n";
print move(0,@data) . "\n";
