#!/usr/bin/perl
use strict;
use warnings;

open my $fh, '<', 'input.txt';
my @frequencies = <$fh>;
close $fh;

my %seen;

my $freq = 0;
while (1) {
	foreach my $frequency (@frequencies) {
		$freq += $frequency;
		if (exists($seen{$freq})) {
			print "$freq\n";
			exit;
		}
		$seen{$freq} = 1;
	}
}
