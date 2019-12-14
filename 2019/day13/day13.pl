#!/usr/bin/perl
use strict;
use warnings;
use Math::BigInt;
use Term::ANSIScreen qw(cls locate);

my $pos = 0;
my $rb = 0;

my ($lx, $hx, $ly, $hy) = (0, 0, 0, 0);

sub display {
	my ($grid, $score) = @_;
	locate 1,1;
	for my $y ($ly..$hy) {
		for my $x ($lx..$hx) {
			if (exists($grid->{$x}{$y})) {
				if ($grid->{$x}{$y} == 0) {
					print ' ';
				}
				elsif ($grid->{$x}{$y} == 1) {
					print '#';
				}
				elsif ($grid->{$x}{$y} == 2) {
					print '@';
				}
				elsif ($grid->{$x}{$y} == 3) {
					print '=';
				}
				elsif ($grid->{$x}{$y} == 4) {
					print 'o';
				}
			}
			else {
				print ' ';
			}
		}
		print "\n";
	}
	print "$score\n";
}

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
			if (@output == 3) {
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
	return [-9];
}

open my $fh, '<', 'input.txt';
chomp(my $line = <$fh>);
close $fh;
my @source_program = split(/,/,$line);

# part 1
my @program = map { Math::BigInt->new($_) } @source_program;

my %grid;
while (1) {
	my $output = run_program(\@program);
	last if $output->[0] == -9;

	my ($x,$y,$val) = @$output;
	
	$grid{$x}{$y} = $val;

	$lx = $x if $x < $lx;
	$hx = $x if $x > $hx;
	$ly = $y if $y < $ly;
	$hy = $y if $y > $ly;	
}

my $count = 0;
foreach my $x (keys %grid) {
	foreach my $y (keys %{$grid{$x}}) {
		if ($grid{$x}{$y} == 2) {
			$count++;
		}
	}
}
print "$count\n";

# part 2
@program = map { Math::BigInt->new($_) } @source_program;
%grid = ();

$program[0] = Math::BigInt->new(2);

my $input = 0;
my $score = 0;
my $paddle;
$pos = 0;
$rb = 0;
cls();
while (1) {
	my $output = run_program(\@program, $input);
	last if $output->[0] == -9;
	my ($x,$y,$val) = @$output;	

	if ($x == -1 && $y == 0) {
		$score = $val;
		next;
	}
	
	if ($grid{$x}{$y} || 0 != $val) {
		$grid{$x}{$y} = $val;
	}

	# paddle
	if ($val == 3) {
		$paddle = $x;
	}

	# ball
	if ($val == 4 && $paddle) {
		display(\%grid, $score);
		if ($paddle < $x) {
			$input = 1;
		}
		elsif ($paddle > $x) {
			$input = -1;
		}
		else {
			$input = 0;
		}
	}

}
print "$score\n";
