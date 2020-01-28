#!/usr/bin/perl
use strict;
use warnings;

use Digest::MD5 qw(md5_hex);

open my $fh, '<', 'input.txt';
chomp (my $key = <$fh>);
close $fh;

my $n = 1;
my ($part1, $part2) = (0, 0);
while (1) {
	my $hash = md5_hex($key . $n);
	if ($part1 == 0 && $hash =~ /^00000/) {
		$part1 = $n;
	}
	if ($part2 == 0 && $hash =~ /^000000/) {
		$part2 = $n;
		last;
	}
	$n++;
}
print "$part1\n$part2\n";
