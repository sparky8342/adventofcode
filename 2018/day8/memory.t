use strict;
use warnings;
use Memory qw(process_data);
use Test::More tests => 2;

my $data = '2 3 0 3 10 11 12 1 1 0 1 99 2 1 1 2';

my ($sum,$value) = process_data($data);
is($sum,138);
is($value,66);
