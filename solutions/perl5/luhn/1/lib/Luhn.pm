package Luhn;

use v5.38;
use Data::Dumper;

use Exporter qw<import>;
our @EXPORT_OK = qw<is_luhn_valid>;

our $FALSE = 0;
our $TRUE  = 1;

sub is_luhn_valid ($number) {

    # remove white space
    $number =~ s/[\s]//g;
    
    my @digits = split //, $number;
    # reverse list of digits
    my @reverse_digits = reverse @digits;
    my $sum = 0;

    # fail if we only have a single digit or less
    if (scalar @reverse_digits <= 1) {
        return $FALSE;
    }

    # traverse the list of digits
    for (my $i = 0; $i < scalar @reverse_digits; $i++) {
        my $product = 0;

        my $digit = $reverse_digits[$i];

        # fail if digit is not a numeric
        if ($digit !~ m/\d/) {
            return $FALSE;
        }

        if ($i % 2 != 0) {
            $product = $digit * 2;

            if ($product > 9) {
                $product -= 9;
            }

        } else {
            $product = $digit;
        }

        $sum += $product;
    }

    if ($sum % 10 == 0) {
        return $TRUE;
    } else {
        return $FALSE;
    }
}
