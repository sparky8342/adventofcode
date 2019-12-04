#!/usr/bin/perl
use strict;
use warnings;

sub valid_part1 {
	my ($pass) = @_;
	return 0 if $pass !~ /(.)\1/;
	return 0 if $pass =~ /(98|97|96|95|94|93|92|91|90|87|86|85|84|83|82|81|80|76|75|74|73|72|71|70|65|64|63|62|61|60|54|53|52|51|50|43|42|41|40|32|31|30|21|20)/;
	return 1;
}

sub valid_part2 {
	my ($pass) = @_;
	$pass =~ s/(.)\1{2,}/$1/g;
	return valid_part1($pass);
}

my $part1 = 0;
my $part2 = 0;
for my $pass (138241..674034) {
	$part1++ if valid_part1($pass);
	$part2++ if valid_part2($pass);
}
print "$part1\n$part2\n";
