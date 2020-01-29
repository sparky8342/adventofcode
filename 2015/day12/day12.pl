#!/usr/bin/perl
use strict;
use warnings;
use JSON::XS qw(decode_json);

sub search {
	my ($val, $ignore_red) = @_;

	my $total = 0;

	if (ref($val) eq 'HASH') {
		if ($ignore_red) {
			my @v = values %$val;
			if (grep { $_ eq 'red' } @v) {
				return 0;
			}
		}
		foreach my $key (keys %$val) {
			$total += search($key, $ignore_red);
			$total += search($val->{$key}, $ignore_red);
		}
	}
	elsif (ref($val) eq 'ARRAY') {
		foreach my $entry (@$val) {
			$total += search($entry, $ignore_red);
		}
	}	
	elsif ($val =~ /^(-)?\d+$/) {
		$total += $val;
	}

	return $total;
}

open my $fh, '<', 'input.txt';
chomp(my $json = <$fh>);
close $fh;

my $data = decode_json($json);

print search($data) . "\n";
print search($data,1) . "\n";
