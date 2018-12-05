package Polymer;
use strict;
use warnings;

use parent "Exporter";
our @EXPORT_OK = qw(reduction find_best_after_removal);

my @ab = ('a'..'z');
my $search = '('.(join('|', map {lc($_).uc($_).'|'.uc($_).lc($_)} @ab)).')';

sub reduction {
	my ($polymer) = @_;
	while (1) {
		last unless $polymer =~ s/$search//g;
	}
	return $polymer;
}

sub find_best_after_removal {
	my ($polymer) = @_;

	my $best = $polymer;

	for my $l (@ab) {
		my $stripped = $polymer;
		$stripped =~ s/$l//ig;
		$stripped = reduction($stripped);
		if (length($stripped) < length($best)) {
			$best = $stripped;
		}
	}

	return $best;
}
