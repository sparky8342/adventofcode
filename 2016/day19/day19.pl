#!/usr/bin/perl
use strict;
use warnings;

use constant NUM_ELVES => 3005290;

sub init_elves() {
    my $head = { id => 1 };

    my $elf = $head;

    for ( my $i = 2 ; $i <= NUM_ELVES ; $i++ ) {
        my $next = { id => $i };
        $elf->{next} = $next;
        $elf = $next;
    }

    $elf->{next} = $head;
    return $head;
}

sub part1() {
    my $elf = init_elves();

    while ( $elf != $elf->{next} ) {
        $elf->{next} = $elf->{next}->{next};
        $elf = $elf->{next};
    }

    return $elf->{id};
}

sub part2() {
    my $elf = init_elves();

    my $opposite = $elf;

    for ( my $i = 1 ; $i < NUM_ELVES / 2 ; $i++ ) {
        $opposite = $opposite->{next};
    }

    while ( $elf != $elf->{next} ) {
        $opposite->{next} = $opposite->{next}->{next};
        $elf = $elf->{next};

        $opposite->{next} = $opposite->{next}->{next};
        $opposite         = $opposite->{next};
        $elf              = $elf->{next};
    }

    return $elf->{id};
}

print part1() . "\n" . part2() . "\n";
