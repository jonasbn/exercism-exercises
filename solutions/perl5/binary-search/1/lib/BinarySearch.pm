package BinarySearch;

use v5.38;
use Data::Dumper;

use Exporter qw<import>;
our @EXPORT_OK = qw<binary_search>;

sub binary_search ( $array, $value ) {

    my @sorted_array = sort { $a <=> $b } @$array;

    return _binary_search(\@sorted_array, $value, 0, scalar(@sorted_array) - 1);
}

sub _binary_search ( $array, $value, $low, $high ) {
    my $middle = int(($low + $high) / 2);
    my $middle_value = $array->[$middle];

    if ($low > $high) {
        die 'value not in array';
    }

    if ($middle_value == $value) {
        return $middle;
    } elsif ($middle_value < $value) {
        return _binary_search($array, $value, $middle + 1, $high);
    } else {
        return _binary_search($array, $value, $low, $middle - 1);
    }
}