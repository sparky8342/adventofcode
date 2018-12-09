use strict;
use warnings;
use Marbles qw(play);
use Test::More tests => 6;

is(play(9,25),32);
is(play(10,1618),8317);
is(play(13,7999),146373);
is(play(17,1104),2764);
is(play(21,6111),54718);
is(play(30,5807),37305);
