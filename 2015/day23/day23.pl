#!/usr/bin/perl
use strict;
use warnings;

open my $fh, '<', 'input.txt';
chomp(my @program = <$fh>);
close $fh;

my $registers = run_program(\@program, { a => 0, b => 0 });
print $registers->{b} . "\n";
$registers = run_program(\@program, { a => 1, b => 0 });
print $registers->{b} . "\n";

sub run_program {
	my ($program, $registers) = @_;
	my $pos = 0;

	while ($pos >= 0 && $pos < scalar(@$program)) {
		my $ins = $program->[$pos];

		if ($ins =~ /^hlf (.)$/) {
			$registers->{$1} /= 2;
			$pos++;
		}
		elsif ($ins =~ /^tpl (.)$/) {
			$registers->{$1} *= 3;
			$pos++;
		} 
		elsif ($ins =~ /^inc (.)$/) {
			$registers->{$1}++;
			$pos++;
		}
		elsif ($ins =~ /^jmp (.+)$/) {
			$pos += $1;
		} 
		elsif ($ins =~/^jie (.), (.+)$/) {
			if ($registers->{$1} % 2 == 0) {
				$pos += $2;
			}
			else {
				$pos++;
			}
		}
		elsif ($ins =~/^jio (.), (.+)$/) {
			if ($registers->{$1} == 1) {
				$pos += $2;
			}
			else {
				$pos++;
			}
		}
	}
	return $registers;
}
