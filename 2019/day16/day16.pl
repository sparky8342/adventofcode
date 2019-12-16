#!/usr/bin/perl
use strict;
use warnings;

open my $fh, '<', 'input.txt';
chomp(my $data = <$fh>);
my $digits = [split//, $data];

for (1..100) {
	$digits = phase($digits);
}	
print join('', @$digits[0..7]) . "\n";

sub phase {
	my ($digits) = @_;

	my @pattern = (0, 1, 0, -1);

	my @new_digits;
	for my $repeat (1..@$digits) {
		my $p = 0;
		if ($repeat == 1) {
			$p = 1;
		}

		my $repeat_count = 1;

		my $new_digit = 0;
		for my $digit (@$digits) {
			$new_digit += $digit * $pattern[$p];
			$repeat_count++;
			if ($repeat_count >= $repeat) {
				$p++;
				$p = 0 if $p == @pattern;
				$repeat_count = 0;
			}
		}

		$new_digit = substr($new_digit, length($new_digit) - 1, 1);
		push @new_digits, $new_digit;
	}
	return \@new_digits;
}
