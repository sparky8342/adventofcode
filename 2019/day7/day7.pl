#!/usr/bin/perl
use strict;
use warnings;

use Algorithm::Combinatorics qw(permutations);

my $pos;
my $p = [];
my $exit = 0;
my @input;
my @output;
my $waiting = 0;

my %commands = (
	1 => {
		"00" => sub { $p->[$p->[$pos+3]] = $p->[$p->[$pos+1]] + $p->[$p->[$pos+2]] },
		"01" => sub { $p->[$p->[$pos+3]] = $p->[$pos+1] + $p->[$p->[$pos+2]] },
		"10" => sub { $p->[$p->[$pos+3]] = $p->[$p->[$pos+1]] + $p->[$pos+2] },
		"11" => sub { $p->[$p->[$pos+3]] = $p->[$pos+1] + $p->[$pos+2] },
		"inc" => 4
	},
	2 => {
		"00" => sub { $p->[$p->[$pos+3]] = $p->[$p->[$pos+1]] * $p->[$p->[$pos+2]] },
		"01" => sub { $p->[$p->[$pos+3]] = $p->[$pos+1] * $p->[$p->[$pos+2]] },
		"10" => sub { $p->[$p->[$pos+3]] = $p->[$p->[$pos+1]] * $p->[$pos+2] },
		"11" => sub { $p->[$p->[$pos+3]] = $p->[$pos+1] * $p->[$pos+2] },
		"inc" => 4
	},
	3 => {
		"00" => sub { if (@input == 0) { $waiting = 1 } else { $p->[$p->[$pos+1]] = shift(@input) } },
		"inc" => 2
	},
	4 => {
		"00" => sub { push @output, $p->[$p->[$pos+1]] },
		"01" => sub { push @output, $p->[$pos+1] },
		"inc" => 2
	},
	5 => {
		"00" => sub { if ($p->[$p->[$pos+1]] != 0) { $pos = $p->[$p->[$pos+2]] } else { $pos += 3 } },
		"01" => sub { if ($p->[$pos+1] != 0) { $pos = $p->[$p->[$pos+2]] } else { $pos += 3 } },
		"10" => sub { if ($p->[$p->[$pos+1]] != 0) { $pos = $p->[$pos+2] } else { $pos += 3 } },
		"11" => sub { if ($p->[$pos+1] != 0) { $pos = $p->[$pos+2] } else { $pos += 3 } },
		"inc" => 0
	},
	6 => {
		"00" => sub { if ($p->[$p->[$pos+1]] == 0) { $pos = $p->[$p->[$pos+2]] } else { $pos += 3 }},
		"01" => sub { if ($p->[$pos+1] == 0) { $pos = $p->[$p->[$pos+2]] } else { $pos += 3 } },
		"10" => sub { if ($p->[$p->[$pos+1]] == 0) { $pos = $p->[$pos+2] } else { $pos += 3} },
		"11" => sub { if ($p->[$pos+1] == 0) { $pos = $p->[$pos+2] } else { $pos += 3 } },
		"inc" => 0
	},
	7 => {
		"00" => sub { $p->[$p->[$pos+3]] = $p->[$p->[$pos+1]] < $p->[$p->[$pos+2]] ? 1 : 0 },
		"01" => sub { $p->[$p->[$pos+3]] = $p->[$pos+1] < $p->[$p->[$pos+2]] ? 1 : 0 },
		"10" => sub { $p->[$p->[$pos+3]] = $p->[$p->[$pos+1]] < $p->[$pos+2] ? 1 : 0 },
		"11" => sub { $p->[$p->[$pos+3]] = $p->[$pos+1] < $p->[$pos+2] ? 1 : 0 },
		"inc" => 4
	},
	8 => {
		"00" => sub { $p->[$p->[$pos+3]] = $p->[$p->[$pos+1]] == $p->[$p->[$pos+2]] ? 1 : 0 },
		"01" => sub { $p->[$p->[$pos+3]] = $p->[$pos+1] == $p->[$p->[$pos+2]] ? 1 : 0 },
		"10" => sub { $p->[$p->[$pos+3]] = $p->[$p->[$pos+1]] == $p->[$pos+2] ? 1 : 0 },
		"11" => sub { $p->[$p->[$pos+3]] = $p->[$pos+1] == $p->[$pos+2] ? 1 : 0 },
		"inc" => 4
	},
	99 => {
		"00" => sub { $exit = 1 }
	}
);

sub run_program {
	$exit = 0;
	$waiting = 0;

	while (1) {
		my $ins = $p->[$pos];
		my $opcode = substr($ins, length($ins) - 2, 2);
		$opcode = sprintf("%d", $opcode);
		my $modes = substr($ins, 0, length($ins) - 2) || "00";
		$modes = sprintf("%02d", $modes);

		$commands{$opcode}{$modes}->();
		last if $exit == 1;

		# stop executing when waiting for more input
		if ($waiting == 1) {
			last;
		}

		$pos += $commands{$opcode}{inc};
		
	}
	return $p->[0];
}

open my $fh, '<', 'input.txt';
chomp(my $line = <$fh>);
close $fh;
my @prog = split(/,/,$line);

for my $part (0..1) {
	my $iter;
	if ($part == 0) {
		$iter = permutations([0,1,2,3,4], 5);
	}
	elsif ($part == 1) {
		$iter = permutations([5,6,7,8,9], 5);
	}

	my $highest = 0;
	while (my $phases = $iter->next) {

		my @amps = ([@prog], [@prog], [@prog], [@prog], [@prog]);
		my @amp_inputs = map { [$_] } @$phases;
		my @amp_pos = (0,0,0,0,0);

		push @{$amp_inputs[0]}, 0;

		$exit = 0;
		while ($exit == 0) {
			for (my $i = 0; $i < 5; $i++) {
				@input = @{$amp_inputs[$i]};
				$amp_inputs[$i] = [];
				$p = $amps[$i];
				@output = ();
				$pos = $amp_pos[$i];
				run_program();
				$amp_pos[$i] = $pos;

				my $next = $i + 1;
				$next = 0 if $next == 5;
				push @{$amp_inputs[$next]}, @output;
			}
		}

		my $value = join('', @output);
		$highest = $value if $value > $highest;

	}
	print "$highest\n";
}
