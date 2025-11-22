package Sieve;

use v5.38;

use Exporter qw<import>;
our @EXPORT_OK = qw<find_primes>;

sub find_primes ($limit) {
    my %numbers;
    for my $number (0..$limit) {
        $numbers{$number}++;
    }

    for (my $i = 2; $i < keys %numbers; $i++) {
        for (my $j = $i; $j < keys %numbers; $j += $i) {
            if ($numbers{$j}) {
                if ($i != $j) {
                    $numbers{$j} = 0;
                }
            }
        }
    }

    my @marked_numbers = ();

    for (my $k = 2; $k < keys %numbers; $k++) {
        if ($numbers{$k}) {
            push @marked_numbers, $k;
        }
    }

    return \@marked_numbers;
}
