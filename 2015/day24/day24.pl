#!/usr/bin/perl
use strict;
use warnings;

use Algorithm::Combinatorics qw(combinations);

open my $fh, '<', 'input.txt';
chomp(my @packages = <$fh>);
close $fh;

@packages = sort { $b <=> $a } @packages;

for my $buckets (3, 4) {
	my $total = 0;
	$total += $_ foreach @packages;
	my $target = $total / $buckets;

	$total = 0;
	my $num = 1;
	for my $package (@packages) {
		$total += $package;
		$num++;
		if ($total >= $target) {
			last;
		}
	}

	my $min_qe;
	my $iter = combinations(\@packages, $num);
	while (my $c = $iter->next) {
		my $total = 0;
		$total += $_ foreach @$c;
		next unless $total == $target;

		my $qe = 1;
		$qe *= $_ foreach @$c;
		if (!defined($min_qe) || $qe < $min_qe) {
			$min_qe = $qe;
		}
	}
	print "$min_qe\n";
}
