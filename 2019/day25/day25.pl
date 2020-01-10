#!/usr/bin/perl
use strict;
use warnings;
use Math::BigInt;
use Algorithm::Combinatorics qw(combinations);

my $pos = 0;
my $rb = 0;
my @output;

sub run_program {
	my ($program, $input) = @_;

	while (1) {
		my $ins = $program->[$pos];

		my $opcode = sprintf("%d", substr($ins, length($ins) - 2, 2));

		if ($opcode == 99) {
			return -9;
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
			if (@$input) {
				$program->[$a1] = shift(@$input);
			}
			else {
				return;
			}
			$pos += 2;
		}
		elsif ($opcode == 4) {
			$pos += 2;
			print chr($a1);
			#return $a1;
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
}

open my $fh, '<', 'input.txt';
chomp(my $line = <$fh>);
close $fh;
my @program = map { Math::BigInt->new($_) } split(/,/,$line);

# get all the stuff
my @commands = (
	'west','west','west','take coin',
	'east','east','east','north','north','take mutex',
	'east','take antenna',
	'west','south','east','take cake',
	'east','north','take pointer',
	'south','west','west','south','east','east','take tambourine',
	'east','take fuel cell',
	'east','take boulder',
	'north'
);

my $input = [map { ord($_) } split(//,join("\n", @commands) . "\n")];
run_program(\@program, $input) while @$input;

my @stuff = (
	'fuel cell',
	'cake',
	'pointer',
	'boulder',
	'mutex',
	'antenna',
	'tambourine',
	'coin'
);

# drop all the stuff
$input = [map { ord($_) } split(//,join("\n",map { "drop $_" } @stuff) . "\n")];
run_program(\@program, $input) while @$input;

# try all the combinations
# (this could be improved by only dropping stuff you don't need for the next round)
for (my $k = 1; $k <= @stuff; $k++) {
	my $iter = combinations(\@stuff,$k);
	while (my $p = $iter->next) {
		$input = [map { ord($_) } split(//,join("\n",map { "take $_" } @$p) . "\n")];
		run_program(\@program, $input) while @$input;
		$input = [map { ord($_) } split(//,"east\n")];
		run_program(\@program, $input) while @$input;
		$input = [map { ord($_) } split(//,join("\n",map { "drop $_" } @$p) . "\n")];
		run_program(\@program, $input) while @$input;
	}
}

while (1) {
	run_program(\@program, $input);
	my $in = <STDIN>;
	$input = [map { ord($_) } split//,$in];
}
