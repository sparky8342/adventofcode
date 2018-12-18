use strict;
use warnings;
use Acres qw(iterate);

use Test::More tests => 1;

my @data = <DATA>;
chomp($_) foreach @data;

is(iterate(10,@data),1147);

__DATA__
.#.#...|#.
.....#|##|
.|..|...#.
..|#.....#
#.#|||#|#|
...#.||...
.|....|...
||...#|.#|
|.||||..|.
...#.|..|.
