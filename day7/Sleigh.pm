package Sleigh;
use strict;
use warnings;

use parent "Exporter";
our @EXPORT_OK = qw(steps);

sub steps {
	my @ins = @_;

	my %tree;
	my %todo;
	foreach my $in (@ins) {
		$in =~ /^Step ([A-Z]) must be finished before step ([A-Z]) can begin\.$/;
		my ($parent,$child) = ($1,$2);
		$tree{$child}{$parent} = 1;
		$todo{$child} = 1;
		$todo{$parent} = 1;
	}

	my $step_order = '';
	my %done;

	while (keys %todo) {
		my %available;
		foreach my $node (keys %todo) {
			if (!exists($tree{$node})) {
				$available{$node} = 1;
			}

			my $ok = 1;
			foreach my $parent (keys %{$tree{$node}}) {
				if (!exists($done{$parent})) {
					$ok = 0;
					last;
				}
			}

			if ($ok) {
				$available{$node} = 1;
			}
		}

		my $step = (sort keys %available)[0];
		delete($todo{$step});
		$done{$step} = 1;
		$step_order .= $step;
	}

	return $step_order;
}
