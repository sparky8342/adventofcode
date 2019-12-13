package Nanobots;
use strict;
use warnings;

use parent "Exporter";
our @EXPORT_OK = qw(in_range most_in_range);

sub in_range {
	my @data = @_;
	my @bots = parse_data(@_);

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


sub most_in_range {
	my @data = @_;
	my @bots = parse_data(@data);

	my $limits = {};

	foreach my $bot (@bots) {
		foreach my $key ('x','y','z') {
			$limits->{minx} = $bot->{x} if $bot->{x} < $limits->{minx} || !exists($limits->{minx});
			$limits->{maxx} = $bot->{x} if $bot->{x} > $limits->{maxx} || !exists($limits->{maxx});
			$limits->{miny} = $bot->{y} if $bot->{y} < $limits->{miny} || !exists($limits->{miny});
			$limits->{maxy} = $bot->{y} if $bot->{y} < $limits->{maxy} || !exists($limits->{maxy});
			$limits->{minz} = $bot->{z} if $bot->{z} < $limits->{minz} || !exists($limits->{minz});
			$limits->{maxz} = $bot->{z} if $bot->{z} < $limits->{maxz} || !exists($limits->{maxz});
		}
	}

	my @points;
	for (1..1000) {
		my $x = $limits->{minx} + int(rand($limits->{maxx} - $limits->{minx}));
		my $y = $limits->{miny} + int(rand($limits->{maxy} - $limits->{miny}));
		my $z = $limits->{minz} + int(rand($limits->{maxz} - $limits->{minz}));
		my $point = { x => $x, y => $y, z => $z };
		$point->{bots} = bots_in_range($point,\@bots);
		push @points, $point;
	}

	for (1..1000) {
		foreach my $point (@points) {
	
			my $dx = int(rand(200)) - 10;
			my $dy = int(rand(200)) - 10;
			my $dz = int(rand(200)) - 10;

			$point->{x} += $dx;
			$point->{y} += $dy;
			$point->{z} += $dz;

			my $amount = bots_in_range($point,\@bots);

			if ($amount < $point->{bots}) {
				$point->{x} -= $dx;
				$point->{y} -= $dy;
				$point->{z} -= $dz;
			}
			elsif ($amount > $point->{bots}) {
				$point->{bots} = $amount;
			}
		}
	}

	@points = sort { $b->{bots} <=> $a->{bots} } @points;

	use Data::Dumper;
	print Dumper $points[0];


}

sub bots_in_range {
	my ($point,$bots) = @_;
	my $in_range = 0;
	foreach my $bot (@$bots) {
		my $dist = abs($bot->{x} - $point->{x}) + abs($bot->{y} - $point->{y}) + abs($bot->{z} - $point->{z});
		if ($dist <= $bot->{r}) {
			$in_range++;
		}
	}
	return $in_range;
}
	


sub most_in_range_old {
	my @data = @_;
	my @bots = parse_data(@data);

	my %galaxy;
	my $best = 0;
	my $bestpoint = [];
	foreach my $bot (@bots) {
		use Data::Dumper;
		print Dumper $bot;
		my $points = get_bot_points2($bot);
		foreach my $point (@$points) {
			$galaxy{$point->{x}}{$point->{y}}{$point->{z}}++;
			if ($galaxy{$point->{x}}{$point->{y}}{$point->{z}} > $best) {
				$best = $galaxy{$point->{x}}{$point->{y}}{$point->{z}};
				$bestpoint = $point;
			}
		}
	}
	
	return $bestpoint;
}

sub get_bot_points {
	my ($bot) = @_;

	my $points = {$bot->{x} => { $bot->{y} => { $bot->{z} => 1}}};
	for (1..$bot->{r}) {
		my $newpoints = {};
		foreach my $x (keys %$points) {
			foreach my $y (keys %{$points->{$x}}) {
				foreach my $z (keys %{$points->{$x}{$y}}) {
					$newpoints->{$x  }{$y  }{$z  } = 1;
					$newpoints->{$x+1}{$y  }{$z  } = 1;
					$newpoints->{$x-1}{$y  }{$z  } = 1;
					$newpoints->{$x  }{$y+1}{$z  } = 1;
					$newpoints->{$x  }{$y-1}{$z  } = 1;
					$newpoints->{$x  }{$y  }{$z+1} = 1;
					$newpoints->{$x  }{$y  }{$z-1} = 1;
				}
			}
		}
		$points = $newpoints;
		#use Data::Dumper;
		#print Dumper $points;
	}	

	return $points;
}

sub get_bot_points2 {
	my ($bot) = @_;

	my $points = [];
	for my $dx (-$bot->{r}..$bot->{r}) {
		for my $dy (-$bot->{r}..$bot->{r}) {
			for my $dz (-$bot->{r}..$bot->{r}) {
				my $point = { x => $bot->{x} + $dx, y => $bot->{y} + $dy, z => $bot->{z} + $dz };
				my $dist = abs($bot->{x} - $point->{x}) + abs($bot->{y} - $point->{y}) + abs($bot->{z} - $point->{z});
				if ($dist <= $bot->{r}) {
					push @$points, $point;
					use Data::Dumper;
					print Dumper $point;
				}
			}
		}
	}

	return $points;
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
