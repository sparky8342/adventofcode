#!/usr/bin/perl
use strict;
use warnings;

open my $fh, '<', 'input.txt';
my @data = <$fh>;
close $fh;

my @reindeers;

foreach my $line (@data) {
	my ($name, $speed, $run_time, $rest_time) = $line =~ /^(.*?) can fly (\d+) km\/s for (\d+) seconds, but then must rest for (\d+)/;
	push @reindeers, {
		name => $name,
		speed => $speed,
		run_time => $run_time,
		rest_time => $rest_time,
		running => 1,
		phase_time => $run_time,
		score => 0,
		distance => 0
	};
}

my $time = 2503;

for (1..$time) {
	foreach my $reindeer (@reindeers) {
		if ($reindeer->{running} == 1) {
			$reindeer->{distance} += $reindeer->{speed};
			$reindeer->{phase_time}--;
			if ($reindeer->{phase_time} == 0) {
				$reindeer->{running} = 0;
				$reindeer->{phase_time} = $reindeer->{rest_time};
			}
		}
		elsif ($reindeer->{running} == 0) {
			$reindeer->{phase_time}--;
			if ($reindeer->{phase_time} == 0) {
				$reindeer->{running} = 1;
				$reindeer->{phase_time} = $reindeer->{run_time};
			}
		}
	}
	# score the reindeer(s) in the lead
	@reindeers = sort { $b->{distance} <=> $a->{distance} } @reindeers;
	my $best_dist = $reindeers[0]->{distance};
	my @scoring = grep { $_->{distance} == $best_dist } @reindeers;
	foreach my $r (@scoring) {
		$r->{score}++;
	}
}

# part 1, furthest distance
@reindeers = sort { $b->{distance} <=> $a->{distance} } @reindeers;
print $reindeers[0]->{distance} . "\n";

# part 2, highest score
@reindeers = sort { $b->{score} <=> $a->{score} } @reindeers;
print $reindeers[0]->{score} . "\n";
