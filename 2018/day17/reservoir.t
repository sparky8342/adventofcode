use strict;
use warnings;
use Reservoir qw(run_water);
use Test::More tests => 2;

my @data = <DATA>;
chomp($_) foreach @data;

my ($total,$standing) = run_water(@data);
is($total,57);
is($standing,29);

__DATA__
x=495, y=2..7
y=7, x=495..501
x=501, y=3..7
x=498, y=2..4
x=506, y=1..2
x=498, y=10..13
x=504, y=10..13
y=13, x=498..504
