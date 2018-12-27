use strict;
use warnings;
use Device2 qw(run_program);

use Test::More tests => 1;

my @data = <DATA>;
chomp($_) foreach @data;

is(run_program([0,0,0,0,0,0],@data),6);

__DATA__
#ip 0
seti 5 0 1
seti 6 0 2
addi 0 1 0
addr 1 2 3
setr 1 0 0
seti 8 0 4
seti 9 0 5
