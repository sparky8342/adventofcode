package MineCarts;
use strict;
use warnings;

use parent "Exporter";
our @EXPORT_OK = qw(move);

use constant CARTS => {
	'>' => { drow => 0, dcol => 1 },
	'<' => { drow => 0, dcol => -1 },
	'^' => { drow => -1, dcol => 0 },
	'v' => { drow => 1, dcol => 0 },
};

use constant DIR_CHANGE => {
	'>' => {
		'/' => '^',
		'\\' => 'v',
		'+' => [ '^', '>', 'v' ]
		},
	'<' => {
		'/' => 'v',
		'\\' => '^',
		'+' => [ 'v', '<', '^' ]
		},
	'^' => {
		'/' => '>',
		'\\' => '<',
		'+' => [ '<', '^', '>' ]
		},
	'v' => {
		'/' => '<',
		'\\' => '>',
		'+' => [ '>', 'v', '<' ]
		}
};

sub move {
	my (@data) = @_;

	my @grid;
	foreach my $row (@data) {
		push @grid, [split//,$row];
	}

	my @carts;

	for (my $row = 0; $row < @grid; $row++) {
		for (my $col = 0; $col < @{$grid[$row]}; $col++) {
			my $square = $grid[$row][$col];
			if (exists(&CARTS->{$square})) {
				push @carts, {
					dir => $square,
					row => $row,
					col => $col,
					turn => 0
				};
			}
		}
	}

	while (1) {
		@carts = sort { $a->{row} <=> $b->{row} || $a->{col} <=> $b->{col} } @carts;

		foreach my $cart (@carts) {
			$cart->{row} += &CARTS->{$cart->{dir}}->{drow};
			$cart->{col} += &CARTS->{$cart->{dir}}->{dcol};
			my $square = $grid[$cart->{row}][$cart->{col}];
			if ($square eq '/' || $square eq '\\') {
				$cart->{dir} = &DIR_CHANGE->{$cart->{dir}}->{$square};
			}
			elsif ($square eq '+') {
				$cart->{dir} = &DIR_CHANGE->{$cart->{dir}}->{'+'}->[$cart->{turn}];
				$cart->{turn}++;
				$cart->{turn} = 0 if $cart->{turn} == 3;
			}

			my $collision = collision_check(@carts);
			return $collision if $collision;
		}

	}

}

sub collision_check {
	my @carts = @_;
	@carts = sort { $a->{row} <=> $b->{row} || $a->{col} <=> $b->{col} } @carts;

	for (my $i = 0; $i < @carts - 1; $i++) {
		if ($carts[$i]->{row} == $carts[$i+1]->{row} && $carts[$i]->{col} == $carts[$i+1]->{col}) {
			return $carts[$i]->{col} . ',' . $carts[$i]->{row};
		}
	}

	return undef;
}
