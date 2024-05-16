#!/usr/bin/perl
use strict;
use warnings;
use List::Util qw(sum);

open my $fh, '<', 'input.txt';
chomp(my $data = <$fh>);
my $digits = [split//, $data];
my $orig_digits = [@$digits];

for (1..100) {
	$digits = phase($digits);
}	
print join('', @$digits[0..7]) . "\n";


$digits = [(@$orig_digits) x 10000];
my $start = join("", @$digits[0..6]);

@$digits = @$digits[$start..@$digits-1];

for (1..100) {
	$digits = phase2($digits);
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

sub phase2 {
	my ($digits, $start) = @_;

	my @new_digits;

	my $sum = sum(@$digits);
	for (my $i = 0; $i < @$digits; $i++) {
		push @new_digits, $sum % 10;
		$sum -= $digits->[$i];
	}

	return \@new_digits;
}
