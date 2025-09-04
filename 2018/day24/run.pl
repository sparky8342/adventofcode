#!/usr/bin/perl
use strict;
use warnings;
use ImmuneSystem qw(battle);

open my $fh, '<', 'input.txt';
my @data = <$fh>;
close $fh;
chomp($_) foreach @data;

my $winning_units = battle(@data);
print "$winning_units\n";
