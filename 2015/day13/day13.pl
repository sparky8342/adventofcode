#!/usr/bin/perl
use strict;
use warnings;

use Algorithm::Combinatorics qw(circular_permutations);

open my $fh, '<', 'input.txt';
my @data = <$fh>;
close $fh;

my %people;
foreach my $line (@data) {
	my ($person1, $gainlose, $units, $person2) = $line =~ /^(.*?) would (gain|lose) (\d+) happiness units by sitting next to (.*?)\./;
	if ($gainlose eq 'lose') {
		$units *= -1;
	}
	$people{$person1}{$person2} += $units;
	$people{$person2}{$person1} += $units;
}

# part 1
print find_best() . "\n";

# part 2
foreach my $p (keys %people) { 
	$people{$p}{me} = 0;
	$people{me}{$p} = 0;
}
print find_best() . "\n";

sub find_best {
	my @plist = keys %people;

	my $best = 0;
	my $iter = circular_permutations(\@plist);
	while (my $p = $iter->next) {
		my $happiness = 0;
		for (my $i = 0; $i < @$p - 1; $i++) {
			$happiness += $people{$p->[$i]}{$p->[$i + 1]};
		}
		$happiness += $people{$p->[@$p - 1]}{$p->[0]};
		if ($happiness > $best) {
			$best = $happiness;
		}
	}
	return $best;
}
