#!/usr/bin/perl
use strict;
use warnings;
use Polymer qw(reduction);

open my $fh, '<', 'input.txt';
chomp(my $polymer = <$fh>);
close $fh;

$polymer = reduction($polymer);
print "$polymer\n";
print length($polymer) . "\n";

$polymer = reduction($polymer);
print "$polymer\n";
print length($polymer) . "\n";


