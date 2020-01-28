#!/usr/bin/perl
use strict;
use warnings;

use Memoize qw(memoize);

open my $fh, '<', 'input.txt';
my @data = <$fh>;
close $fh;

my $program = '';
foreach my $line (@data) {
	chomp($line);
	my ($exp, $var) = split(' -> ', $line);
	$exp =~ s/AND/&/g;
	$exp =~ s/OR/|/g;
	$exp =~ s/LSHIFT/<</g;
	$exp =~ s/RSHIFT/>>/g;
	$exp =~ s/NOT/~/g;
	$exp =~ s/([a-z]+)/&$1/g;

	$program .= "sub $var { $exp }\n";
	$program .= "memoize('$var');\n";
}

$program .= "return &a()\n";

my $part1 = eval "$program";
$program =~ s/^sub b.*?$/sub b { $part1 }/m;
no warnings 'redefine';
my $part2 = eval "$program";
print "$part1\n$part2\n";
