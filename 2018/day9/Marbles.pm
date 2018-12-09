package Marbles;
use strict;
use warnings;

use parent "Exporter";
our @EXPORT_OK = qw(play);

sub play {
	my ($players,$no_marbles) = @_;

	# The key to this is to use a linked list.
	# It gives an incredible speed up when compared
	# to a regular array, because elements are only
	# added and no shifting along is needed as in a
	# splice operation.

	# I don't bother properly deleting elements, just
	# link past them.

	# The format is:
	# (value, previous, next)

	my @marbles = [(0,0,0)];

	my $current = 0;
	my $player = 0;
	my @scores = (0) x $players;

	for my $marble (1..$no_marbles) {
		if ($marble % 23 == 0) {
			$scores[$player] += $marble;

			# go back 7 marbles
			for (1..7) {
				$current = $marbles[$current]->[1];
			}

			$scores[$player] += $marbles[$current]->[0];

			# 'delete' the marble by linking past it
			my $next = $marbles[$current]->[2];
			my $previous = $marbles[$current]->[1];
			$marbles[$previous]->[2] = $next;
			$current = $next;
		}
		else {
			# move to next marble
			$current = $marbles[$current]->[2];

			# get next from current position
			my $next = $marbles[$current]->[2];

			# create new marble in between current and next
			push @marbles, [$marble,$current,$next];
			my $new = @marbles - 1;

			# set the current one to point to the new one as the next marble
			$marbles[$current]->[2] = $new;

			# set the old next to point back to the new one as the previous marble
			$marbles[$next]->[1] = $new;

			# move current to the new one
			$current = $new;
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
