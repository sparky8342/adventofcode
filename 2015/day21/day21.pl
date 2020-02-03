#!/usr/bin/perl
use strict;
use warnings;

use Algorithm::Combinatorics qw(variations);

use constant WEAPONS => [
	{ name => 'Dagger',     cost => 8,  damage => 4 },
	{ name => 'Shortsword', cost => 10, damage => 5 },
	{ name => 'Warhammer',  cost => 25, damage => 6 },
	{ name => 'Longsword',  cost => 40, damage => 7 },
	{ name => 'Greataxe',   cost => 74, damage => 8 }
];

use constant GARMENTS => [
	{ name => 'Plain Clothes', cost => 0,   armour => 0 },
	{ name => 'Leather',       cost => 13,  armour => 1 },
	{ name => 'Chainmail',     cost => 31,  armour => 2 },
	{ name => 'Splintmail',    cost => 53,  armour => 3 },
	{ name => 'Bandedmail',    cost => 75,  armour => 4 },
	{ name => 'Platemail',     cost => 102, armour => 5 }
];

use constant RINGS => [
	{ name => 'Damage +1',  cost => 25,  damage => 1, armour => 0 },
	{ name => 'Damage +2',  cost => 50,  damage => 2, armour => 0 },
	{ name => 'Damage +3',  cost => 100, damage => 3, armour => 0 },
	{ name => 'Defense +1', cost => 20,  damage => 0, armour => 1 },
	{ name => 'Defense +2', cost => 40,  damage => 0, armour => 2 },
	{ name => 'Defense +3', cost => 80,  damage => 0, armour => 3 }
];

my $enemy = {};
open my $fh, '<', 'input.txt';
foreach my $key ('hit_points', 'damage', 'armour') {
	chomp(my $line = <$fh>);
	$line =~ /(\d+)/;
	$enemy->{$key} = $1;
}
close $fh;

my @ring_ids;
for (0..5) {
	push @ring_ids, $_;
}

my $cheapest = 9999999;
my $losing_gold = 0;

foreach my $weapon (@{&WEAPONS}) {
	my $wcost = $weapon->{cost};
	my $damage = $weapon->{damage};

	foreach my $clothes (@{&GARMENTS}) {
		my $cost = $wcost + $clothes->{cost}; 
		my $armour = $clothes->{armour};

		for (my $no_rings = 0; $no_rings <= 2; $no_rings++) {
			if ($no_rings == 0) {
				if (simulate($damage, $armour)) {
					if ($cost < $cheapest) {
						$cheapest = $cost;
					}
				}
				elsif ($cost > $losing_gold) {
					$losing_gold = $cost;
				}
			}
			else {
				my $iter = variations(\@ring_ids, $no_rings);
				while (my $v = $iter->next) {

					my ($rdamage, $rarmour, $rcost) = ($damage, $armour, $cost);
					foreach my $id (@$v) {
						$rdamage += RINGS->[$id]->{damage};
						$rarmour += RINGS->[$id]->{armour};
						$rcost += RINGS->[$id]->{cost};
					}
					if (simulate($rdamage, $rarmour)) {
						if ($rcost < $cheapest) {
							$cheapest = $rcost;
						}
					}
					elsif ($rcost > $losing_gold) {
						$losing_gold = $rcost;
					}
				}
			}
		}
	}
}
print "$cheapest\n$losing_gold\n";

sub simulate {
	my ($damage, $armour) = @_;

	my $me = { hit_points => 100, damage => $damage, armour => $armour };
	my $opponent = { hit_points => $enemy->{hit_points}, damage => $enemy->{damage}, armour => $enemy->{armour} };

	while (1) {
		my $attack = $me->{damage} - $opponent->{armour};
		$attack = 1 if $attack < 1;
		$opponent->{hit_points} -= $attack;
		if ($opponent->{hit_points} <= 0) {
			return 1;
		}

		$attack = $opponent->{damage} - $me->{armour};
		$attack = 1 if $attack < 1;
		$me->{hit_points} -= $attack;
		if ($me->{hit_points} <= 0) {
			return 0;
		}
	}
}
