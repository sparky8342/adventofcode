use strict;
use warnings;
use Coordinates qw(find_areas);
use Test::More tests => 2;

my @data = (
	'1, 1',
	'1, 6',
	'8, 3',
	'3, 4',
	'5, 5',
	'8, 9'
);

my ($area,$safe_area) = find_areas(32,@data);
is($area,17);
is($safe_area,16);
