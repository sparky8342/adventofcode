#!/usr/bin/perl
use strict;
use warnings;

use Algorithm::Combinatorics qw(combinations);

use constant TARGET => 150;

open my $fh, '<', 'input.txt';
chomp(my @data = <$fh>);
close $fh;

my @seq = (0..@data - 1);

my $combos = 0;
my $min_containers = 0;
my $min_combos = 0;
for (my $k = 1; $k <= @data; $k++) {
	my $iter = combinations(\@seq, $k);
	ITER:
	while (my $c = $iter->next) {
		my $total = 0;
		foreach my $val (@$c) {
			$total += $data[$val];
			if ($total > TARGET) {
				next ITER;
			}
		}
		if ($total == TARGET) {
			$combos++;
			if ($min_containers == 0) {
				$min_containers = $k;
			}
			if ($min_containers == $k) {
				$min_combos++;
			}
		}

	}
}
print "$combos\n$min_combos\n";
