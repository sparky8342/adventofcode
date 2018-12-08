package Memory;
use strict;
use warnings;

use parent "Exporter";
our @EXPORT_OK = qw(sum_metadata);

sub sum_metadata {
	my ($data) = @_;
	my @n = split/ /,$data;

	my $sum = { sum => 0 };
	build_tree($sum,@n);

	return $sum->{sum};
}

sub build_tree {
	my ($sum,@n) = @_;

	#print '->' . join(",",@n) . "\n";

	my $no_children = shift(@n);
	my $no_metadata = shift(@n);

	for (my $i = 0; $i < $no_children; $i++) {
		@n = build_tree($sum,@n);
	}

	my @metadata;
	for (1..$no_metadata) {
		push @metadata, shift(@n);
	}

	$sum->{sum} += $_ foreach @metadata;
	#print join(",",@metadata) . "\n";

	return @n;
}
