package Allergies;

use v5.38;

use Exporter qw<import>;
our @EXPORT_OK = qw<allergic_to list_allergies>;

my %allergy_table = (
    1 => "eggs",
    2 => "peanuts",
    4 => "shellfish",
    8 => "strawberries",
    16 => "tomatoes",
    32 => "chocolate",
    64 => "pollen",
    128 => "cats",
);


sub allergic_to ( $item, $score ) {

    my @allergies = @{list_allergies($score)};

    if ( grep(/$item/, @allergies)) {
        return 1;
    } else {
        return 0;
    }

}

sub list_allergies ($score) {

    my $sum = 0;

    foreach my $key (keys %allergy_table) {
        $sum += $key;
    }

    my @allergies = ();

    if ($score > $sum) {
        $score -= $sum+1;
    }

    foreach my $key (sort {$b <=> $a} keys %allergy_table) {
        if ($score >= $key) {
            $score -= $key;
            push @allergies, $allergy_table{$key};
        }
    }

    return \@allergies;
}
