package Device2;
use strict;
use warnings;

use parent "Exporter";
our @EXPORT_OK = qw(run_program);

my @r = (0,0,0,0,0,0);
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

sub run_program {
	my @data = @_;

	my $bind = shift(@data);
	my ($bind_reg) = $bind =~ /(\d)/;
	my $ip = $r[$bind_reg];

	my @program;
	foreach my $line (@data) {
		push @program, [split/ /,$line];
	}

	@r = (0,0,0,0,0,0);

	while (1) {
		$r[$bind_reg] = $ip;
		#print "$ip\n";
		$opcodes{$program[$ip]->[0]}->($program[$ip]->[1],$program[$ip]->[2],$program[$ip]->[3]);
		$ip = $r[$bind_reg];
		$ip++;
		last if $ip < 0 || $ip >= @program;
	}

	return $r[0];
}
