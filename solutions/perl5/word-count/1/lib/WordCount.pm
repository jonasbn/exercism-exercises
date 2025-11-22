package WordCount;

use v5.40;

use Exporter qw<import>;
our @EXPORT_OK = qw<count_words>;

sub count_words ($sentence) {

    $sentence =~ s/[\s,]{2,}/ /g;
    $sentence =~ s/^\s+|\s+$//g;
    $sentence =~ s/\n//g;

    my @words = split /[\s,]+/, $sentence;
    my %word_count;

    foreach my $word (@words) {
        $word =~ s/^\W+|\W+$//g;
        $word =~ s/^'|'$//g;
        $word = lc $word;
        $word_count{$word}++;
    }
    return \%word_count;
}

1;
