#!/usr/bin/perl
use strict;
use warnings;
use Map qw(furthest_path);

open my $fh, '<', 'input.txt';
my $route = <$fh>;
close $fh;

print furthest_path($route) . "\n";
