use strict;
use warnings;
use MineCarts qw(move);
use Test::More tests => 1;

my @data = <DATA>;
is(move(0,@data),'6,4');

__DATA__
/>-<\  
|   |  
| /<+-\
| | | v
\>+</ |
  |   ^
  \<->/
