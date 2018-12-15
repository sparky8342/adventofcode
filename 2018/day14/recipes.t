use strict;
use warnings;
use Recipes qw(make);
use Test::More tests => 4;

is(make(9),'5158916779');
is(make(5),'0124515891');
is(make(18),'9251071085');
is(make(2018),'5941429882');
