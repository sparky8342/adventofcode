#!/usr/bin/perl
use strict;
use warnings;
use Plants qw(iterate);

open my $fh, '<', 'input.txt';
my @data = <$fh>;
close $fh;
chomp($_) foreach @data;

my $initial = shift(@data);
($initial) = $initial =~ /: (.*)/;
shift(@data);
print iterate($initial,20,@data) . "\n";
print iterate($initial,50000000000,@data) . "\n";
