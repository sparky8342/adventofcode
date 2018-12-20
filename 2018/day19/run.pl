#!/usr/bin/perl
use strict;
use warnings;
use Device2 qw(run_program);

open my $fh, '<', 'input.txt';
my @data = <$fh>;
close $fh;
chomp($_) foreach @data;

print run_program(@data) . "\n";
