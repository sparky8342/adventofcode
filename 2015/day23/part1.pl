#!/usr/bin/perl
use strict;
use warnings;

# auto-generated from the input source

my $registers = { a => 0, b => 0 };
if ($registers->{a} == 1) { goto LINE18 }
$registers->{a}++;
$registers->{a} *= 3;
$registers->{a}++;
$registers->{a} *= 3;
$registers->{a} *= 3;
$registers->{a} *= 3;
$registers->{a}++;
$registers->{a} *= 3;
$registers->{a}++;
$registers->{a} *= 3;
$registers->{a}++;
$registers->{a}++;
$registers->{a} *= 3;
$registers->{a} *= 3;
$registers->{a} *= 3;
$registers->{a}++;
goto LINE39;
LINE18: $registers->{a} *= 3;
$registers->{a}++;
$registers->{a} *= 3;
$registers->{a}++;
$registers->{a}++;
$registers->{a} *= 3;
$registers->{a}++;
$registers->{a} *= 3;
$registers->{a}++;
$registers->{a}++;
$registers->{a} *= 3;
$registers->{a} *= 3;
$registers->{a}++;
$registers->{a}++;
$registers->{a} *= 3;
$registers->{a}++;
$registers->{a}++;
$registers->{a} *= 3;
$registers->{a}++;
$registers->{a}++;
$registers->{a} *= 3;
LINE39: if ($registers->{a} == 1) { goto LINE47 }
$registers->{b}++;
if ($registers->{a} % 2 == 0) { goto LINE45 }
$registers->{a} *= 3;
$registers->{a}++;
goto LINE46;
LINE45: $registers->{a} /= 2;
LINE46: goto LINE39;
LINE47:
print $registers->{b} . "\n";
