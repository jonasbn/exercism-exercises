package CryptoSquare;

use v5.40;

use Exporter qw<import>;
our @EXPORT_OK = qw<cipher>;
use POSIX qw<ceil floor>;
use Data::Dumper;

sub cipher ($text) {
    my $normalized = lc $text;
    $normalized =~ s/[^a-z0-9]//g;

    my $length_of_message = length $normalized;

    my $c = ceil(sqrt $length_of_message); # colunms
    my $r = floor(sqrt $length_of_message); # rows

    my @chars = split //, $normalized;

    if ($r * $c < $length_of_message) {
        $r++;
    }

    my @padding = (' ') x (($r * $c) - $length_of_message);
    push @chars, @padding;

    my @square;
    my $p = 0;
    my $encoded = '';
    for(my $i = 0; $i < $r; $i++) {
        for(my $j = 0; $j < $c; $j++) {
            my $character = @chars[$p++];
            $square[$j][$i] = $character;
        }
    }

    foreach my $first (@square) {
        foreach my $second (@{$first}) {
            $encoded .= $second;
        }
        $encoded .= " ";
    }
    chop $encoded;

    return $encoded;
}

1;
