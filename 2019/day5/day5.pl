#!/usr/bin/perl
use strict;
use warnings;

use Term::ReadKey;

ReadMode 'cbreak';

END {
	ReadMode 'normal';
}

my $pos;
my @p;
my $exit = 0;

my %commands = (
	1 => {
		"00" => sub { $p[$p[$pos+3]] = $p[$p[$pos+1]] + $p[$p[$pos+2]] },
		"01" => sub { $p[$p[$pos+3]] = $p[$pos+1] + $p[$p[$pos+2]] },
		"10" => sub { $p[$p[$pos+3]] = $p[$p[$pos+1]] + $p[$pos+2] },
		"11" => sub { $p[$p[$pos+3]] = $p[$pos+1] + $p[$pos+2] },
		"inc" => 4
	},
	2 => {
		"00" => sub { $p[$p[$pos+3]] = $p[$p[$pos+1]] * $p[$p[$pos+2]] },
		"01" => sub { $p[$p[$pos+3]] = $p[$pos+1] * $p[$p[$pos+2]] },
		"10" => sub { $p[$p[$pos+3]] = $p[$p[$pos+1]] * $p[$pos+2] },
		"11" => sub { $p[$p[$pos+3]] = $p[$pos+1] * $p[$pos+2] },
		"inc" => 4
	},
	3 => {
		"00" => sub { $p[$p[$pos+1]] = ReadKey(0) },
		"inc" => 2
	},
	4 => {
		"00" => sub { print $p[$p[$pos+1]] },
		"01" => sub { print $p[$pos+1] },
		"inc" => 2
	},
	99 => {
		"00" => sub { $exit = 1 }
	}
);

sub run_program {
	$pos = 0;
	$exit = 0;

	while (1) {
		my $ins = $p[$pos];
		my $opcode = substr($ins, length($ins) - 2, 2);
		$opcode = sprintf("%d", $opcode);
		my $modes = substr($ins, 0, length($ins) - 2) || "00";
		$modes = sprintf("%02d", $modes);

		$commands{$opcode}{$modes}->();
		last if $exit == 1;

		$pos += $commands{$opcode}{inc};
	}
	return $p[0];
}

open my $fh, '<', 'input.txt';
chomp(my $line = <$fh>);
close $fh;
@p = split(/,/,$line);
run_program();

