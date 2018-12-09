package Marbles;
use strict;
use warnings;

use parent "Exporter";
our @EXPORT_OK = qw(play);

sub play {
	my ($players,$no_marbles) = @_;

	my @marbles = (0);
	my $current = 0;
	my $player = 0;
	my @scores = (0) x $players;

	for my $marble (1..$no_marbles) {
		if ($marble % 23 == 0) {
			$scores[$player] += $marble;
			$current -= 7;
			if ($current < 0) {
				$current = scalar(@marbles) + $current;
			}
			$scores[$player] += $marbles[$current];
			splice(@marbles,$current,1);
		}
		else {
			$current++;
			$current = 0 if $current == scalar(@marbles);
			splice(@marbles,$current+1,0,$marble);
			$current++;
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
