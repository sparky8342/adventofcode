#!/usr/bin/perl
use strict;
use warnings;
use Sleigh qw(steps);

open my $fh, '<', 'input.txt';
my @ins = <$fh>;
close $fh;
chomp($_) foreach @ins;

print steps(@ins) . "\n";
