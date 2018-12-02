use strict;
use warnings;
use Inventory qw(checksum);
use Test::More tests => 1;

my @data = qw/abcdef bababc abbcde abcccd aabcdd abcdee ababab/;
is(checksum(@data),12);
