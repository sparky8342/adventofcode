#!/usr/bin/perl
use strict;
use warnings;
use Repose qw(most_asleep);

open my $fh, '<', 'input.txt';
my @records = <$fh>;
close $fh;
chomp($_) foreach @records;

print most_asleep(@records) . "\n";
