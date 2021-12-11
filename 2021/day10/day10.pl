#!/usr/bin/perl
use strict;
use warnings;

my %error = (')' => 3, ']' => 57, '}' => 1197, '>' => 25137);
my %complete = ('(' => 1, '[' => 2, '{' => 3, '<' => 4);

my $fh;
open $fh, '<', 'input.txt';
chomp(my @lines = <$fh>);
close $fh;

my $error_score = 0;
my @complete_scores;
for my $line (@lines) {
	my $len = 0;
	while ($len != length($line)) {
		$len = length($line);
		$line =~ s/\(\)|\[\]|\{\}|\<\>//g;
	}
	my $li = $line;
	$line =~ s/^[\(\[\{\<]+//;
	if (length($line) > 0) {
		$error_score += $error{substr($line, 0, 1)};
	}
	else {
		$li = reverse($li);
		my $complete_score = 0;
		$li =~ s/(.)/$complete_score = $complete_score * 5 + $complete{$1}/ge;
		push @complete_scores, $complete_score;
	}	
}

print "$error_score\n";
@complete_scores = sort { $a <=> $b } @complete_scores;
print $complete_scores[@complete_scores / 2] . "\n";
