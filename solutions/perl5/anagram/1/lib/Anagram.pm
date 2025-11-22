package Anagram;

use v5.38;

use Exporter qw<import>;
our @EXPORT_OK = qw<match_anagrams>;

sub match_anagrams ( $subject, $candidates ) {
    my @anagrams;
    foreach my $candidate ( @{$candidates} ) {
        next if $candidate =~ m/$subject/i;
        push @anagrams, $candidate if _is_anagram( $subject, $candidate );
    }

    return \@anagrams;
}

sub _is_anagram( $subject, $candidate ) {

    if ( length($subject) != length($candidate) ) {
        return 0;
    }

    my %candidate_letters;
    my %subject_letters;
    foreach (split //, $subject) {
        $subject_letters{lc($_)}++;
    };
    foreach (split //, $candidate) {
        $candidate_letters{lc($_)}++;
    };

    foreach my $letter ( keys %subject_letters ) {
        if (grep { $_ eq $letter } keys %candidate_letters) {

            if ($subject_letters{$letter} == 1) {
                delete $subject_letters{$letter};

            } else {
                $subject_letters{$letter}--;
            }

            if ($candidate_letters{$letter} == 1) {
                delete $candidate_letters{$letter};
            } else {
                $candidate_letters{$letter}--;
            }

            return _is_anagram( join('', map { $_ x $subject_letters{$_} } keys %subject_letters), join('', map { $_ x $candidate_letters{$_} } keys %candidate_letters) );

        } else {
            return 0;
        }
    }
    return 1;
}
 