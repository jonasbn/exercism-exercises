package ArmstrongNumbers;

use v5.40;

use Exporter qw<import>;
our @EXPORT_OK = qw<is_armstrong_number>;
use Math::BigInt;

sub is_armstrong_number ($number) {
    my $sum = Math::BigInt->new();
    my $factor = length $number;

    if ($number == 0) {
        return 1;
    }

    my @numbers = split //, $number;

    foreach my $digit (@numbers) {
        $sum += Math::BigInt->new($digit) ** $factor;
    }


    return $sum->beq($number)?1:0;
}

1;