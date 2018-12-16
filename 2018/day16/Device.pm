package Device;
use strict;
use warnings;

use parent "Exporter";
our @EXPORT_OK = qw(process);

my @r;
my %opcodes = (
	addr => sub { $r[$_[2]] = $r[$_[0]] + $r[$_[1]] },
	addi => sub { $r[$_[2]] = $r[$_[0]] + $_[1] },
	mulr => sub { $r[$_[2]] = $r[$_[0]] * $r[$_[1]] },
	muli => sub { $r[$_[2]] = $r[$_[0]] * $_[1] },
	banr => sub { $r[$_[2]] = $r[$_[0]] & $r[$_[1]] },
	bani => sub { $r[$_[2]] = $r[$_[0]] & $_[1] },
	borr => sub { $r[$_[2]] = $r[$_[0]] | $r[$_[1]] },
	bori => sub { $r[$_[2]] = $r[$_[0]] | $_[1] },
	setr => sub { $r[$_[2]] = $r[$_[0]] },
	seti => sub { $r[$_[2]] = $_[0] },
	grir => sub { $r[$_[2]] = $_[0] > $r[$_[1]] ? 1 : 0 }, 
	gtri => sub { $r[$_[2]] = $r[$_[0]] > $_[1] ? 1 : 0 }, 
	gtrr => sub { $r[$_[2]] = $r[$_[0]] > $r[$_[1]] ? 1 : 0 }, 
	eqir => sub { $r[$_[2]] = $_[0] == $r[$_[1]] ? 1 : 0 }, 
	eqri => sub { $r[$_[2]] = $r[$_[0]] == $_[1] ? 1 : 0 }, 
	gtrr => sub { $r[$_[2]] = $r[$_[0]] == $r[$_[1]] ? 1 : 0 }, 
);

sub process {
	my (@data) = @_;
	my @samples;

	for (my $i = 0; $i < @data; $i += 4) {
		my @b = $data[$i] =~ /(\d), (\d), (\d), (\d)/;
		my ($op,@ins) = $data[$i+1] =~ /(\d) (\d) (\d) (\d)/;
		my @a = $data[$i+2] =~ /(\d), (\d), (\d), (\d)/;
		push @samples, { before => \@b, op => $op, ins => \@ins, after => \@a };
	}

	my $count = 0;
	foreach my $sample (@samples) {
		my $opcount = 0;
		foreach my $op (keys %opcodes) {
			@r = @{$sample->{before}};
			$opcodes{$op}->(@{$sample->{ins}});
			if (join('',@{$sample->{after}}) == join('',@r)) {
				$opcount++;
			}
		}
		$count++ if $opcount >= 3;
	}

	return $count;
}
