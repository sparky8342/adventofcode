#!/usr/bin/perl
use strict;
use warnings;

use Test::More tests => 7;

sub fuel_needed {
	my ($mass) = @_;
	return int($mass / 3) - 2;
}

sub complete_fuel_needed {
	my ($mass) = @_;
	my $total = 0;
	my $fuel = fuel_needed($mass);
	while ($fuel > 0) {
		$total += $fuel;
		$fuel = fuel_needed($fuel);
	}
	return $total;
}

open my $fh, '<', 'input.txt';
my @modules = <$fh>;
close $fh;

# part 1
my @tests = (
	{ mass => 12, fuel => 2 },
	{ mass => 14, fuel => 2 },
	{ mass => 1969, fuel => 654 },
	{ mass => 100756, fuel => 33583 },
);
foreach my $test (@tests) {
	is(fuel_needed($test->{mass}), $test->{fuel});
}

my $total = 0;
foreach my $module (@modules) {
	$total += fuel_needed($module);
}
print "$total\n";

# part 2
@tests = (
	{ mass => 14, fuel => 2 },
	{ mass => 1969, fuel => 966 },
	{ mass => 100756, fuel => 50346 },
);
foreach my $test (@tests) {
	is(complete_fuel_needed($test->{mass}), $test->{fuel});
}

$total = 0;
foreach my $module (@modules) {
	$total += complete_fuel_needed($module);
}
print "$total\n";
