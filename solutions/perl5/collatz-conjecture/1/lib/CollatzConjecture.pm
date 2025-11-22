package CollatzConjecture;

use v5.40;

use Exporter qw<import>;
our @EXPORT_OK = qw<steps>;

sub steps ($number) {
    if ($number <= 0) {
        die "Only positive integers are allowed\n";
    }

    my $steps = 0;
    while ($number != 1) {
        $number = $number % 2 == 0 ? $number / 2 : $number * 3 + 1;
        $steps++;
    }

    return $steps;
}

1;
