#!/usr/bin/perl
use strict;
use warnings;

print sum_factors(976) . "\n" . sum_factors(10551376) . "\n";

sub sum_factors {
	my ($n) = @_;
	my $sum = 0;
	for my $x (1..sqrt($n)) {
		my $d = $n / $x;
		if ($d == int($d)) {
			$sum += $x + $d;
		}
	}
	return $sum;
}

__END__

Original decompiled (by hand) code is here:
After studying it for a while, I can see that there are effectively 2 nested loops
counting from 1 up to $c. If the 2 loop values multiply to become $c, then one
of them is added to $b, which is returned at the end.

So whats actually happening is that all the factors of $c are being added up.
Hence the above code to do that in a much more efficient way.
The numbers for $c for part 1 and 2 are found by running the original program
and dumping the instructions and registers for each step.


my $a = 0;
my $c = 976;

my $b = 1;
my $f = 1;

while (1) {
	if ($b * $f == $c) {
		$a = $a + $b;
	}

	$f++;

	if ($f > $c) {
		$b++;
		if ($b > $c) {
			last;
		}
		else {
			$f = 1;
		}
	}
}

print $a . "\n";
