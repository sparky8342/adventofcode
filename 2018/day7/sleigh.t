use strict;
use warnings;
use Sleigh qw(steps);
use Test::More tests => 2;

my $data = <<'END_DATA';
Step C must be finished before step A can begin.
Step C must be finished before step F can begin.
Step A must be finished before step B can begin.
Step A must be finished before step D can begin.
Step B must be finished before step E can begin.
Step D must be finished before step E can begin.
Step F must be finished before step E can begin.
END_DATA

my @ins = split("\n",$data);
chomp($_) foreach @ins;

my ($steps,$time) = steps(2,0,@ins);
is($steps,'CABDFE');
is($time,15);
