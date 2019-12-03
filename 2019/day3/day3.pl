#!/usr/bin/perl
use strict;
use warnings;

use Test::More tests => 3;

sub closest_crossover {
	my ($wires) = @_;
	my %grid;
	my $best_dist = 99999;

	foreach my $wire_no (0..1) {
		my $wire = $wires->[$wire_no];
		my $x = 0;
		my $y = 0;
		foreach my $move (@$wire) {
			$move =~ /^([A-Z])(\d+)$/;
			my ($dir, $length) = ($1, $2);
			my ($dx, $dy) = (0, 0);
			if    ($dir eq 'U') { $dy = -1 }
			elsif ($dir eq 'D') { $dy =  1 }
			elsif ($dir eq 'L') { $dx = -1 }
			elsif ($dir eq 'R') { $dx =  1 }
			for (1..$length) {
				$x = $x + $dx;
				$y = $y + $dy;
				$grid{$x}{$y}++;

				if ($wire_no == 1 && $grid{$x}{$y} == 2) {
					my $dist = abs($x) + abs($y);
					if ($dist < $best_dist) {
						$best_dist = $dist;
					}
				}
			}
		}
	}

	return $best_dist;
}

my @tests = (
	{
		wires => [
			['R8','U5','L5','D3'],
			['U7','R6','D4','L4']
		],
		distance => 6
	},
	{
		wires => [
			['R75','D30','R83','U83','L12','D49','R71','U7','L72'],
			['U62','R66','U55','R34','D71','R55','D58','R83']
		],
		distance => 159
	},
	{
		wires => [
			['R98','U47','R26','D63','R33','U87','L62','D20','R33','U53','R51'],
			['U98','R91','D20','R16','D67','R40','U7','R15','U6','R7']
		],
		distance => 135
	}
);

foreach my $test (@tests) {
	is(closest_crossover($test->{wires}), $test->{distance});
}

my @wires;
open my $fh, '<', 'input.txt';
chomp(my $line = <$fh>);
push @wires, [split/,/,$line];
chomp($line = <$fh>);
push @wires, [split/,/,$line];
close $fh;

print closest_crossover(\@wires) . "\n";
