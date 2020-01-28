#!/usr/bin/perl
use strict;
use warnings;

use constant START => 97;
use constant EN => 122;

my $ab = 'abcdefghijklmnopqrstuvwxyz';
my @sets;
for (my $i = 0; $i < 24; $i++) {
	push @sets, substr($ab, $i, 3);
}
my $regex = '(' . join('|', @sets) . ')';

sub next_password {
	my ($password) = @_;
	my @c = map { ord($_) } split//, $password;
	my $e = @c - 1;
	my $str;
	while (1) {
		my $inc = $e;
		$c[$inc]++;
		while ($c[$inc] > EN) {
			$c[$inc] = START;
			$inc--;
			$c[$inc]++;
		}

		$str = join '', map { chr($_) } @c;
		next if $str =~ /[iol]/;
		next if $str !~ /(.)\1.*?(.)\2/;
		next if $str !~ /$regex/;
		return $str;
	}
}

open my $fh, '<', 'input.txt';
chomp(my $pass = <$fh>);
close $fh;

for (1..2) {
	$pass = next_password($pass);
	print "$pass\n";
}
