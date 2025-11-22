package NucleotideCount;

use v5.40;

use Exporter qw<import>;
our @EXPORT_OK = qw<count_nucleotides>;

sub count_nucleotides ($strand) {
    my @nucleotides = split //, $strand;

    my %nycleotide_count = (A => 0, C => 0, G => 0, T => 0);

    foreach my $nucleotide (@nucleotides) {
        die "Invalid nucleotide in strand" unless exists $nycleotide_count{$nucleotide};
        $nycleotide_count{$nucleotide}++;
    }

    return \%nycleotide_count;
}

1;
