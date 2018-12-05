use strict;
use warnings;
use Polymer qw(reduction find_best_after_removal);
use Test::More tests => 6;

is(reduction('aA'),'');
is(reduction('abBA'),'');
is(reduction('abAB'),'abAB');
is(reduction('aabAAB'),'aabAAB');
is(reduction('dabAcCaCBAcCcaDA'),'dabCBAcaDA');

is(find_best_after_removal('dabAcCaCBAcCcaDA'),'daDA');
