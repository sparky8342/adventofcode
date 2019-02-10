package Map;
use strict;
use warnings;

use parent "Exporter";
our @EXPORT_OK = qw(furthest_path find_amount);

sub furthest_path {
	my ($route) = @_;
	($route) = $route =~ /^\^(.*)\$$/;
	1 while $route =~ s/\([^\(\)]+\|\)//g;
	1 while $route =~ s/\(([^\(\)]+)\)/(sort { length($b) <=> length($a) } split('\|',$1))[0]/eg;
	return length($route);
}

sub find_amount {
	my ($route,$doors) = @_;

	($route) = $route =~ /^\^(.*)\$$/;

	my %map = (0 => { 0 => 0 });
	my %pos = (x => 0, y => 0);
	my @stack;

	for (my $i = 0; $i < length($route); $i++) {
		my $c = substr($route,$i,1);
		if ($c eq '(') {
			push @stack, { x => $pos{x}, y => $pos{y} };
		}
		elsif ($c eq '|') {
			%pos = %{$stack[-1]};
		}
		elsif ($c eq ')') {
			%pos = %{pop @stack};
		}
		else {
			my $steps = $map{$pos{x}}{$pos{y}};
			if ($c eq 'N') { $pos{y}-- }
			if ($c eq 'S') { $pos{y}++ }
			if ($c eq 'W') { $pos{x}-- }
			if ($c eq 'E') { $pos{x}++ }
			$steps++;
			if (!exists($map{$pos{x}}{$pos{y}}) || $map{$pos{x}}{$pos{y}} > $steps) {
				$map{$pos{x}}{$pos{y}} = $steps;
			}
		}
	}

	my $count = 0;
	foreach my $x (keys %map) {
		foreach my $y (keys %{$map{$x}}) {
			if ($map{$x}{$y} >= $doors) {
				$count++;
			}
		}
	}
	return $count;
}
