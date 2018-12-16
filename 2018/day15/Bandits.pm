package Bandits;
use strict;
use warnings;

use parent "Exporter";
our @EXPORT_OK = qw(combat);

sub combat {
	my @data = @_;

	my @grid;
	foreach my $row (@data) {
		push @grid, [split//,$row];
	}

	for (my $row = 0; $row < @grid; $row++) {
		for (my $col = 0; $col < @{$grid[$row]}; $col++) {
			if ($grid[$row][$col] =~ /^[EG]$/) {
				my $unit = {
					type => $grid[$row][$col],
					hitpoints => 200,
					last_row => $row,
					last_col => $col
				};
				$grid[$row][$col] = $unit;
			}
		}
	}

	my $turns = 0;

	LOOP:
	while (1) {
		pp(\@grid);

		for (my $row = 0; $row < @grid; $row++) {
			for (my $col = 0; $col < @{$grid[$row]}; $col++) {
				if (ref($grid[$row][$col])) {
					$grid[$row][$col]->{moved} = 0;
				}
			}
		}
		
		for (my $row = 0; $row < @grid; $row++) {
			for (my $col = 0; $col < @{$grid[$row]}; $col++) {
				if (ref($grid[$row][$col]) eq 'HASH') {	
					if (end_check(\@grid)) {
						last LOOP;
					}
					my $unit = $grid[$row][$col];
					next if $unit->{moved} == 1;
					my ($new_row,$new_col) = bfs(\@grid,$unit,$row,$col);
					if ($new_row == -1) {
						next;
					}
					$grid[$row][$col] = '.';
					$unit->{last_row} = $row;
					$unit->{last_col} = $col;
					$unit->{moved} = 1;
					$grid[$new_row][$new_col] = $unit;
					attack(\@grid,$new_row,$new_col);
				}
			}
		}
		$turns++;
	}

	my $sum = 0;
	for (my $row = 0; $row < @grid; $row++) {
		for (my $col = 0; $col < @{$grid[$row]}; $col++) {
			if (ref($grid[$row][$col])) {
				$sum += $grid[$row][$col]->{hitpoints};
			}
		}
	}

	return $sum * $turns;
}


sub bfs {
	my ($grid,$unit,$row,$col) = @_;

	my @queue = ({row => $row, col => $col, distance => 0, parent => 0});

	my $enemy_type = $unit->{type} eq 'E' ? 'G' : 'E';

	my @enemies;

	my %visited;
	while (@queue) {
		my $square = shift(@queue);
		next if (exists($visited{$square->{row}}{$square->{col}}));
		$visited{$square->{row}}{$square->{col}} = 1;

		if ($square->{row} != $row || $square->{col} != $col) {
			my $gridsquare = $grid->[$square->{row}][$square->{col}];
			if (ref($gridsquare)) {
				if ($gridsquare->{type} eq $enemy_type) {
					push @enemies, {
						row => $square->{row},
						col => $square->{col},
						parent => $square->{parent},
						distance => $square->{distance},
						last_row => $gridsquare->{last_row},
						last_col => $gridsquare->{last_col}
					};
				}
				next;
			}
		}

		for my $dir (
			{row => -1, col =>  0},
			{row =>  0, col => -1},
			{row =>  0, col =>  1},
			{row =>  1, col =>  0}
		) {
			my $drow = $square->{row} + $dir->{row};
			my $dcol = $square->{col} + $dir->{col};
			if (defined($grid->[$drow]) && defined($grid->[$drow][$dcol]) && $grid->[$drow][$dcol] ne '#') {
				push @queue, {
					row => $drow,
					col => $dcol,
					distance => $square->{distance} + 1,
					parent => $square
				};
			}
		}
	}

	if (@enemies == 0) {
		return (-1,-1);
	}

	@enemies = sort {
		$a->{parent}->{distance} <=> $b->{parent}->{distance}
		#$a->{distance} <=> $b->{distance}
		#$a->{last_row} <=> $b->{last_row} ||
		#$a->{last_col} <=> $b->{last_col}
	} @enemies;
	my $target = $enemies[0];

	while ($target->{parent}->{parent} != 0) {
		$target = $target->{parent};
	}

	if (ref($grid->[$target->{row}][$target->{col}])) {
		return($row,$col);
	}
	else {
		return ($target->{row},$target->{col});
	}
}

sub attack {
	my ($grid,$row,$col) = @_;

	my $unit = $grid->[$row][$col];
	my $enemy = $unit->{type} eq 'E' ? 'G' : 'E';

	my $hitpoints = 300;
	my $attack_row;
	my $attack_col;

	for my $dir (
		{row => -1, col =>  0},
		{row =>  0, col => -1},
		{row =>  0, col =>  1},
		{row =>  1, col =>  0}
	) {
		my $r = $row + $dir->{row};
		my $c = $col + $dir->{col};
		my $square = $grid->[$r][$c];
		if (ref($square) && $square->{type} eq $enemy) {
			if ($square->{hitpoints} < $hitpoints) {	
				$hitpoints = $square->{hitpoints};
				$attack_row = $r;
				$attack_col = $c;
			}
		}
	}

	if ($attack_row) {
		$grid->[$attack_row][$attack_col]->{hitpoints} -= 3;
	
		if ($grid->[$attack_row][$attack_col]->{hitpoints} <= 0) {
			$grid->[$attack_row][$attack_col] = '.';
			return 1;
		}
	}

	return 0;
}
					
sub end_check {
	my ($grid) = @_;
	my %u = ( E => 0, G => 0 );
	
	for (my $row = 0; $row < @$grid; $row++) {
		for (my $col = 0; $col < @{$grid->[$row]}; $col++) {
			if (ref($grid->[$row][$col])) {
				$u{$grid->[$row][$col]->{type}}++;
			}
		}
	}

	if ($u{E} == 0 || $u{G} == 0) {
		return 1;
	}
	else {
		return 0;
	}
}
		
sub pp {
	my ($grid) = @_;

	my @units;

	for (my $row = 0; $row < @$grid; $row++) {
		my $unitstr = '   ';
		for (my $col = 0; $col < @{$grid->[$row]}; $col++) {
			my $sq = $grid->[$row][$col];
	
			if (ref($sq)) {
				print $sq->{type};
				$unitstr .= $sq->{type} . '(' . $sq->{hitpoints} . '), ';
			}
			else {
				print $sq;
			}
		}
		print "$unitstr\n";
	}
	foreach my $unit (@units) {
		print join(',',@$unit) . "\n";
	}
	print "\n";
}
