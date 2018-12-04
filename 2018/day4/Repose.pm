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
			for (my $i = $sleep_start; $i < $minute; $i++) {
				$guards{$current_guard}{minutes}{$i}++;
			}
		}
	}

	my $longest_guard = (sort { $guards{$b}{sleep} <=> $guards{$a}{sleep} } keys %guards)[0];
	my $minutes = $guards{$longest_guard}{minutes};
	my $longest_minute = (sort { $minutes->{$b} <=> $minutes->{$a} } keys %$minutes)[0];

	return $longest_guard * $longest_minute;	
}	
