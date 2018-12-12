use strict;
use warnings;
use FuelGrid qw(find_best_square);
use Test::More tests => 4;

is(find_best_square(18,3,3),'33,45');
is(find_best_square(42,3,3),'21,61');
is(find_best_square(18,1,300),'90,269,16');
is(find_best_square(42,1,300),'232,251,12');
