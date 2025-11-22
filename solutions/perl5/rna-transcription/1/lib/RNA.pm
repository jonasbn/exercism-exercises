package RNA;

use v5.40;

use Exporter qw<import>;
our @EXPORT_OK = qw<to_rna>;



sub to_rna ($dna) {
    my @nucleotides = split //, $dna;

    my %complement = (
        G => 'C',
        C => 'G',
        T => 'A',
        A => 'U',
    );

    my $rna = '';
    foreach my $nucleotide (@nucleotides) {
        $rna .= $complement{$nucleotide};
    }
    return $rna;

}

1;
