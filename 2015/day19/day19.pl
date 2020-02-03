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

my %seen;
my $smallest_steps = 999999;
no warnings 'recursion';
search($molecule, 0);
print "$smallest_steps\n";

sub search {
	my ($m, $steps) = @_;

	if ($m eq 'e') {
		if ($steps < $smallest_steps) {
			print "$steps\n";
			$smallest_steps = $steps;
		}
		return;
	}

	if (exists($seen{$m})) {
		return;
	}
	$seen{$m} = 1;

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
				search($mo, $steps + 1);
				$n++;
			}
		}
	}
}
