#!/usr/bin/perl
use strict;
use warnings;

use Test::More tests => 2;

sub make_tree {
	my ($data) = @_;
	my $tree = {};
	foreach my $entry (@$data) {
		my ($a, $b) = split('\)', $entry);
		$tree->{$b} = $a;
	}
	return $tree;
}

sub total_orbits {
	my ($tree) = @_;
	my $orbits = 0;
	foreach my $o (keys %$tree) {
		$orbits++;
		while ($tree->{$o} ne 'COM') {
			$o = $tree->{$o};
			$orbits++;
		}
	}
	return $orbits;
}

sub distance {
	my ($tree,$start,$end) = @_;
	my %path;
	my $node = $start;
	my $steps = 0;
	while ($tree->{$node} ne 'COM') {
		$node = $tree->{$node};
		$steps++;
		$path{$node} = $steps;
	}

	$steps = 0;
	$node = $end;
	while ($tree->{$node} ne 'COM') {
		if (exists($path{$node})) {
			my $total = $path{$node} + $steps - 2;
			return $total;
		}
		$node = $tree->{$node};
		$steps++;
	}
}

my $test = [qw/COM)B B)C C)D D)E E)F B)G G)H D)I E)J J)K K)L/];
my $tree = make_tree($test);
is(total_orbits($tree), 42);

$test = [qw/COM)B B)C C)D D)E E)F B)G G)H D)I E)J J)K K)L K)YOU I)SAN/];
$tree = make_tree($test);
is(distance($tree, 'YOU', 'SAN'), 4);

open my $fh, '<', 'input.txt';
chomp(my @data = <$fh>);
close $fh;

$tree = make_tree(\@data);
my $orbits = total_orbits($tree);
my $distance = distance($tree, 'YOU', 'SAN');
print "$orbits\n$distance\n";
