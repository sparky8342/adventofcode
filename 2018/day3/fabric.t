use strict;
use warnings;
use Fabric qw(overlap);
use Test::More tests => 2;

my @data = (
	'#1 @ 1,3: 4x4',
	'#2 @ 3,1: 4x4',
	'#3 @ 5,5: 2x2'
);

my ($overlap_size,$non_overlap_id) = overlap(@data);
is($overlap_size,4);
is($non_overlap_id,3);
