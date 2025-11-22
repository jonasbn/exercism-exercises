package Hamming;

use v5.40;

use Exporter qw<import>;
our @EXPORT_OK = qw<hamming_distance>;

sub hamming_distance ( $strand1, $strand2 ) {
    my @left_strand  = split //, $strand1;
    my @right_strand = split //, $strand2;

    if ( @left_strand != @right_strand ) {
        die 'left and right strands must be of equal length';
    }

    my $distance = 0;
    for my $i ( 0 .. $#left_strand ) {
        $distance++ if $left_strand[$i] ne $right_strand[$i];
    }

    return $distance;
}

1;
