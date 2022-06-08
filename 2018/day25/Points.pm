package Points;
use strict;
use warnings;

use parent "Exporter";
our @EXPORT_OK = qw(count_groups);

sub process_input {
	my ($input) = @_;
	my @points;
	foreach my $line (@$input) {
		my @coords = split/,/,$line;
		push @points, { x => $coords[0], y => $coords[1], z => $coords[2], t => $coords[3] };
	}

	return \@points;
}

sub count_groups {
	my @input = @_;
	my $points = process_input(\@input);
	my @groups = map {[ $_ ]} @$points;

	while (1) {
		my $move_from;
		my $move_to;
		OUTER:
		for (my $i = 0; $i < scalar(@groups); $i++) {
			for (my $j = $i + 1; $j < scalar(@groups); $j++) {
				my $group1 = $groups[$i];
				my $group2 = $groups[$j];
				for (my $k = 0; $k < scalar(@$group1); $k++) {
					for (my $l = 0; $l < scalar(@$group2); $l++) {
						my $point1 = $group1->[$k];
						my $point2 = $group2->[$l];
						my $dist = abs($point2->{x} - $point1->{x}) + abs($point2->{y} - $point1->{y}) + abs($point2->{z} - $point1->{z}) + abs($point2->{t} - $point1->{t});
						if ($dist <= 3) {
							$move_from = $i;
							$move_to = $j;
							last OUTER;
						}
					}
				}
			}
		}
		if (defined($move_from)) {
			push @{$groups[$move_to]}, @{$groups[$move_from]};
			splice @groups, $move_from, 1;
			next;
		}
		last;
	}

	return scalar(@groups);
}

1;
