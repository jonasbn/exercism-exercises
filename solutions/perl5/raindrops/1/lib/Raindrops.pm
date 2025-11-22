package Raindrops;

use v5.38;

use Exporter qw<import>;
our @EXPORT_OK = qw<raindrop>;

sub raindrop ($number) {
    my $result = '';

    $result .= 'Pling' if $number % 3 == 0;
    $result .= 'Plang' if $number % 5 == 0;
    $result .= 'Plong' if $number % 7 == 0;
    
    if ($result) {
        return $result;
    } else {
        return $number;
    }
}
