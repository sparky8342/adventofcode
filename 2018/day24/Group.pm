package Group;
use strict;
use warnings;

sub new {
	my ($class, $self) = @_;
	if (!defined($self)) {
		$self = {};
	}
	bless $self, $class;
	return $self;   
}

sub effective_power {
	my ($self) = @_;
	return $self->{units} * $self->{attack};
}

sub damage {
	my ($self, $target) = @_;

	my $damage = $self->effective_power();

	if (exists($target->{immune}->{$self->{attack_type}})) {
		$damage = 0;
	} elsif (exists($target->{weak}->{$self->{attack_type}})) {
		$damage *= 2;
	}

	return $damage;
}

1;
