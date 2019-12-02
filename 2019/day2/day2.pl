#!/usr/bin/perl
use strict;
use warnings;

sub run_program {
	my ($program) = @_;

	my $pos = 0;

	while (1) {
		my $opcode = $program->[$pos];
		if ($opcode == 99) {
			last;
		}
		my @args = @$program[$pos+1..$pos+3];
		if ($opcode == 1) {
			$program->[$args[2]] = $program->[$args[0]] + $program->[$args[1]];
		}
		elsif ($opcode == 2) {
			$program->[$args[2]] = $program->[$args[0]] * $program->[$args[1]];
		}
		$pos += 4;
	}
	return $program->[0];
}

open my $fh, '<', 'input.txt';
chomp(my $line = <$fh>);
close $fh;
my @source_program = split(/,/,$line);

# part 1
my @program = @source_program;
$program[1] = 12;
$program[2] = 2;
print run_program(\@program) . "\n";

# part 2
my $target_result = 19690720;
for my $noun (0..99) {
	for my $verb (0..99) {
		my @program = @source_program;
		$program[1] = $noun;
		$program[2] = $verb;
		my $result = run_program(\@program);
		if ($result == $target_result) {
			print 100 * $noun + $verb . "\n";
			exit;
		}
	}
}
