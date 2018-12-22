package Map;
use strict;
use warnings;

use parent "Exporter";
our @EXPORT_OK = qw(furthest_path);

sub furthest_path {
	my ($route) = @_;
	($route) = $route =~ /^\^(.*)\$$/;
	1 while $route =~ s/\([^\(\)]+\|\)//g;
	1 while $route =~ s/\(([^\(\)]+)\)/(sort { length($b) <=> length($a) } split('\|',$1))[0]/eg;
	return length($route);
}
