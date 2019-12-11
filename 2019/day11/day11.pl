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

my @program = map { Math::BigInt->new($_) } @source_program;

my %grid = (0 => { 0 => 0 });

my $robot = { x => 0, y => 0, dir => '^' };
my $input = 0;

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
}

my $painted = 0;
foreach my $x (keys %grid) {
	foreach my $y ( keys %{$grid{$x}} ) {
		$painted++;
	}
}
print "$painted\n";
