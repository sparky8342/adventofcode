#!/usr/bin/perl
use strict;
use warnings;

use Algorithm::Combinatorics qw(variations_with_repetition);

open my $fh, '<', 'input.txt';
my @data = <$fh>;
close $fh;

my @ingredients;
my @properties;

for my $line (@data) {
	my ($name, $props) = split(': ', $line);
	my $entry = {map { split(' ',$_) } split(', ', $props)};
	if (!@properties) {
		@properties = keys %$entry;
	}
	$entry->{name} = $name;
	push @ingredients, $entry;
}

my @sizes;
for (1..97) {
	push @sizes, $_;
}

my $best = 0;
my $best_with_500 = 0;
my $iter = variations_with_repetition(\@sizes, scalar(@ingredients));
while (my $v = $iter->next) {
	my $sum = 0;
	$sum += $_ foreach @$v;
	if ($sum == 100) {
		my %combo;
		for (my $i = 0; $i < @ingredients; $i++) {
			foreach my $property (@properties) {
				$combo{$property} += $ingredients[$i]->{$property} * $v->[$i];
			}
		}
		if (grep { $_ <= 0 } values %combo) {
			next;
		}
		my $total = 1;
		foreach my $key (grep { $_ ne 'calories' } keys %combo) {
			$total *= $combo{$key};
		}
		if ($total > $best) {
			$best = $total;
		}
		if ($combo{calories} == 500 && $total > $best_with_500) {
			$best_with_500 = $total;
		}
	}
}
print "$best\n$best_with_500\n";
