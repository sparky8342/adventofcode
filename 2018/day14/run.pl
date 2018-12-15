#!/usr/bin/perl
use strict;
use warnings;
use Recipes qw(make);

open my $fh, '<', 'input.txt';
chomp(my $amount = <$fh>);
close $fh;

print make($amount) . "\n";
