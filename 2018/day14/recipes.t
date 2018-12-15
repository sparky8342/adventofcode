use strict;
use warnings;
use Recipes qw(make find);
use Test::More tests => 8;

is(make(9),'5158916779');
is(make(5),'0124515891');
is(make(18),'9251071085');
is(make(2018),'5941429882');

is(find('51589'),9);
is(find('01245'),5);
is(find('92510'),18);
is(find('59414'),2018);
