#!/usr/bin/perl
use strict;
use warnings;

open my $fh, '<', 'input.txt';
chomp(my @data = <$fh>);
close $fh;

my @moons;
foreach my $line (@data) {
	$line =~ /^\<x=(.*?), y=(.*?), z=(.*?)>$/;
	push @moons, { 
		position => { x => $1, y => $2, z => $3 },
		velocity => { x => 0, y => 0, z => 0 }
	};
} 

for (1..1000) {
	for (my $i = 0; $i < @moons; $i++) {
		my $moon = $moons[$i];
		for (my $j = 0; $j < @moons; $j++) {
			next if $i == $j;
			my $moon2 = $moons[$j];
			foreach my $axis ('x', 'y', 'z') {
				if ($moon->{position}->{$axis} > $moon2->{position}->{$axis}) {
					$moon->{velocity}->{$axis}--;
				}
				elsif ($moon->{position}->{$axis} < $moon2->{position}->{$axis}) {
					$moon->{velocity}->{$axis}++;
				}
			}
		}
	}

	for (my $i = 0; $i < @moons; $i++) {
		my $moon = $moons[$i];
		foreach my $axis ('x', 'y', 'z') {
			$moon->{position}->{$axis} += $moon->{velocity}->{$axis};
		}
	}
}

my $energy = 0;
foreach my $moon (@moons) {
	$energy += (abs($moon->{position}->{x}) + abs($moon->{position}->{y}) + abs($moon->{position}->{z}))
		* (abs($moon->{velocity}->{x}) + abs($moon->{velocity}->{y}) + abs($moon->{velocity}->{z}));
}

print "$energy\n";

