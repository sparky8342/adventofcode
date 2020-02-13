#!/usr/bin/perl
use strict;
use warnings;

use Storable qw(dclone);

use constant SPELLS => [
	{
		name => 'magic missile',
		cost => 53,
		damage => 4,
		heal => 0
	},
	{
		name => 'drain',
		cost => 73,
		damage => 2,
		heal => 2
	},
	{
		name => 'shield',
		cost => 113,
		damage => 0,
		heal => 0,
		effect => { time => 6, action => { armour => 7 }}
	},
	{
		name => 'poison',
		cost => 173,
		damage => 0,
		heal => 0,
		effect => { time => 6, action => { damage => 3 }}
	},
	{
		name => 'recharge',
		cost => 229,
		damage => 0,
		heal => 0,
		effect => { time => 5, action => { mana => 101 }}
	}
];

use constant MIN_MANA => 53; # cheapest spell (magic missile)

my $start_enemy = {};
open my $fh, '<', 'input.txt';
foreach my $key ('hit_points', 'damage') {
	chomp(my $line = <$fh>);
	$line =~ /(\d+)/;
	$start_enemy->{$key} = $1;
}
close $fh;

my $start_player = { hit_points => 50, mana => 500, armour => 0, spent => 0, effects => [] };

my $min_mana = 9999;

fight($start_player, $start_enemy);

print "$min_mana\n";

sub fight {
	my ($pl, $en) = @_;

	if ($pl->{mana} < MIN_MANA) {
		return;
	}

	foreach my $sp (@{&SPELLS}) {
		if ($sp->{cost} > $pl->{mana}) {
			# player can't afford spell
			next;
		}

		my $player = dclone($pl);
		my $enemy = dclone($en);

		my $spell = dclone($sp); # need to avoid altering the source spell :(

		# player turn
		$player->{mana} -= $spell->{cost};
		$player->{spent} += $spell->{cost};

		if ($player->{spent} > $min_mana) {
			next;
		}

		$player->{hit_points} += $spell->{heal};
		my $damage = $spell->{damage};

		# add effect if one was cast
		if (exists($spell->{effect})) {
			push @{$player->{effects}}, $spell->{effect};
		}

		# process effects
		foreach my $effect (@{$player->{effects}}) {
			my $action = $effect->{action};
			if (exists($action->{damage})) {
				$damage += $action->{damage};
			}
			elsif (exists($action->{mana})) {
				$player->{mana} += $action->{mana};
			}
			elsif (exists($action->{armour})) {
				$player->{armour} += $action->{armour};
			}
			$effect->{time}--;
		}
		$player->{effects} = [ grep { $_->{time} > 0 } @{$player->{effects}} ];

		$enemy->{hit_points} -= $damage;
		if ($enemy->{hit_points} <= 0) {
			if ($player->{spent} < $min_mana) {
				$min_mana = $player->{spent};
			}
			next;
		}

		# enemy turn
		$damage = $enemy->{damage} - $player->{armour};
		$damage = 1 if $damage < 1;

		$player->{hit_points} -= $damage;
		if ($player->{hit_points} <= 0) {
			next;
		}

		$player->{armour} = 0;

		fight($player, $enemy);
	}
}
