#!/usr/bin/perl
use strict;
use warnings;

my @replacements;

open my $fh, '<', 'input.txt';
while (my $line = <$fh>) {
	chomp($line);
	last if $line eq '';
	my ($val, $repl) = split(' => ', $line);
	push @replacements, [ $val, $repl ];
}
chomp(my $molecule = <$fh>);
close $fh;

my %mo;

LOOP:
foreach my $r (@replacements) {
	my ($val, $repl) = @$r;
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

print scalar keys %mo;
print "\n";

no warnings 'recursion';
my $steps = search($molecule, 0);
print "$steps\n";

sub search {
	my ($m, $steps) = @_;

	if ($m eq 'e') {
		return $steps;
	}

	SEARCH:
	foreach my $r (@replacements) {
		my ($val, $repl) = @$r;

		if ($m =~ /$repl/) {	
			my $n = 0;
			while (1) {
				my $mo = $m;
				for (1..$n) {
					$mo =~ s/$repl/-/;
				}
				my $found = $mo =~ s/$repl/$val/;
				next SEARCH unless $found;
				$mo =~ s/-/$repl/g;
				if (my $steps = search($mo, $steps + 1)) {
					return $steps;
				}
				$n++;
			}
		}
	}
}
