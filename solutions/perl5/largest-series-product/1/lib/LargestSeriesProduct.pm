package LargestSeriesProduct;

use v5.40;

use Exporter qw<import>;
our @EXPORT_OK = qw<largest_product>;
use Data::Dumper;

sub largest_product ( $digits, $span ) {

    if ($digits =~ m/\D/) {
        die "digits input must only contain digits";
    }

    if ($span < 0) {
        die "span must not be negative";
    }

    if ($span > length $digits) {
        die "span must be smaller than string length";
    }

    my $largest_product = 0;

    for (my $i = 0; $i < length($digits); $i++) {

        if ($i + $span > length($digits)) {
            last;
        }

        my $number = substr($digits, $i, $span);
        my $product = product($number);

        if ($product > $largest_product) {
            $largest_product = $product;
        }
    }

    return $largest_product;
}

sub product ( $digits ) {

    my @numbers = split //, $digits;

    my $product = 1;
    foreach my $number (@numbers) {
        $product *= $number;
    }
    return $product;

}
1;
