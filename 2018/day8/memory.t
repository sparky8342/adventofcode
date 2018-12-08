use strict;
use warnings;
use Memory qw(sum_metadata);
use Test::More tests => 2;

my $data = '2 3 0 3 10 11 12 1 1 0 1 99 2 1 1 2';

my ($sum,$value) = sum_metadata($data);
is($sum,138);
is($value,66);
