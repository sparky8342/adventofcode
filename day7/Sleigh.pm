package Sleigh;
use strict;
use warnings;

use parent "Exporter";
our @EXPORT_OK = qw(steps);

sub steps {
	my ($no_workers,$inc,@ins) = @_;

	my $tree = {};
	my $todo = {};
	foreach my $in (@ins) {
		$in =~ /^Step ([A-Z]) must be finished before step ([A-Z]) can begin\.$/;
		my ($parent,$child) = ($1,$2);
		$tree->{$child}{$parent} = 1;
		$todo->{$child} = 1;
		$todo->{$parent} = 1;
	}

	my $step_order = '';
	my $todo1 = { map { $_ => 1 } keys %$todo };
	my $done = {};
	while (keys %$todo1) {
		my $step = get_next_available($tree,$todo1,$done);
		delete($todo1->{$step});
		$done->{$step} = 1;
		$step_order .= $step;
	}

	my @workers;
	for (1..$no_workers) {
		push @workers, { step => undef, time => 0 };
	}

	$done = {};
	my $amount = scalar keys %$todo;
	my $time = -1;
	while (keys %$done < $amount) {
                for (my $i = 0; $i < @workers; $i++) {
                        my $worker = $workers[$i];
                        if ($worker->{time} > 0) {
                                $worker->{time}--;
                        }

                        if ($worker->{time} == 0) {
                                if ($worker->{step}) {
					$done->{$worker->{step}} = 1;
					$worker->{step} = undef;
                                }
                                my $step = get_next_available($tree,$todo,$done);
				next unless $step;
                                $worker->{step} = $step;
				delete($todo->{$worker->{step}});
                                $worker->{time} = ord($step) - 64 + $inc;
                        }
                }
		$time++;
        }

	return ($step_order,$time);
}

sub get_next_available {
	my ($tree,$todo,$done) = @_;
	my %available;

	foreach my $node (keys %$todo) {
		my $ok = 1;

		foreach my $parent (keys %{$tree->{$node}}) {
			if (!exists($done->{$parent})) {
				$ok = 0;
				last;
			}
		}

		if ($ok) {
			$available{$node} = 1;
		}
	}

	my $step = (sort keys %available)[0];
	return $step;
}
