#!/usr/bin/perl
use strict;
use warnings;

open my $fh, '<', 'input.txt';
chomp(my @program = <$fh>);
close $fh;

# create perl scripts from the code
for my $part ('part1', 'part2') {
	my $a = $part eq 'part1' ? 0 : 1;
	my $script = "#!/usr/bin/perl\nuse strict;\nuse warnings;\n\n";
	$script .= "# auto-generated from the input source\n\n";
	$script .= "my \$registers = { a => $a, b => 0 };\n";

	my $line_no = 0;
	my %labels;
	my %labels_called;

	foreach my $line (@program) {
		my $label = 'LINE' . $line_no;
		$labels{$label} = 1;
		$script .= "$label: ";
		if ($line =~ /^hlf (.)$/) {
			$script .= "\$registers->{$1} /= 2;\n";
		}
		elsif ($line =~ /^tpl (.)$/) {
			$script .= "\$registers->{$1} *= 3;\n";
		}
		elsif ($line =~ /^inc (.)$/) {
			$script .= "\$registers->{$1}++;\n";
		}
		elsif ($line =~ /^jmp (.+)$/) {
			my $label = 'LINE' . ($line_no + $1);
			$labels_called{$label} = 1;
			$script .= "goto $label;\n";
		} 
		elsif ($line =~/^jie (.), (.+)$/) {
			my $label = 'LINE' . ($line_no + $2);
			$labels_called{$label} = 1;
			$script .= "if (\$registers->{$1} % 2 == 0) { goto $label }\n";
		}
		elsif ($line =~/^jio (.), (.+)$/) {
			my $label = 'LINE' . ($line_no + $2);
			$labels_called{$label} = 1;
			$script .= "if (\$registers->{$1} == 1) { goto $label }\n";
		}
		$line_no++;
	}
	foreach my $label (sort keys %labels_called) {
		if (!exists($labels{$label})) {
			$script .= "$label:\n";
		}
	}
	foreach my $label (keys %labels) {
		if (!exists($labels_called{$label})) {
			$script =~ s/$label: //g;
		}
	}

	$script .= "print \$registers->{b} . \"\\n\";\n";
	open my $fh, '>', "$part.pl";
	print $fh $script;
	close $fh;
}
