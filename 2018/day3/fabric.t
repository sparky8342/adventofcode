use strict;
use warnings;
use Fabric qw(overlap);
use Test::More tests => 1;

my @data = (
	'#1 @ 1,3: 4x4',
	'#2 @ 3,1: 4x4',
	'#3 @ 5,5: 2x2'
);

is(overlap(@data),4);
