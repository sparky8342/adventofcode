package Memory;
use strict;
use warnings;

use parent "Exporter";
our @EXPORT_OK = qw(sum_metadata);

sub sum_metadata {
	my ($data) = @_;
	my @n = split/ /,$data;

	my $tree = { node_number => 0, sum => 0 };

	#build_tree($sum,@n);

	build_tree($tree,@n);

	use Data::Dumper;
	print Dumper $tree;

	my $sum = find_sum($tree);
	my $value = find_value($tree);

	return ($sum,$value);
	#return $tree->{sum};
}

sub build_tree {
	my ($tree,@n) = @_;

	#print '->' . join(",",@n) . "\n";

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

	#$tree->{sum} += $_ foreach @metadata;

	#print join(",",@metadata) . "\n";

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

	use Data::Dumper;
	print Dumper $node;

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





__END__
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
