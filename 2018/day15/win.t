use strict;
use warnings;
use Bandits qw(find_winning_attack_power);
use Test::More tests => 1;

my @data = <DATA>;
chomp($_) foreach @data;

is(find_winning_attack_power(@data),4988);

__DATA__
#######
#.G...#
#...EG#
#.#.#G#
#..G#E#
#.....#
#######
