#!/usr/bin/perl
use strict;
use warnings;
use Bandits qw(combat find_winning_attack_power);

open my $fh, '<', 'input.txt';
my @data = <$fh>;
close $fh;
chomp($_) foreach @data;

print combat(3,0,@data) . "\n";
print find_winning_attack_power(@data) . "\n";
