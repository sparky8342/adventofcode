use strict;
use warnings;
use Plants qw(iterate);
use Test::More tests => 1;

my @data = <DATA>;
my $initial = shift(@data);
($initial) = $initial =~ /: (.*)/;
shift(@data);
is(iterate($initial,@data),325);

__DATA__
initial state: #..#.#..##......###...###

...## => #
..#.. => #
.#... => #
.#.#. => #
.#.## => #
.##.. => #
.#### => #
#.#.# => #
#.### => #
##.#. => #
##.## => #
###.. => #
###.# => #
####. => #
