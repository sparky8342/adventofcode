package Nanobots;
use strict;
use warnings;

use parent "Exporter";
our @EXPORT_OK = qw(in_range);

sub in_range {
	my @data = @_;

	my @bots;
	foreach my $line (@data) {
		$line =~ /<(-?[0-9]+),(-?[0-9]+),(-?[0-9]+)>, r=(\d+)/;
		push @bots, { x => $1, y => $2, z => $3, r => $4 };
	}

	@bots = sort { $b->{r} <=> $a->{r} } @bots;
	my $strongest = $bots[0];

	my $in_range = 0;
	foreach my $bot (@bots) {
		my $dist = abs($bot->{x} - $strongest->{x}) + abs($bot->{y} - $strongest->{y}) + abs($bot->{z} - $strongest->{z});
		if ($dist <= $strongest->{r}) {
			$in_range++;
		}
	}

	return $in_range;
}
