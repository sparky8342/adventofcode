package Nanobots;
use strict;
use warnings;

use Data::Dumper;

use parent "Exporter";
our @EXPORT_OK = qw(in_range most_in_range);

sub in_range {
	my @data = @_;
	my @bots = parse_data(@_);

	@bots = sort { $b->{r} <=> $a->{r} } @bots;
	my $strongest = $bots[0];

	return bots_in_range($strongest, \@bots);	
}

sub bots_in_range {
	my ($point, $bots) = @_;
	my $in_range = 0;
	foreach my $bot (@$bots) {
		my $dist = abs($bot->{x} - $point->{x}) + abs($bot->{y} - $point->{y}) + abs($bot->{z} - $point->{z});
		if ($dist <= $point->{r}) {
			$in_range++;
		}
	}
	return $in_range;
}
	
sub most_in_range {
	my @data = @_;
	my @bots = parse_data(@data);

	my $start = {};

	foreach my $bot (@bots) {
		foreach my $key ('x','y','z') {
			$start->{min_x} = $bot->{x} if !exists($start->{min_x}) || $bot->{x} < $start->{min_x};
			$start->{max_x} = $bot->{x} if !exists($start->{max_x}) || $bot->{x} > $start->{max_x};
			$start->{min_y} = $bot->{y} if !exists($start->{min_y}) || $bot->{y} < $start->{min_y};
			$start->{max_y} = $bot->{y} if !exists($start->{max_y}) || $bot->{y} > $start->{max_y};
			$start->{min_z} = $bot->{z} if !exists($start->{min_z}) || $bot->{z} < $start->{min_z};
			$start->{max_z} = $bot->{z} if !exists($start->{max_z}) || $bot->{z} > $start->{max_z};
		}
	}

	print Dumper $start;
}

sub divide {
	my ($box) = @_;

	my $mid_x = $box->{min_x} + int(($box->{max_x} - $box->{min_x}) / 2);
	my $mid_y = $box->{min_y} + int(($box->{max_y} - $box->{min_y}) / 2);
	my $mid_z = $box->{min_z} + int(($box->{max_z} - $box->{min_z}) / 2);

	my @new_boxes;

	foreach my $x ('min_x', 'max_x') {
		foreach my $y ('min_y', 'max_y') {
			foreach my $z ('min_z', 'max_z') {
				my $b = {
					min_x => $box->{min_x}, max_x => $box->{max_x},
					min_y => $box->{min_y}, max_y => $box->{max_y},
					min_z => $box->{min_z}, max_z => $box->{max_z}
				};
				$b->{$x} = $mid_x;
				$b->{$y} = $mid_y;
				$b->{$z} = $mid_z;
				push @new_boxes, $b;
			}
		}
	}

	return \@new_boxes;
}

sub parse_data {
	my @data = @_;

	my @bots;
	foreach my $line (@data) {
		$line =~ /<(-?[0-9]+),(-?[0-9]+),(-?[0-9]+)>, r=(\d+)/;
		push @bots, { x => $1, y => $2, z => $3, r => $4 };
	}

	return @bots;
}
