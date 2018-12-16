#!/usr/bin/perl
use strict;
use warnings;
use Bandits qw(combat);

open my $fh, '<', 'input.txt';
my @data = <$fh>;
close $fh;
chomp($_) foreach @data;

print combat(@data) . "\n";
