package PigLatin;

use v5.40;

use Exporter qw<import>;
our @EXPORT_OK = qw<translate>;

sub translate ($phrase) {

    my $vowels = 'aeiou';
    my $consonants = 'bcdfghjklmnpqrstvwxyz';

    my $translation = '';
    my @words = $phrase =~ m/(\w*|\s*)\n?/g;

    foreach my $word (@words) {
        
        if ($word =~ /\A([$vowels]|yt|xr)/) { # rule 1: words that start with a vowel sound
            $translation .= $word . 'ay';
        } elsif ($word =~ /\A[$consonants]+y/) { # rule 4: words that start with a consonant sound followed by 'y'
            $word =~ s/\A([$consonants]+)y(.*)/$2$1/;
            $translation .= 'y' . $word . 'ay';
        } elsif ($word =~ /\A[$consonants]*qu/) { # rule 3: words that start with a consonant sound followed by 'qu'
            $word =~ s/\A([$consonants]*qu)(.*)/$2$1/;
            $translation .= $word . 'ay';
        } elsif ($word =~ /\A[$consonants]+/) { # rule 2: words that start with a consonant sound
            $word =~ s/\A([$consonants]+)(.*)/$2$1/;
            $translation .= $word . 'ay';
        } else {
            $translation .= $word;
        }
    }

    return $translation;
}

1;
