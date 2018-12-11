use strict;
use warnings;
use FuelGrid qw(find_best_square);
use Test::More tests => 1;


is(find_best_square(18),'33,45');
