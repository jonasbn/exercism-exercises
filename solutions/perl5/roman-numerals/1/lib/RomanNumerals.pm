package RomanNumerals;

use v5.38;

use Exporter qw<import>;
our @EXPORT_OK = qw<to_roman>;

sub to_roman ($number) {

    my @positional_sets = (
        {
            primary   => 'I',
            secondary => 'V',
            tertiary  => 'X',
        },
        {
            primary   => 'X',
            secondary => 'L',
            tertiary  => 'C',
        },
        {
            primary   => 'C',
            secondary => 'D',
            tertiary  => 'M',
        },
        {
            primary => 'M',
        },
    );

    my @digits = reverse split //, $number;

    my $complete_roman_numeral = '';

    for (my $i = 0; $i < scalar @digits; $i++) {
        my ($remainder, $multiplier) = _calculate_remainder_and_multiplier($digits[$i]);

        my $roman_numeral = '';
        my $set = $positional_sets[$i];

        # The roman numeral does not require duplication and we can just append it
        if ($multiplier > 0 and $remainder < 4) {
            $roman_numeral = _postfix($roman_numeral, $set->{secondary}, 1);
        }
        
        if ($remainder == 4) {
            # The roman numeral requires subtraction, meaning we append the highest available
            # numeral in the set and prefix the required amount of the lowest available numeral
            # from the set the required amount of times
            if ($multiplier > 0) {
                $roman_numeral = _postfix($roman_numeral, $set->{tertiary}, 1);
                $roman_numeral = _prefix($roman_numeral, $set->{primary}, 5 - $remainder);

            # The roman numeral requires subtraction, meaning we append the next highest available
            # numeral in the set and prefix the required amount of the lowest available numeral
            # from the set the required amount of times
            } else {
                $roman_numeral = _postfix($roman_numeral, $set->{secondary}, 1);
                $roman_numeral = _prefix($roman_numeral, $set->{primary}, 5 - $remainder);
            }
        } elsif ($remainder > 0) {

            # The roman numeral requires addition meaning we append the lowest available numeral
            # from the set the required amount of times
            $roman_numeral = _postfix($roman_numeral, $set->{primary}, $remainder);
        }

        $complete_roman_numeral =  $roman_numeral . $complete_roman_numeral;
    };

    return $complete_roman_numeral;
}

sub _prefix {
    my ($roman_numeral, $character, $amount) = @_;

    my $prefix = $character x $amount;

    $roman_numeral = $prefix .= $roman_numeral;

    return $roman_numeral;
}

sub _postfix {
    my ($roman_numeral, $character, $amount) = @_;

    $roman_numeral .= $character x $amount;

    return $roman_numeral;
}

sub _calculate_remainder_and_multiplier {
    my $number = shift;

    my $remainder = $number % 5;
    my $multiplier = int($number / 5);

    return ($remainder, $multiplier);
}