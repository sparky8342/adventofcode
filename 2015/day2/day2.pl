#!/usr/bin/perl
use strict;
use warnings;

open my $fh, '<', 'input.txt';
my @data = <$fh>;
close $fh;

my $paper = 0;
my $ribbon = 0;
foreach my $line (@data) {
	my ($l, $w, $h) = split/x/,$line;
	my @sides = (
		{ size => $l * $w, perimeter => $l * 2 + $w * 2 },
		{ size => $w * $h, perimeter => $w * 2 + $h * 2 },
		{ size => $h * $l, perimeter => $h * 2 + $l * 2 }
	);
	@sides = sort { $a->{size} <=> $b->{size} } @sides;

	foreach my $side (@sides) {
		$paper += $side->{size} * 2;
	}

	$paper += $sides[0]->{size};
	$ribbon += $sides[0]->{perimeter} + $l * $w * $h;
}
print "$paper\n$ribbon\n";
