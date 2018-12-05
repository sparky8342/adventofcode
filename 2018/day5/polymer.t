use strict;
use warnings;
use Polymer qw(reduction);
use Test::More tests => 5;

is(reduction('aA'),'');
is(reduction('abBA'),'');
is(reduction('abAB'),'abAB');
is(reduction('aabAAB'),'aabAAB');
is(reduction('dabAcCaCBAcCcaDA'),'dabCBAcaDA');
