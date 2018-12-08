package Memory;
use strict;
use warnings;

use parent "Exporter";
our @EXPORT_OK = qw(process_data);

sub process_data {
	my ($data) = @_;
	my @n = split/ /,$data;

	my $tree = {};
	build_tree($tree,@n);

	my $sum = find_sum($tree);
	my $value = find_value($tree);

	return ($sum,$value);
}

sub build_tree {
	my ($tree,@n) = @_;

	my $no_children = shift(@n);
	my $no_metadata = shift(@n);

	$tree->{children} = [];
	for (my $i = 0; $i < $no_children; $i++) {
		my $child = {};
		push @{$tree->{children}}, $child;
		@n = build_tree($child,@n);
	}

	my @metadata;
	for (1..$no_metadata) {
		push @metadata, shift(@n);
	}

	$tree->{metadata} = \@metadata;
	return @n;
}

sub find_sum {
	my ($node) = @_;
	my $metadata = $node->{metadata};
	my $sum = 0;
	$sum += $_ foreach @{$node->{metadata}};
	$sum += find_sum($_) foreach @{$node->{children}};
	return $sum;
}

sub find_value {
	my ($node) = @_;

	my $metadata = $node->{metadata};
	my $children = $node->{children};
	my $sum = 0;
	if (@$children == 0) {
		$sum += $_ foreach @{$node->{metadata}};
	}
	else {
		foreach my $entry (@{$node->{metadata}}) {
			if (defined($children->[$entry - 1])) {
				$sum += find_value($children->[$entry - 1]);
			}
		}
	}
	return $sum;
}
