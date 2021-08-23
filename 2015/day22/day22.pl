#!/usr/bin/perl
use strict;
use warnings;

use constant START_HEALTH => 50;
use constant START_MANA => 500;

use constant MAGIC_MISSILE_COST => 53;
use constant MAGIC_MISSILE_DAMAGE => 4;

use constant DRAIN_COST => 73;
use constant DRAIN_DAMAGE => 2;
use constant DRAIN_HEAL => 2;

use constant SHIELD_COST => 113;
use constant SHIELD_TIME => 6;
use constant SHIELD_ARMOUR => 7;

use constant POISON_COST => 173;
use constant POISON_TIME => 6;
use constant POISON_DAMAGE => 3;

use constant RECHARGE_COST => 229;
use constant RECHARGE_TIME => 5;
use constant RECHARGE_MANA => 101;

open my $fh, '<', 'input.txt';
my $enemy_hp = <$fh>;
my $enemy_damage = <$fh>;
close $fh;
($enemy_hp) = $enemy_hp =~ /(\d+)/;
($enemy_damage) = $enemy_damage =~ /(\d+)/;

my $mana = bfs($enemy_hp, $enemy_damage, 0);
print "$mana\n";

$mana = bfs($enemy_hp, $enemy_damage, 1);
print "$mana\n";

sub process_effects {
	my ($state) = @_;

	if ($state->{poison} > 0) {
		$state->{enemy_hp} -= POISON_DAMAGE;
		$state->{poison}--;
	}
	if ($state->{recharge} > 0) {
		$state->{mana} += RECHARGE_MANA;
		$state->{recharge}--;
	}
	if ($state->{shield} > 0) {
		$state->{armour} = SHIELD_ARMOUR;
		$state->{shield}--;
	}
}

sub bfs {
	my ($enemy_hp, $enemy_damage, $hard_mode) = @_;

	my $smallest_mana_used = 9999;
	
	my $start = {
		player_turn => 1,
		hp => START_HEALTH,
		mana => START_MANA,
		enemy_hp => $enemy_hp,
		shield => 0,
		poison => 0,
		recharge => 0,
		armour => 0,
		mana_used => 0
	};

	my @queue = ($start);

	while (@queue > 0) {
		my $state = shift @queue;

		if ($state->{mana_used} >= $smallest_mana_used) {
			next;
		}

		if ($state->{player_turn} == 1) {

			my $hp = $state->{hp};
			if ($hard_mode) {
				$hp--;
				if ($hp <= 0) {
					next;
				}
			}

			my $state2 = {
				player_turn => 0,
				hp => $hp,
				mana => $state->{mana},
				enemy_hp => $state->{enemy_hp},
				shield => $state->{shield},
				poison => $state->{poison},
				recharge => $state->{recharge},
				armour => 0,
				mana_used => $state->{mana_used}
			};

			process_effects($state2);

			if ($state2->{enemy_hp} <= 0) {
				if ($state2->{mana_used} < $smallest_mana_used) {
					$smallest_mana_used = $state2->{mana_used};
				}
				next;
			}

			foreach my $spell ('magic_missile', 'drain', 'shield', 'poison', 'recharge') {
				my $new_state = {
					player_turn => 0,
					hp => $hp,
					mana => $state2->{mana},
					enemy_hp => $state2->{enemy_hp},
					shield => $state2->{shield},
					poison => $state2->{poison},
					recharge => $state2->{recharge},
					armour => 0,
					mana_used => $state2->{mana_used}
				};

				if ($spell eq 'magic_missile' && $new_state->{mana} >= MAGIC_MISSILE_COST) {
					$new_state->{mana} -= MAGIC_MISSILE_COST;
					$new_state->{mana_used} += MAGIC_MISSILE_COST;
					$new_state->{enemy_hp} -= MAGIC_MISSILE_DAMAGE;
				}
				elsif ($spell eq 'drain' && $new_state->{mana} >= DRAIN_COST) {
						$new_state->{mana} -= DRAIN_COST;
						$new_state->{mana_used} += DRAIN_COST;
						$new_state->{enemy_hp} -= DRAIN_DAMAGE;
						$new_state->{hp} += 2;
				}
				elsif ($spell eq 'shield' && $new_state->{shield} == 0 && $new_state->{mana} >= SHIELD_COST) {
						$new_state->{mana} -= SHIELD_COST;
						$new_state->{mana_used} += SHIELD_COST;
						$new_state->{shield} = SHIELD_TIME;
				}
				elsif ($spell eq 'poison' && $new_state->{poison} == 0 && $new_state->{mana} >= POISON_COST) {
					$new_state->{mana} -= POISON_COST;
					$new_state->{mana_used} += POISON_COST;
					$new_state->{poison} = POISON_TIME;
				}
				elsif ($spell eq 'recharge' && $new_state->{recharge} == 0 && $new_state->{mana} >= RECHARGE_COST) {
						$new_state->{mana} -= RECHARGE_COST;
						$new_state->{mana_used} += RECHARGE_COST;
						$new_state->{recharge} = RECHARGE_TIME;
				}
				else {
					next;
				}

				if ($new_state->{enemy_hp} <= 0) {
					if ($new_state->{mana_used} < $smallest_mana_used) {
						$smallest_mana_used = $new_state->{mana_used};
					}
					next;
				}

				push @queue, $new_state;
			}
		}
		elsif ($state->{player_turn} == 0) {
			my $new_state = {
				player_turn => 1,
				hp => $state->{hp},
				mana => $state->{mana},
				enemy_hp => $state->{enemy_hp},
				shield => $state->{shield},
				poison => $state->{poison},
				recharge => $state->{recharge},
				armour => 0,
				mana_used => $state->{mana_used}
			};

			process_effects($new_state);

			if ($new_state->{enemy_hp} <= 0) {
				if ($new_state->{mana_used} < $smallest_mana_used) {
					$smallest_mana_used = $new_state->{mana_used};
				}
				next;
			}

			my $damage = $enemy_damage - $new_state->{armour};
			$damage = 1 if $damage < 1;
			$new_state->{hp} -= $damage;

			if ($new_state->{hp} <= 0) {
				next;
			}

			push @queue, $new_state;
		}
	}

	return $smallest_mana_used;
}
