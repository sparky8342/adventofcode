#!/usr/bin/perl
use strict;
use warnings;

sub gcd {
	my ($x, $y) = @_;
	while ($x) { ($x, $y) = ($y % $x, $x) }
	$y
}

sub lcm {
	my ($x, $y) = @_;
	($x && $y) and $x / gcd($x, $y) * $y or 0
}

open my $fh, '<', 'input.txt';
chomp(my @data = <$fh>);
close $fh;

my @moons;
my @states;
foreach my $line (@data) {
	$line =~ /^\<x=(.*?), y=(.*?), z=(.*?)>$/;
	push @moons, { 
		position => { x => $1, y => $2, z => $3 },
		velocity => { x => 0, y => 0, z => 0 }
	};
	push @states, { 
		position => { x => $1, y => $2, z => $3 },
		velocity => { x => 0, y => 0, z => 0 }
	};
} 

my $steps = 0;
my %cycles;

while (1) {
	for (my $i = 0; $i < @moons - 1; $i++) {
		my $moon = $moons[$i];
		for (my $j = $i + 1; $j < @moons; $j++) {
			my $moon2 = $moons[$j];
			foreach my $axis ('x', 'y', 'z') {
				if ($moon->{position}->{$axis} > $moon2->{position}->{$axis}) {
					$moon->{velocity}->{$axis}--;
					$moon2->{velocity}->{$axis}++;
				}
				elsif ($moon->{position}->{$axis} < $moon2->{position}->{$axis}) {
					$moon->{velocity}->{$axis}++;
					$moon2->{velocity}->{$axis}--;
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
	
	# part 1
	$steps++;
	if ($steps == 1000) {
		my $energy = 0;
		foreach my $moon (@moons) {
			$energy += (abs($moon->{position}->{x}) + abs($moon->{position}->{y}) + abs($moon->{position}->{z}))
				* (abs($moon->{velocity}->{x}) + abs($moon->{velocity}->{y}) + abs($moon->{velocity}->{z}));
		}

		print "$energy\n";
	}

	# part 2
	for my $axis ('x','y','z') {
		next if exists $cycles{$axis};
		my $ok = 1;
		for (my $i = 0; $i < @moons; $i++) {
			if (!($moons[$i]->{position}->{$axis} == $states[$i]->{position}->{$axis}
				&& $moons[$i]->{velocity}->{$axis} == $states[$i]->{velocity}->{$axis})) {
					$ok = 0;
					last;
			}
		}
		if ($ok == 1) {
			$cycles{$axis} = $steps;
			if (keys %cycles == 3) {
				print lcm(lcm($cycles{x}, $cycles{y}), $cycles{z}) . "\n";
				exit;
			}
		}
	}
}
