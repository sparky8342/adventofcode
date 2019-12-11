#!/usr/bin/perl
use strict;
use warnings;
use Math::BigInt;

use constant TURNS => {
	0 => ['^', '<', 'v', '>', '^'],
	1 => ['^', '>', 'v', '<', '^']
};

my $pos = 0;
my $rb = 0;

sub run_program {
	my ($program, $input) = @_;

	my @output;

	while (1) {
		my $ins = $program->[$pos];

                my $opcode = sprintf("%d", substr($ins, length($ins) - 2, 2));

		if ($opcode == 99) {
			last;
		}

		my ($mode3, $mode2, $mode1) = split(//, sprintf("%03d", substr($ins, 0, length($ins) - 2) || "000"));

		my ($a1, $a2, $a3) = @$program[$pos+1..$pos+3];
		if ($opcode != 3) {
			if ($mode1 == 0) {
				$a1 = $program->[$a1] || 0;
			}
			elsif ($mode1 == 2) {
				$a1 = $program->[$a1 + $rb] || 0;
			}
		}
		if ($mode2 == 0) {
			$a2 = $program->[$a2] || 0;
		}
		elsif ($mode2 == 2) {
			$a2 = $program->[$a2 + $rb] || 0;
		}
		if ($mode3 == 2) {
			$a3 += $rb;
		}

		if ($opcode == 1) {
			$program->[$a3] = $a1 + $a2;
			$pos += 4;
		}
		elsif ($opcode == 2) {
			$program->[$a3] = $a1 * $a2;
			$pos += 4;
		}
		elsif ($opcode == 3) {
			$a1 += $rb if $mode1 == 2;
			$program->[$a1] = $input;
			$pos += 2;
		}
		elsif ($opcode == 4) {
			push @output, $a1;
			$pos += 2;
			if (@output == 2) {
				my @out = @output;
				@output = ();
				return \@out;
			}
		}
		elsif ($opcode == 5) {
			$pos = $a1 != 0 ? $a2 : $pos + 3;
		}
		elsif ($opcode == 6) {
			$pos = $a1 == 0 ? $a2 : $pos + 3;
		}
		elsif ($opcode == 7) {
			$program->[$a3] = $a1 < $a2 ? 1 : 0;
			$pos += 4;
		}
		elsif ($opcode == 8) {
			$program->[$a3] = $a1 == $a2 ? 1 : 0;
			$pos += 4;
		}
		elsif ($opcode == 9) {
			$rb += $a1;
			$pos += 2;
		}
	}
	return [-1];
}

open my $fh, '<', 'input.txt';
chomp(my $line = <$fh>);
close $fh;
my @source_program = split(/,/,$line);

my $limits = { lx => 0, hx => 0, ly => 0, hy => 0 };

for my $part (1..2) {
	my @program = map { Math::BigInt->new($_) } @source_program;
	$pos = 0;
	$rb = 0;

	my %grid = (0 => { 0 => 0 });
	if ($part == 2) {
		$grid{0}{0} = 1;
	}

	my $robot = { x => 0, y => 0, dir => '^' };

	while (1) {
		my $output = run_program(\@program, $grid{$robot->{x}}{$robot->{y}});
		last if $output->[0] == -1;

		$grid{$robot->{x}}{$robot->{y}} = $output->[0];
		my $turn = $output->[1];
		for (my $i = 0; $i <= 5; $i++) {
			if (TURNS->{$turn}->[$i] eq $robot->{dir}) {
				$robot->{dir} = TURNS->{$turn}->[$i + 1];
				last;
			}
		}
		if    ($robot->{dir} eq '^') { $robot->{y}-- }
		elsif ($robot->{dir} eq '>') { $robot->{x}++ }
		elsif ($robot->{dir} eq 'v') { $robot->{y}++ }
		elsif ($robot->{dir} eq '<') { $robot->{x}-- }

		if ($part == 2) {
			if ($robot->{x} < $limits->{lx}) { $limits->{lx} = $robot->{x} }
			if ($robot->{x} > $limits->{hx}) { $limits->{hx} = $robot->{x} }
			if ($robot->{y} < $limits->{ly}) { $limits->{ly} = $robot->{y} }
			if ($robot->{y} > $limits->{hy}) { $limits->{hy} = $robot->{y} }
		}
	}

	if ($part == 1) {
		my $painted = 0;
		foreach my $x (keys %grid) {
			foreach my $y ( keys %{$grid{$x}} ) {
				$painted++;
			}
		}
		print "$painted\n";
	}
	elsif ($part == 2) {
		for my $y ($limits->{ly}..$limits->{hy}) {
			for my $x ($limits->{lx}..$limits->{hx}) {
				if (exists($grid{$x}{$y})) {
					if ($grid{$x}{$y} == 1) {
						print '#';
					}
					else {
						print ' ';
					}
				}
				else {
					print ' ';
				}
			}
			print "\n";
		}
	}
}
