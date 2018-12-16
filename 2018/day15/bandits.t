use strict;
use warnings;
use Bandits qw(combat);
use Test::More tests => 1;

my @data = <DATA>;
chomp($_) foreach @data;

is(combat(@data),27730);

__DATA__
#######
#.G...#
#...EG#
#.#.#G#
#..G#E#
#.....#
#######
