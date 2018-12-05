package Polymer;
use strict;
use warnings;

use parent "Exporter";
our @EXPORT_OK = qw(reduction);

use constant AB => qw/abcdefghijklmnopqrstuvwxyz/;

my $search;
foreach my $l (split//,&AB) {
	$search .= '|' if $search;
	$search .= lc($l) . uc($l) . '|' . uc($l) . lc($l);
}
$search = "($search)";

sub reduction {
	my ($polymer) = @_;
	while (1) {
		last unless $polymer =~ s/$search//g;
	}
	return $polymer;
}
