package ImmuneSystem;
use strict;
use warnings;

use Group;

use parent "Exporter";
our @EXPORT_OK = qw(battle);

sub battle {
	my @data = @_;
	my ($immune, $infection) = parse_data(@data);

	while (@$immune && @$infection) {
		target_selection($immune, $infection);

		my @all = (@$immune, @$infection);
		@all = sort { $b->{initiative} <=> $a->{initiative} } @all;

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
			$target->{units} -= $units_lost;
		}

		@$immune = grep { $_->{units} > 0 } @$immune;
		@$infection = grep { $_->{units} > 0 } @$infection;

	}

	my $total = 0;
	foreach my $group (@$immune, @$infection) {
		$total += $group->{units};
	}

	return $total;
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

		push @damage, [$power, $j];
	}

	@damage = sort { $b->[0] <=> $a->[0] } @damage;
	if (@damage) {
		return $damage[0][1];
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

		print "$line\n";
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
