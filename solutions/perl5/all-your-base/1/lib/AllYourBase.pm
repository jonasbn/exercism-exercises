package AllYourBase;

use v5.38;
use Data::Dumper;

use Exporter qw<import>;
our @EXPORT_OK = qw<rebase>;

my @errors = (
    'input base must be >= 2',
    'output base must be >= 2',
    'all digits must satisfy 0 <= d < input base',
);

sub rebase ( $digits, $input_base, $output_base ) {

    if ( $input_base < 2 ) {
        die $errors[0];
    } elsif ( $output_base < 2 ) {
        die $errors[1];
    } elsif ( grep { $_ < 0 || $_ >= $input_base } @{$digits} ) {
        die $errors[2];
    }

    if ( $input_base == $output_base ) {
        return $digits;
    }

    if ($input_base != 10) {

        my @rebased_digits_in_base_10 = ();
        my $rebased_digit_in_base_10 = 0;
        my $exponent = 0;

        foreach my $digit (reverse @{$digits}) {
            $rebased_digit_in_base_10 += $digit * ($input_base ** $exponent);
            $exponent++;
        }
        if ($rebased_digit_in_base_10 > 9) {
            push @rebased_digits_in_base_10, split //, $rebased_digit_in_base_10;
        } else {
            push @rebased_digits_in_base_10, $rebased_digit_in_base_10;
        }
        $digits = \@rebased_digits_in_base_10;
    }

    # Converting to other bases from base 10:
    # REF: https://math.libretexts.org/Courses/College_of_the_Canyons/Math_130%3A_Math_for_Elementary_School_Teachers_(Lagusker)/02%3A_Empathy_and_Primary_Mathematics/2.06%3A_Converting_Between_(our)_Base_10_and_Any_Other_Base_(and_vice_versa)
    if ($output_base != 10) {
    
        my $base_10_number = join '', @{$digits};
        my $division_result;
        my @rebased_digits;

        # I have never used do/while before, but it seems like the right tool for this job
        # REF: https://www.perltutorial.org/perl-do-while/
        do {
            my $remainder = $base_10_number % $output_base;
            # Rounding without Math::Round:
            # REF: https://www.oreilly.com/library/view/perl-cookbook/1565922433/ch02s04.html
            # I chose int() over sprintf()
            $division_result = int($base_10_number / $output_base);
            push @rebased_digits, $remainder;
            $base_10_number = $division_result;
        } while ($division_result > 0);

        @rebased_digits = reverse @rebased_digits;
        $digits = \@rebased_digits;
    }
    
    return $digits;
}
