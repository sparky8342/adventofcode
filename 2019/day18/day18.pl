#!/usr/bin/perl
use strict;
use warnings;
use Math::BigInt;
use Memoize;

sub normalize {
	my ($grid) = @_;
	my $str = '';
	foreach my $row (@$grid) {
		$str .= join('',@$row);
	}
	return $str;
} 

memoize('search', NORMALIZER => 'normalize');

open my $fh, '<', 'input.txt';
chomp(my @data = <$fh>);
close $fh;

my $grid = [];
foreach my $row (@data) {
	push @$grid, [split//, $row];
}

my $height = @$grid;
my $width = @{$grid->[0]};

my $me;
my %doors;

for (my $y = 0; $y < $height; $y++) {
	for (my $x = 0; $x < $width; $x++) {
		if ($grid->[$y][$x] eq '@') {
			$me = { x => $x, y => $y };
		}
		elsif ($grid->[$y][$x] =~ /^([A-Z])$/) {
			$doors{$1} = [$x, $y];
		}
	}
}

# part 1
print_grid($grid);

my $steps = search($grid, [[$me->{x}, $me->{y}]]);
print "$steps\n";

$grid->[$me->{y} - 1][$me->{x} - 1] = '1';
$grid->[$me->{y} - 1][$me->{x}] = '#';
$grid->[$me->{y} - 1][$me->{x} + 1] = '2';
$grid->[$me->{y}][$me->{x} - 1] = '#';
$grid->[$me->{y}][$me->{x}] = '#';
$grid->[$me->{y}][$me->{x} + 1] = '#';
$grid->[$me->{y} + 1][$me->{x} - 1] = '3';
$grid->[$me->{y} + 1][$me->{x}] = '#';
$grid->[$me->{y} + 1][$me->{x} + 1] = '4';

# part 2
print_grid($grid);

$steps = search($grid, [
	[$me->{x} - 1, $me->{y} - 1],
	[$me->{x} + 1, $me->{y} - 1],
	[$me->{x} - 1, $me->{y} + 1],
	[$me->{x} + 1, $me->{y} + 1]
]);
print "$steps\n";

sub search {
	my ($grid, $robots) = @_;

	my @step_count;
	foreach my $robot (@$robots) {
		my $x = $robot->[0];
		my $y = $robot->[1];	

		my $keys = bfs($grid, $x, $y);

		if (@$keys == 0) {
			next;
		}

		foreach my $key (@$keys) {
			# move robot to key
			my $r = $grid->[$y][$x];
			$grid->[$y][$x] = '.';
			$grid->[$key->{y}][$key->{x}] = $r;
			$robot->[0] = $key->{x};
			$robot->[1] = $key->{y};

			# delete corresponding door from map
			my $door = uc($key->{key});
			my $d;
			if (exists($doors{$door})) {
				$d = $doors{$door};
				$grid->[$d->[1]][$d->[0]] = '.';
			}

			my $s = search($grid, $robots) + $key->{dist};
			push @step_count, $s;

			# put key, door and robot back
			$grid->[$key->{y}][$key->{x}] = $key->{key};
			if ($d) {
				$grid->[$d->[1]][$d->[0]] = $door;
			}
			$grid->[$y][$x] = $r;
			$robot->[0] = $x;
			$robot->[1] = $y;
		}
	}

	return 0 unless @step_count;
	@step_count = sort { $a <=> $b } @step_count;
	return $step_count[0];
}

sub print_grid {
	my ($grid) = @_;
	for my $y (0..@$grid - 1) {
		for my $x (0..@{$grid->[0]} - 1) {
			print $grid->[$y][$x];
		}
		print "\n";
	}
}

sub bfs {
	my ($grid, $x, $y) = @_;

	my @queue = ({ x => $x, y => $y, dist => 0 });
	my %visited;

	my @keys;
	my @doors;

	while (@queue) {
		my $space = shift(@queue);

		if (exists($visited{$space->{x}}{$space->{y}})) {
			next;
		}
		$visited{$space->{x}}{$space->{y}} = 1;

		if ($grid->[$space->{y}][$space->{x}] eq '#') {
			next;
		}
		elsif ($grid->[$space->{y}][$space->{x}] =~ /^([A-Z])$/) {
			next;
		}
		elsif ($grid->[$space->{y}][$space->{x}] =~ /^([a-z])$/) {
			$space->{key} = $1;
			push @keys, $space;
			next;
		}

		push @queue, (
			{ x => $space->{x} + 1, y => $space->{y},     dist => $space->{dist} + 1, dir => $space->{dir} || 4 },
			{ x => $space->{x} - 1, y => $space->{y},     dist => $space->{dist} + 1, dir => $space->{dir} || 3 },
			{ x => $space->{x}    , y => $space->{y} + 1, dist => $space->{dist} + 1, dir => $space->{dir} || 2 },
			{ x => $space->{x}    , y => $space->{y} - 1, dist => $space->{dist} + 1, dir => $space->{dir} || 1 }
		);
	}
	return \@keys;
}
