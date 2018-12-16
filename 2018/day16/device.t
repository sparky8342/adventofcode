use strict;
use warnings;
use Device qw(process);
use Test::More tests => 1;

my @data = <DATA>;
chomp($_) foreach @data;

is(process(@data),1);

__DATA__
Before: [3, 2, 1, 1]
9 2 1 2
After:  [3, 2, 2, 1]
