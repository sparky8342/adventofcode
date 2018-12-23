use strict;
use warnings;
use Cave qw(risk_level);

use Test::More tests => 1;

my @data = <DATA>;
chomp($_) foreach @data;

my ($rl) = risk_level(@data);
is($rl,114);

__DATA__
depth: 510
target: 10,10
