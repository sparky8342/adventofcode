#!/usr/bin/perl
use strict;
use warnings;

open my $fh, '<', 'input.txt';
my @data = <$fh>;
close $fh;

my $code = 0;
my $mem = 0;
my $encoded = 0;
foreach my $line (@data) {
	chomp($line);
	$code += length($line);
	my $encode = $line;

	$line =~ s/^"(.*)"$/$1/;
	$line =~ s/\\\\/A/g;
	$line =~ s/\\x[0-9a-f]{2}/A/g;
	$line =~ s/\\\"/A/g;
	$mem += length($line);

	$encode =~ s/\\/AA/g;
	$encode =~ s/\"/AA/g;
	$encoded += length($encode) + 2;
}
# part 1
print $code - $mem . "\n";

# part 2
print $encoded - $code . "\n"; 
