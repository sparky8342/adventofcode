#!/usr/bin/perl
use strict;
use warnings;
use Cave qw(risk_level);

open my $fh, '<', 'input.txt';
my @data = <$fh>;
close $fh;
chomp($_) foreach @data;

print risk_level(@data) . "\n";
