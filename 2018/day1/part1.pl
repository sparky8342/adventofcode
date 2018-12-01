#!/usr/bin/perl
use strict;
use warnings;

open my $fh, '<', 'input.txt';
my @frequencies = <$fh>;
close $fh;

my $freq = 0;
foreach my $frequency (@frequencies) {
	$freq += $frequency;
}

print "$freq\n";
