use strict;
use warnings;
use Nanobots qw(in_range);
use Test::More tests => 1;

my @data = <DATA>;
is(in_range(@data),7);

__DATA__
pos=<0,0,0>, r=4
pos=<1,0,0>, r=1
pos=<4,0,0>, r=3
pos=<0,2,0>, r=1
pos=<0,5,0>, r=3
pos=<0,0,3>, r=1
pos=<1,1,1>, r=1
pos=<1,1,2>, r=1
pos=<1,3,1>, r=1
