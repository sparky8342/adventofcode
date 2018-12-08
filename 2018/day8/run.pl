#!/usr/bin/perl
use strict;
use warnings;
use Memory qw(process_data);

open my $fh, '<', 'input.txt';
chomp(my $data = <$fh>);
close $fh;

my ($sum,$value) = process_data($data);
print "$sum\n$value\n";
