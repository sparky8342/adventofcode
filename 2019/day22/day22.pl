#!/usr/bin/perl
use strict;
use warnings;
use Data::Dumper;

use constant MAX_CARD => 10006;

my @cards;
for (0..MAX_CARD) {
	push @cards, $_;
}

open my $fh, '<', 'input.txt';
chomp(my @ops = <$fh>);
close $fh;

foreach my $op (@ops) {
	if ($op eq 'deal into new stack') {
		@cards = reverse @cards;
	}
	elsif ($op =~ /^cut (.*)$/) {
		my $cut = $1;
		if ($cut > 0) {
			@cards = (@cards[$cut..@cards-1], @cards[0..$cut-1]);
		}
		else {
			$cut *= -1;
			@cards = (@cards[@cards-$cut..@cards-1], @cards[0..@cards-$cut-1]);
		}
	}
	elsif ($op =~ /^deal with increment (\d+)$/) {
		my $inc = $1;
		my $pos = 0;
		my @new_cards;
		foreach my $card (@cards) {
			$new_cards[$pos] = $card;
			$pos += $inc;
			$pos = $pos % scalar(@cards);
		}
		@cards = @new_cards;
	}
}	
		
for (my $i = 0; $i <= MAX_CARD; $i++) {
	if ($cards[$i] == 2019) {
		print "$i\n";
		last;
	}
}
