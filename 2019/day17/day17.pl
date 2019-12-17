#!/usr/bin/perl
use strict;
use warnings;
use Math::BigInt;
use Term::ANSIScreen qw(cls locate);

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
			$program->[$a1] = shift(@$input);
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

my %grid = ();

my ($x, $y) = (0, 0);
my ($hx, $hy) = (0, 0);

while (1) {
	my $output = run_program(\@program);
	if ($output == -9) {
		$hy = $y - 2;
		last;
	}
	if ($output == 10) {
		$hx = $x - 1 unless $hx;
		$y++;
		$x = 0;
	}
	else {
		my $char = chr($output);
		$grid{$x}{$y} = $char;
		$x++;
	}
}
print_grid();

my $sum = 0;
my $robot;
for my $y (0..$hy) {
	for my $x (0..$hx) {
		if ($grid{$x}{$y} eq '#') {
			if (($grid{$x-1}{$y} || '') eq '#'
			&&  ($grid{$x+1}{$y} || '') eq '#'
			&&  ($grid{$x}{$y-1} || '') eq '#'
			&&  ($grid{$x}{$y+1} || '') eq '#') {
				$sum += $x * $y;
			}
		}
		elsif ($grid{$x}{$y} =~ /^([<>^v])$/) {
			$robot = { dir => $1, x => $x, y => $y };
		}
	}
}
print "$sum\n";

# part 2
my @dirs = (
	{
		x => -1,
		y => 0,
		dir => {
			'^' => { new_dir => '<', turn => 'L' },
			'v' => { new_dir => '<', turn => 'R' }
		}
	},
	{
		x => 1,
		y => 0,
		dir => {
			'^' => { new_dir => '>', turn => 'R' },
			'v' => { new_dir => '>', turn => 'L' }
		}
	},
	{
		x => 0,
		y => -1,
		dir => {
			'<' => { new_dir => '^', turn => 'R' },
			'>' => { new_dir => '^', turn => 'L' }
		}
	},
	{
		x => 0,
		y => 1,
		dir => {
			'<' => { new_dir => 'v', turn => 'L' },
			'>' => { new_dir => 'v', turn => 'R' }
		}
	}
);

my @commands;
while (1) {
	my $move;
	my $found = 0;
	foreach my $dir (@dirs) {
		if (($grid{$robot->{x} + $dir->{x}}{$robot->{y} + $dir->{y}} || '') eq '#'
		&& exists($dir->{dir}->{$robot->{dir}})) {
			$move = { x => $dir->{x}, y => $dir->{y}, turn => $dir->{dir}->{$robot->{dir}}->{turn}, steps => 0 };
			$robot->{dir} = $dir->{dir}->{$robot->{dir}}->{new_dir};
			$found = 1;
			last;
		}
	}
	last if !$found;

	while (($grid{$robot->{x} + $move->{x}}{$robot->{y} + $move->{y}} || '') eq '#') {
		$robot->{x} += $move->{x};
		$robot->{y} += $move->{y};
		$move->{steps}++;
	}

	push @commands, $move->{turn} . $move->{steps};
}

my $str = join(',',@commands);
$str =~ s/,//g;

my %groups;
foreach my $group ('A', 'B', 'C') {
	if ($str =~ /([^A-C]+).*?\1/) {
		$groups{$group} = $1;
		$str =~ s/$groups{$group}/$group/g;
		$groups{$group} =~ s/(\d)(\D)/$1,$2/g;
		$groups{$group} =~ s/(\D)(\d)/$1,$2/g;
	}
}
$str = join(',',split(//,$str));

my @input;
foreach my $line ($str, $groups{A}, $groups{B}, $groups{C}) {
	push @input, map { ord($_) } split(//,$line);
	push @input, 10;
}
push @input, ord('n');
push @input, 10;

@program = map { Math::BigInt->new($_) } @source_program;
$program[0] = Math::BigInt->new(2);
$pos = 0;
$rb = 0;

my $last_out;
while (1) {
	my $output = run_program(\@program, \@input);
	last if $output == -9;
	$last_out = $output;
}
print "$last_out\n";

sub print_grid {
	#locate(1,1);
	for my $y (0..$hy) {
		for my $x (0..$hx) {
			print $grid{$x}{$y} || ' ';
		}
		print "\n";
	}
}
