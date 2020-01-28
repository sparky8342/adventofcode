#!/usr/bin/perl
use strict;
use warnings;

sub iterate {
	my ($seq) = @_;
	my $new_seq = '';
	while ($seq =~ /(\d)(\1*)/g) {
		my $n = $1;
		my $m = $n . $2;
		$new_seq .= length($m) . $n;
	}
	return $new_seq;
}

open my $fh, '<', 'input.txt';
my $seq = <$fh>;
close $fh;

for (1..40) {
	$seq = iterate($seq);
}
print length($seq) . "\n";

for (1..10) {
	$seq = iterate($seq);
}
print length($seq) . "\n";
