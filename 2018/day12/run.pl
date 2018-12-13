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
print iterate($initial,@data) . "\n";
