#!/usr/bin/perl
use strict;
use warnings;
use Cave qw(risk_level search);

open my $fh, '<', 'input.txt';
my @data = <$fh>;
close $fh;
chomp($_) foreach @data;

my ($rl, $grid, $target_x, $target_y) = risk_level(@data);
print "$rl\n";

my $time = search($grid, $target_x, $target_y);
print "$time\n";
