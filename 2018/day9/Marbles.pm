package Marbles;
use strict;
use warnings;

use parent "Exporter";
our @EXPORT_OK = qw(play);

sub play {
	my ($players,$no_marbles) = @_;

	# linked list, each entry is a hash with prev, next and val keys

	my $marble = { val => 0 };
	$marble->{next} = $marble;
	$marble->{prev} = $marble;

	my $current = 0;
	my $player = 0;
	my @scores = (0) x $players;

	for my $marble_n (1..$no_marbles) {
		if ($marble_n % 23 == 0) {
			$scores[$player] += $marble_n;

			# go back 7 marbles
			for (1..7) {
				$marble = $marble->{prev};
			}

			$scores[$player] += $marble->{val};

			# 'delete' the marble by linking past it
			my $next = $marble->{next};
			my $prev = $marble->{prev};
			$marble->{prev}->{next} = $next;
			$marble = $next;
		}
		else {
			# move to next marble
			$marble = $marble->{next};

			# get next from current position
			my $next = $marble->{next};

			# create new marble in between current and next
			my $new = { prev => $marble, next => $next, val => $marble_n };

			# set the current one to point to the new one as the next marble
			$marble->{next} = $new;

			# set the old next to point back to the new one as the previous marble
			$next->{prev} = $new;

			# move current to the new one
			$marble = $new;
		}
		$player++;
		$player = 0 if $player == $players;
	}	

	my $max = 0;
	for (my $i = 0; $i < @scores; $i++) {
		$max = $scores[$i] if $scores[$i] > $max;
	}

	return $max;
}

sub pp {
	my @m = @_;

	my $p = 0;
	while (1) {
		print $m[$p]->[0] . ' ';
		$p = $m[$p]->[2];
		last if $p == 0;
	}
	print "\n";
}
