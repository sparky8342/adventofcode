#!/usr/bin/perl
use strict;
use warnings;
use Repose qw(most_asleep);

open my $fh, '<', 'input.txt';
my @records = <$fh>;
close $fh;
chomp($_) foreach @records;

my ($l,$h) = most_asleep(@records);
print "$l\n$h\n";
