package Plants;
use strict;
use warnings;
no warnings 'portable';
use Math::BigInt;

use parent "Exporter";
our @EXPORT_OK = qw(iterate);

sub iterate {
	my ($initial,@data) = @_;

	$initial =~ s/#/1/g;
	$initial =~ s/\./0/g;
	$initial = Math::BigInt->from_bin($initial); 

	my @commands;
	my @change;
	foreach my $row (@data) {
		$row =~ /(.*) => (.)/;
		my ($pattern,$set) = ($1,$2);

		$pattern =~ s/#/1/g;
		$pattern =~ s/\./0/g;
		$pattern = oct("0b" . $pattern);

		$set = $set eq '#' ? 1 : 0;

		push @commands, $pattern;
		push @change, $set;
	}

	my $state = $initial;

	my $shift = 0;

	for (1..20) {
		my $next = Math::BigInt->new(0);
		my $next_shift = 0;

		$state->blsft(4);

		my $leading_zeroes = 0;
		while ($state > 0) {
			my $match = 0;
			my $chunk = $state & 31;
			for (my $i = 0; $i < @commands; $i++) {
				my $command = $commands[$i];
				if ($chunk == $command) {
					if ($change[$i] == 0) {
						$leading_zeroes++;
					}
					else {
						$next += Math::BigInt->new(1)->blsft($next_shift);
						$leading_zeroes = 0;
					}
					$match = 1;
					last;
				}
			}
			if ($match == 0) {
				$leading_zeroes++;
			}
			$next_shift++;

			$state->brsft(1);
		}

		$shift -= 2 - $leading_zeroes;

		$state = $next;
		while (($state & 1) == 0) {
			$state->brsft(1)
		}
	}

	my $sum = 0;
	my $bin = $state->as_bin;
	$bin =~ s/0b//;
	foreach my $digit (split//,$bin) {
		$sum += $shift if $digit == 1;
		$shift++;
	}

	return $sum;
}
