use strict;
use warnings;
use Coordinates qw(largest_area);
use Test::More tests => 1;

my @data = (
	'1, 1',
	'1, 6',
	'8, 3',
	'3, 4',
	'5, 5',
	'8, 9'
);

is(largest_area(@data),17);
