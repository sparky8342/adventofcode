#!/usr/bin/perl
use strict;
use warnings;

my %match = (
	children => 3,
	cats => 7,
	samoyeds => 2,
	pomeranians => 3,
	akitas => 0,
	vizslas => 0,
	goldfish => 5,
	trees => 3,
	cars => 2,
	perfumes => 1
);

open my $fh, '<', 'input.txt';
my @data = <$fh>;
close $fh;

my @aunts;
foreach my $line (@data) {
	chomp($line);
	my ($no, $info) = $line =~ /^Sue (\d+): (.*)/; 
	push @aunts, { map { split(': ', $_) } split(', ', $info), no => $no };
}

# part 1
my @search = @aunts;
foreach my $key (keys %match) {
	my $val = $match{$key};
	@search = grep { !exists($_->{$key}) || $_->{$key} == $val } @search;
}
print $search[0]->{no} . "\n";

# part 2
@search = @aunts;
foreach my $key (keys %match) {
	my $val = $match{$key};
	if ($key eq 'cats' || $key eq 'trees') {
		@search = grep { !exists($_->{$key}) || $_->{$key} > $val } @search;
	}
	elsif ($key eq 'pomeranians' || $key eq 'goldfish') {
		@search = grep { !exists($_->{$key}) || $_->{$key} < $val } @search;
	}
	else {
		@search = grep { !exists($_->{$key}) || $_->{$key} eq $val } @search;
	}
}
print $search[0]->{no} . "\n";
