#!/usr/bin/perl
use strict;
use warnings;
use Memory qw(sum_metadata);

open my $fh, '<', 'input.txt';
chomp(my $data = <$fh>);
close $fh;

print sum_metadata($data) . "\n";
