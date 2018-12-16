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
	eqrr => sub { $r[$_[2]] = $r[$_[0]] == $r[$_[1]] ? 1 : 0 }, 
);

sub process {
	my ($data,$program) = @_;
	my @samples;

	for (my $i = 0; $i < @$data; $i += 4) {
		my @b = $data->[$i] =~ /(\d), (\d), (\d), (\d)/;
		my ($op,@ins) = $data->[$i+1] =~ /(\d+) (\d) (\d) (\d)/;
		my @a = $data->[$i+2] =~ /(\d), (\d), (\d), (\d)/;
		push @samples, { before => \@b, op => $op, ins => \@ins, after => \@a };
	}

	my $count = 0;
	foreach my $sample (@samples) {
		my @ops;
		foreach my $op (sort keys %opcodes) {
			@r = @{$sample->{before}};
			$opcodes{$op}->(@{$sample->{ins}});
			if (join('',@{$sample->{after}}) == join('',@r)) {
				push @ops, $op;
			}
		}
		$count++ if @ops >= 3;
	}

	my %ops;
	foreach my $op (sort keys %opcodes) {
		for (0..15) { 
			$ops{$op}{$_} = 1;
		}
	}

	foreach my $sample (@samples) {
		foreach my $op (sort keys %opcodes) {
			@r = @{$sample->{before}};
			$opcodes{$op}->(@{$sample->{ins}});
			if (join('',@{$sample->{after}}) != join('',@r)) {
				delete($ops{$op}{$sample->{op}});
			}
		}
	}

	while (1) {
		foreach my $op (keys %ops) {
			my @k = keys %{$ops{$op}};
			if (@k == 1) {
				foreach my $op2 (keys %ops) {
					next if $op eq $op2;
					delete($ops{$op2}{$k[0]});
				}
			}
		}
		my %r;	
		foreach my $op (keys %ops) {
			foreach my $k (keys %{$ops{$op}}) {
				$r{$k}{$op} = 1;
			}
		}
		foreach my $k (keys %r) {
			my @ke = keys %{$r{$k}};
			if (@ke == 1) {
				$ops{$ke[0]} = {$k => 1};
			}
		}
		my $done = 1;
		foreach my $op (keys %ops) {
			if (keys %{$ops{$op}} > 1) {
				$done = 0;
				last;
			}
		}

		last if $done;
	}
	my %op;
	foreach my $key (keys %ops) {
		$op{(keys %{$ops{$key}})[0]} = $key;
	}

	@r = (0,0,0,0);
	foreach my $line (@$program) {
		my ($opc,@ar) = split(/ /,$line);
		$opcodes{$op{$opc}}->(@ar);
	}

	return ($count,$r[0]);
}
