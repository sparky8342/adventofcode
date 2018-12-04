package Repose;
use strict;
use warnings;

use parent "Exporter";
our @EXPORT_OK = qw(most_asleep);

sub most_asleep {
	my @events = @_;

	@events = sort @events;

	my %guards;
	my $current_guard;
	my $sleep_start;

	my %longest = (guard => 0, sleep => 0, minute => 0, minute_frequency => 0);
	my %frequent = (guard => 0, minute => 0, minute_frequency => 0);

	foreach my $entry (@events) {
		$entry =~ /^.*:(\d{2})\] (.*)/;
		my $minute = $1;
		my $action = $2;	

		if ($action =~ /^Guard #(\d+)/) {
			$current_guard = $1;
		}
		elsif ($action =~ /falls asleep/) {
			$sleep_start = $minute;
		}
		elsif ($action =~ /wakes up/) {
			$guards{$current_guard}{sleep} += $minute - $sleep_start;
			my $longest_so_far = 0;
			if ($guards{$current_guard}{sleep} > $longest{sleep}) {
				$longest{sleep} = $guards{$current_guard}{sleep};
				$longest{guard} = $current_guard;
			}
			for (my $i = $sleep_start; $i < $minute; $i++) {
				$guards{$current_guard}{minutes}{$i}++;
				if ($guards{$current_guard}{minutes}{$i} > $longest{minute_frequency}) {
					$longest{minute_frequency} = $guards{$current_guard}{minutes}{$i};
					$longest{minute} = $i;
				}
				if ($guards{$current_guard}{minutes}{$i} > $frequent{minute_frequency}) {
					$frequent{minute_frequency} = $guards{$current_guard}{minutes}{$i};
					$frequent{minute} = $i;
					$frequent{guard} = $current_guard;
				}
			}
		}
	}

	return ($longest{guard} * $longest{minute}, $frequent{guard} * $frequent{minute});
}	
