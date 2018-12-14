use strict;
use warnings;
use MineCarts qw(move);
use Test::More tests => 1;

my @data = <DATA>;
is(move(@data),'7,3');

__DATA__
/->-\        
|   |  /----\
| /-+--+-\  |
| | |  | v  |
\-+-/  \-+--/
  \------/
