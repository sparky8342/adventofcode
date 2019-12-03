#!/usr/bin/perl
use strict;
use warnings;

use Test::More tests => 3;

sub closest_crossover {
	my ($wires) = @_;
	my %grid;
	# part1
	my $best_dist = 99999;
	# part2
	my $best_steps = 99999;

	foreach my $wire_no (0..1) {
		my $wire = $wires->[$wire_no];
		my $x = 0;
		my $y = 0;
		my $steps = 0;
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
				$steps++;
				if (!exists($grid{$wire_no}{$x}{$y})) {
					$grid{$wire_no}{$x}{$y} = $steps;
				}

				if ($wire_no == 1 && exists($grid{0}{$x}{$y})) {
					my $dist = abs($x) + abs($y);
					if ($dist < $best_dist) {
						$best_dist = $dist;
					}

					my $total_steps = $grid{0}{$x}{$y} + $grid{1}{$x}{$y};
					if ($total_steps < $best_steps) {
						$best_steps = $total_steps;
					}
				}

			}
		}
	}

	return ($best_dist, $best_steps);
}

my @tests = (
	{
		wires => [
			['R8','U5','L5','D3'],
			['U7','R6','D4','L4']
		],
		distance => 6,
		steps => 30
	},
	{
		wires => [
			['R75','D30','R83','U83','L12','D49','R71','U7','L72'],
			['U62','R66','U55','R34','D71','R55','D58','R83']
		],
		distance => 159,
		steps => 610
	},
	{
		wires => [
			['R98','U47','R26','D63','R33','U87','L62','D20','R33','U53','R51'],
			['U98','R91','D20','R16','D67','R40','U7','R15','U6','R7']
		],
		distance => 135,
		steps => 410
	}
);

foreach my $test (@tests) {
	is(closest_crossover($test->{wires}), ($test->{distance}, $test->{steps}));
}

open my $fh, '<', 'input.txt';
chomp(my @wires = <$fh>);
close $fh;

@wires = map { [split(/,/,$_)] } @wires;

print join("\n",closest_crossover(\@wires)) . "\n";
