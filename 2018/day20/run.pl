#!/usr/bin/perl
use strict;
use warnings;
use Map qw(furthest_path find_amount);

open my $fh, '<', 'input.txt';
my $route = <$fh>;
close $fh;

print furthest_path($route) . "\n";
print find_amount($route,1000) . "\n";
