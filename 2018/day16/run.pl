#!/usr/bin/perl
use strict;
use warnings;
use Device qw(process);

open my $fh, '<', 'input.txt';
my @data = <$fh>;
close $fh;
chomp($_) foreach @data;

@data = reverse @data;
my @program;
while ($data[0] !~ /^After/) {
	my $line = shift(@data);
	push @program, $line if $line =~ /^\d/;
}
@data = reverse @data;
@program = reverse @program;

my ($count,$r0) = process(\@data,\@program);
print "$count\n$r0\n";
