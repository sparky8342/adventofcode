#!/usr/bin/perl
use strict;
use warnings;

my %replacements;

open my $fh, '<', 'input.txt';
while (my $line = <$fh>) {
	chomp($line);
	last if $line eq '';
	my ($val, $repl) = split(' => ', $line);
	push @{$replacements{$val}}, $repl;
}
my $molecule = <$fh>;
close $fh;

my %mo;

foreach my $val (keys %replacements) {
	LOOP:
	foreach my $repl ( @{$replacements{$val}} ) {
		my $n = 0;

		while (1) {
			my $m = $molecule;
			for (1..$n) {
				$m =~ s/$val/-/;
			}
			my $found = $m =~ s/$val/$repl/;
			next LOOP unless $found;
			$m =~ s/-/$val/g;
			$mo{$m} = 1;
			$n++;
		}
	}
}

print scalar keys %mo;
print "\n";
