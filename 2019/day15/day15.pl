#!/usr/bin/perl
use strict;
use warnings;
use Math::BigInt;
use Term::ANSIScreen qw(cls locate);

my $pos = 0;
my $rb = 0;

my ($lx, $hx, $ly, $hy) = (0, 0, 0, 0);

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
			$pos += 2;
			return $a1;
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
	return -9;
}

open my $fh, '<', 'input.txt';
chomp(my $line = <$fh>);
close $fh;
my @source_program = split(/,/,$line);

# part 1
my @program = map { Math::BigInt->new($_) } @source_program;

my %grid = (0 => { 0 => '.' });
my ($x, $y) = (0, 0);
my ($ox, $oy);
cls();
while (1) {
	my ($dx, $dy) = (0, 0);
	my $input = bfs($x,$y,1);

	if    ($input == 1) { $dy = -1 }
	elsif ($input == 2) { $dy =  1 }
	elsif ($input == 3) { $dx = -1 }
	elsif ($input == 4) { $dx =  1 }
	$dx += $x;
	$dy += $y;

	my $output = run_program(\@program,$input);
	last if $output == -9;

	if ($output == 0) {
		$grid{$dx}{$dy} = '#';
	}
	elsif ($output == 1) {
		$grid{$dx}{$dy} = '.';
		($x, $y) = ($dx, $dy);
	}
	elsif ($output == 2) {
		$grid{$dx}{$dy} = 'o';
		($x, $y) = ($dx, $dy);
		($ox, $oy) = ($dx, $dy);
	}

	print_grid($x,$y);

	$lx = $dx if $dx < $lx;
	$hx = $dx if $dx > $hx;
	$ly = $dy if $dy < $ly;
	$hy = $dy if $dy > $hy;	
}
my $steps = bfs(0,0,2,$ox,$oy);
print "$steps\n";

# part 2
my $longest = bfs($ox, $oy, 3);
print "$longest\n";

# 311 too high

sub print_grid {
	locate(1,1);
	my ($rx, $ry) = @_;
	for my $y ($ly..$hy) {
		for my $x ($lx..$hx) {
			if ($x == $rx && $y == $ry) {
				print 'D';
			}
			else {
				print $grid{$x}{$y} || ' ';
			}
		}
		print "\n";
	}
}

sub bfs {
	my ($x, $y, $mode, $x2, $y2) = @_;

	my @queue = { x => $x, y => $y, dist => 0 };
	my %visited;
	my $max_dist = 0;

	while (@queue) {
		my $space = shift(@queue);
		if ($mode == 1) {
			# nearest unknown space
			if (!exists($grid{$space->{x}}{$space->{y}})) {
				return $space->{dir};
			}
		}
		elsif ($mode == 2) {
			# find path to x2, y2
			if ($space->{x} == $x2 && $space->{y} == $y2) {
				return $space->{dist};
			}
		}

		if (exists($visited{$space->{x}}{$space->{y}})) {
			next;
		}
		$visited{$space->{x}}{$space->{y}} = 1;
		if ($grid{$space->{x}}{$space->{y}} eq '#') {
			next;
		}

		if ($mode == 3) {
			# find longest path
			$max_dist = $space->{dist} if $space->{dist} > $max_dist;
		}

		push @queue, (
			{ x => $space->{x} + 1, y => $space->{y},     dist => $space->{dist} + 1, dir => $space->{dir} || 4 },
			{ x => $space->{x} - 1, y => $space->{y},     dist => $space->{dist} + 1, dir => $space->{dir} || 3 },
			{ x => $space->{x}    , y => $space->{y} + 1, dist => $space->{dist} + 1, dir => $space->{dir} || 2 },
			{ x => $space->{x}    , y => $space->{y} - 1, dist => $space->{dist} + 1, dir => $space->{dir} || 1 }
		);
	}
	return $max_dist if $mode == 3;
}	
