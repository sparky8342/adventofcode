#!/usr/bin/perl
use strict;
use warnings;
use Polymer qw(reduction);

open my $fh, '<', 'input.txt';
chomp(my $polymer = <$fh>);
close $fh;

$polymer = reduction($polymer);
print length($polymer) . "\n";
