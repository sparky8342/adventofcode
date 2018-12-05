#!/usr/bin/perl
use strict;
use warnings;
use Polymer qw(reduction find_best_after_removal);

open my $fh, '<', 'input.txt';
chomp(my $polymer = <$fh>);
close $fh;

my $reduced = reduction($polymer);
print length($reduced) . "\n";

my $best = find_best_after_removal($polymer);
print length($best) . "\n";
