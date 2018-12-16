#!/usr/bin/perl
use strict;
use warnings;
use Device qw(process);

open my $fh, '<', 'input.txt';
my @data = <$fh>;
close $fh;
chomp($_) foreach @data;

@data = reverse @data;
while ($data[0] !~ /^After/) {
	shift(@data);
}
@data = reverse @data;

print process(@data) . "\n";
