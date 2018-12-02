package Inventory;
use strict;
use warnings;

use parent "Exporter";
our @EXPORT_OK = qw(checksum find_similar);

sub checksum {
	my @boxes = @_;
	my ($two,$three);

	foreach my $box (@boxes) {
		my %chars;
		$chars{$_}++ foreach split//,$box;

		my %v;
		$v{$_} = 1 foreach values %chars;

		$two++ if exists($v{2});
		$three++ if exists($v{3});	
	}

	return $two * $three;
}

sub find_similar {
	my @boxes = @_;

	while(1) {
		my @box = split//,shift(@boxes);
		BOX:
		foreach my $bx (@boxes) {
			my @b = split//,$bx;
			my $diff = 0;
			my $diff_pos;
			for (my $i = 0; $i < @box; $i++) {
				if ($box[$i] ne $b[$i]) {
					$diff++;
					if ($diff > 1) {
						next BOX;
					}
					$diff_pos = $i;
				}
			}
			splice(@box,$diff_pos,1);
			return join('',@box);
		}
	}
}

