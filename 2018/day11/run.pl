#!/usr/bin/perl
use strict;
use warnings;
use FuelGrid qw(find_best_square);

open my $fh, '<', 'input.txt';
chomp(my $serial = <$fh>);
close $fh;

print find_best_square($serial,1,300) . "\n";
