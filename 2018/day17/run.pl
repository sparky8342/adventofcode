#!/usr/bin/perl
use strict;
use warnings;
use Reservoir qw(run_water);

open my $fh, '<', 'input.txt';
my @data = <$fh>;
close $fh;
chomp($_) foreach @data;

my ($total,$standing) = run_water(@data);
print "$total\n$standing\n";
