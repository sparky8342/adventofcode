#!/usr/bin/perl
use strict;
use warnings;
use Marbles qw(play);

open my $fh, '<', 'input.txt';
my $data = <$fh>;
close $fh;

$data =~ /^(\d+).*?(\d+)/;
my ($players,$no_marbles) = ($1,$2);
print "$players,$no_marbles\n";
print play($players,$no_marbles) . "\n";
