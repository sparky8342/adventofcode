#!/usr/bin/perl
use strict;
use warnings;
use Math::BigInt;

my $pos = 0;
my $rb = 0;

sub run_program {
	my ($program) = @_;

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
			#$program->[$a1] = $input;
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
	return [-1];
}

open my $fh, '<', 'input.txt';
chomp(my $line = <$fh>);
close $fh;
my @source_program = split(/,/,$line);

my ($lx, $hx, $ly, $hy) = (0, 0, 0, 0);

my @program = map { Math::BigInt->new($_) } @source_program;

my %grid;

while (1) {
	my $output = run_program(\@program);
	last if $output->[0] == -1;
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

for my $y ($ly..$hy) {
	for my $x ($lx..$hx) {
		if (exists($grid{$x}{$y})) {
			if ($grid{$x}{$y} == 0) {
				print ' ';
			}
			elsif ($grid{$x}{$y} == 1) {
				print '#';
			}
			elsif ($grid{$x}{$y} == 2) {
				print '@';
			}
			elsif ($grid{$x}{$y} == 3) {
				print '=';
			}
			elsif ($grid{$x}{$y} == 4) {
				print 'o';
			}
		}
		else {
			print ' ';
		}
	}
	print "\n";
}


__END__

The software draws tiles to the screen with output instructions: every three output instructions specify the x position (distance from the left), y position (distance from the top), and tile id. The tile id is interpreted as follows:

    0 is an empty tile. No game object appears in this tile.
    1 is a wall tile. Walls are indestructible barriers.
    2 is a block tile. Blocks can be broken by the ball.
    3 is a horizontal paddle tile. The paddle is indestructible.
    4 is a ball tile. The ball moves diagonally and bounces off objects.

