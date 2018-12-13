package Plants;
use strict;
use warnings;

use parent "Exporter";
our @EXPORT_OK = qw(iterate);

sub iterate {
	my ($initial,$generations,@data) = @_;

	my @state = split//,$initial;

	my @commands;
	my @change;
	foreach my $row (@data) {
		$row =~ /(.*) => (.)/;
		my ($pattern,$set) = ($1,$2);
		next if $set eq '.';
		push @commands, $pattern;
	}

	my $left_add = 0;

	for (1..$generations) {
		my @previous_state = @state;
		my $previous_left_add = $left_add;
		@state = ('.','.','.','.',@state,'.','.','.','.');

		$left_add += 2;

		my @new_state;
		for (my $i = 0; $i < @state - 5; $i++) {
			my $chunk = join('',@state[$i..$i+4]);
			my $match = 0;
			foreach my $command (@commands) {
				if ($chunk eq $command) {
					push @new_state, '#';
					$match = 1;
					last;
				}
			}
			push @new_state, '.' if $match == 0;
		}

		while ($new_state[0] eq '.') {
			shift(@new_state);
			$left_add--;
		}

		while ($new_state[-1] eq '.') {
			pop(@new_state);
		}

		@state = @new_state;
		if (join('',@state) eq join('',@previous_state)) {
			my $change = $left_add - $previous_left_add;
			$left_add += $change * ($generations - $_);
			last;
		}
	}

	my $sum = 0;
	my $value = -$left_add;
	foreach my $plant (@state) {
		$sum += $value if $plant eq '#';
		$value++;
	}

	return $sum;
}
