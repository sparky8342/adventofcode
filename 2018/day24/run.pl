#!/usr/bin/perl
use strict;
use warnings;
use ImmuneSystem qw(parse_data battle find_winning_boost);

open my $fh, '<', 'input.txt';
my @data = <$fh>;
close $fh;
chomp($_) foreach @data;

my ($immune, $infection) = parse_data(@data);
my ($winner, $winning_units) = battle($immune, $infection);
print "$winning_units\n";

$winning_units = find_winning_boost($immune, $infection);
print "$winning_units\n";
