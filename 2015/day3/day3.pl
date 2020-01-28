#!/usr/bin/perl
use strict;
use warnings;

sub move {
	my ($x, $y, $step, $visited) = @_;
	if ($step eq '^') { $y-- }
	if ($step eq 'v') { $y++ }
	if ($step eq '<') { $x-- }
	if ($step eq '>') { $x++ }
	$visited->{"$x $y"} = 1;
	return ($x, $y);
}

open my $fh, '<', 'input.txt';
my $data = <$fh>;
close $fh;

my ($x, $y) = (0, 0);
my $visited = {"0 0" => 1};

my @steps = split//, $data;
foreach my $step (@steps) {
	($x, $y) = move($x, $y, $step, $visited);
}

print scalar(keys %$visited) . "\n";

$visited = {"0 0" => 1};

for my $start (0, 1) {
	my ($x, $y) = (0, 0);
	for (my $i = $start; $i < @steps; $i += 2) {
		my $step = $steps[$i];
		($x, $y) = move($x, $y, $step, $visited);
	}
}

print scalar(keys %$visited) . "\n";
