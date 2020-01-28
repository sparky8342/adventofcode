#!/usr/bin/perl
use strict;
use warnings;

open my $fh, '<', 'input.txt';
my $data = <$fh>;
close $fh;

my @chars = split//, $data;
my $floor = 0;
my $pos = 0;

for (my $i = 0; $i < @chars; $i++) {
	my $char = $chars[$i];
	if ($char eq '(') {
		$floor++;
	}
	elsif ($char eq ')') {
		$floor--;
		if ($floor == -1 && $pos == 0) {
			$pos = $i + 1;
		}
	}
}
print "$floor\n$pos\n";
