use strict;
use warnings;
use Inventory qw(find_similar);
use Test::More tests => 1;

my @data = qw/abcde fghij klmno pqrst fguij axcye wvxyz/;
is(find_similar(@data),'fgij');
