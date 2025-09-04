package ImmuneSystem;
use strict;
use warnings;

use Storable qw(dclone);

use Group;

use parent "Exporter";
our @EXPORT_OK = qw(parse_data battle find_winning_boost);

sub battle {
	my ($immune, $infection) = @_;

	$immune = dclone($immune);
	$infection = dclone($infection);

	while (@$immune && @$infection) {
		target_selection($immune, $infection);

		my @all = (@$immune, @$infection);
		@all = sort { $b->{initiative} <=> $a->{initiative} } @all;

		my $damage_done = 0;

		foreach my $group (@all) {
			if ($group->{units} <= 0 || !defined($group->{target})) {
				next;
			}

			my $target;
			if ($group->{type} eq 'immune') {
				$target = $infection->[$group->{target}];
			} elsif ($group->{type} eq 'infection') {
				$target = $immune->[$group->{target}];
			}

			my $damage = $group->damage($target);

			my $units_lost = int($damage / $target->{hit_points});
			if ($units_lost) {
				$target->{units} -= $units_lost;
				$damage_done = 1;
			}
		}

		@$immune = grep { $_->{units} > 0 } @$immune;
		@$infection = grep { $_->{units} > 0 } @$infection;

		if (!$damage_done) {
			return 'draw', 0;
		}
	}

	my $total = 0;
	foreach my $group (@$immune, @$infection) {
		$total += $group->{units};
	}
	my $winner;
	if (@$immune) {
		$winner = 'immune';
	} else {
		$winner = 'infection';
	}

	return $winner, $total;
}

sub target_selection {
	my ($immune, $infection) = @_;

	init_group($immune);
	init_group($infection);
	
	@$immune = sort { $b->effective_power() <=> $a->effective_power() || $b->{initiative} <=> $a->{initiative} } @$immune;
	@$infection = sort { $b->effective_power() <=> $a->effective_power() || $b->{initiative} <=> $a->{initiative} } @$infection;

	for my $group (@$immune) {
		my $target = pick_target($group, $infection);
		if (defined($target)) {
			$group->{target} = $target;
			$infection->[$target]->{targetted} = 1;
		}
	}

	for my $group (@$infection) {
		my $target = pick_target($group, $immune);
		if (defined($target)) {
			$group->{target} = $target;
			$immune->[$target]->{targetted} = 1;
		}
	}
}

sub pick_target {
	my ($group, $targets) = @_;

	my @damage;

	for (my $j = 0; $j < @$targets; $j++) {
		if ($targets->[$j]->{targetted} == 1 ) {
			next;
		}

		my $power = $group->damage($targets->[$j]);

		if ($power > 0) {
			push @damage, [$power, $targets->[$j]->effective_power(), $targets->[$j]->{initiative}, $j];
		}
	}

	if (@damage) {
		@damage = sort { $b->[0] <=> $a->[0] || $b->[1] <=> $a->[1] || $b->[2] <=> $a->[2] } @damage;
		return $damage[0][3];
	} else {
		return undef;
	}
}

sub init_group {
	my ($groups) = @_;
	foreach my $group (@$groups) {
		$group->{target} = undef;
		$group->{targetted} = 0;
	}
}

sub boost {
	my ($groups, $amount) = @_;
	foreach my $group (@$groups) {
		$group->{attack} += $amount;
	}
}

sub find_winning_boost {
	my ($immune, $infection) = @_;

	while (1) {
		boost($immune, 1);

		my ($winner, $winning_units) = battle($immune, $infection);
		if ($winner eq 'immune') {
			return $winning_units;
		}
	}
}

sub parse_data {
	my @data = @_;

	shift @data;

	my @immune;
	my @infection;
	my $infections = 0;

	foreach my $line (@data) {

		if ($line =~ /^\s*$/) {
			next;
		} elsif ($line =~ /Infection:/) {
			$infections = 1;
			next;
		}

		$line =~ /^(\d+) units each with (\d+) hit points(.*?)with an attack that does (\d+) (\w+) damage at initiative (\d+)$/;

		my $modifiers = $3;

		my $group = Group->new({
			units => $1,
			hit_points => $2,
			attack => $4,
			attack_type => $5,
			initiative => $6,
			weak => {},
			immune => {},
		});

		if ($modifiers ne " ") {
			$modifiers =~ s/ \((.*?)\) /$1/;
			foreach my $part (split/; /, $modifiers) {
				$part =~ /^(\w+) to (.*)/;
				my $type = $1;
				my $things = $2;
				foreach my $thing (split/, /, $things) {
					$group->{$type}->{$thing} = 1;
				}
			}
		}

		if ($infections) {
			$group->{type} = 'infection';
			push @infection, $group;
		} else {
			$group->{type} = 'immune';
			push @immune, $group;
		}
	}

	return \@immune, \@infection;
}
