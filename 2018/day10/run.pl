#!/usr/bin/perl
use strict;
use warnings;
use Stars qw(animate);

open my $fh, '<', 'input.txt';
my @records = <$fh>;
close $fh;
chomp($_) foreach @records;

animate(20000,@records);
