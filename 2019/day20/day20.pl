#!/usr/bin/perl
use strict;
use warnings;

my $grid = [];
open my $fh, '<', 'input.txt';
while (my $row = <$fh>) {
	chomp($row);
	push @$grid, [split//, $row];
}
close($fh);

my %portals;

my %portal_info;

my $height = @$grid;
my $width = @{$grid->[0]};

my $start;
my $end;

for (my $y = 0; $y < $height; $y++) {
	for (my $x = 0; $x < $width; $x++) {

		if ($grid->[$y][$x] =~ /^([A-Z])$/) {
			my $portal = $1; 
			my $use_no;
			my ($x2, $y2) = ($x, $y);
			if ($grid->[$y+1][$x] =~ /^([A-Z])$/) {
				$portal .= $1;
				$y2++;
			}
			elsif ($grid->[$y][$x+1] =~ /^([A-Z])$/) {
				$portal .= $1;
				$x2++;
			}

			if ($portal eq 'AA' || $portal eq 'ZZ') {
				my ($pathx, $pathy) = neighbour_path($x, $y);
				if ($pathx == -1) {	
					($pathx, $pathy) = neighbour_path($x2, $y2);
				}
				if ($portal eq 'AA') {
					$start = { x => $pathx, y => $pathy, dist => 0 };
				}
				elsif ($portal eq 'ZZ') {
					$end = { x => $pathx, y => $pathy };
				}
				$grid->[$y][$x] = '#';
				$grid->[$y2][$x2] = '#';
			}
			else {
				my ($pathx, $pathy) = neighbour_path($x, $y);
				if ($pathx != -1) {
					$portal_info{$portal}{"$x $y"} = [$pathx, $pathy];
					$grid->[$y][$x] = $portal;
					$grid->[$y2][$x2] = ' ';
				}
				else {
					($pathx, $pathy) = neighbour_path($x2, $y2);
					$portal_info{$portal}{"$x2 $y2"} = [$pathx, $pathy];
					$grid->[$y2][$x2] = $portal;
					$grid->[$y][$x] = ' ';
				}
			}
		}
	}
}

# swap destinations
foreach my $no (keys %portal_info) {
	my @coord = keys %{$portal_info{$no}};
	my @paths = ($portal_info{$no}{$coord[0]}, $portal_info{$no}{$coord[1]});
	($portal_info{$no}{$coord[0]}, $portal_info{$no}{$coord[1]}) = ($paths[1], $paths[0]);
}
		
my $steps = bfs();
print "$steps\n";

sub neighbour_path {
	my ($x, $y) = @_;
	if ($y > 0 && $grid->[$y-1][$x] eq '.') {
		return ($x, $y-1);
	}
	elsif ($y < $height - 1 && $grid->[$y+1][$x] eq '.') {
		return ($x, $y+1);
	}
	elsif ($x > 0 && $grid->[$y][$x-1] eq '.') {
		return ($x-1, $y);
	}
	elsif ($x < $width - 1 && $grid->[$y][$x+1] eq '.') {
		return ($x+1, $y);
	}

	return (-1, -1);
}

sub print_grid {
	for (my $y = 0; $y < $height; $y++) {
		for (my $x = 0; $x < $width; $x++) {
			printf("%-2s",$grid->[$y][$x]);
		}
		print "\n";
	}
}	

sub bfs {
	my @queue = ($start);
	my %visited;

	while (@queue) {
		my $space = shift(@queue);

		if (exists($visited{$space->{x}}{$space->{y}})) {
			next;
		}
		$visited{$space->{x}}{$space->{y}} = 1;

		if ($grid->[$space->{y}][$space->{x}] eq '#') {
			next;
		}
		elsif ($grid->[$space->{y}][$space->{x}] =~ /^[A-Z]{2}$/) {
			my $tele = $portal_info{$grid->[$space->{y}][$space->{x}]}{$space->{x} . ' ' . $space->{y}};
			unshift @queue, { x => $tele->[0], y => $tele->[1], dist => $space->{dist} };
			next;
		}
		elsif ($space->{x} == $end->{x} && $space->{y} == $end->{y}) {
			return $space->{dist};
		}

		push @queue, (
			{ x => $space->{x} + 1, y => $space->{y},     dist => $space->{dist} + 1, dir => $space->{dir} || 4 },
			{ x => $space->{x} - 1, y => $space->{y},     dist => $space->{dist} + 1, dir => $space->{dir} || 3 },
			{ x => $space->{x}    , y => $space->{y} + 1, dist => $space->{dist} + 1, dir => $space->{dir} || 2 },
			{ x => $space->{x}    , y => $space->{y} - 1, dist => $space->{dist} + 1, dir => $space->{dir} || 1 }
		);
	}
}
