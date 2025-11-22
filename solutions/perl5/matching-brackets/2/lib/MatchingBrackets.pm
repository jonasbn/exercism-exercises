package MatchingBrackets;

use v5.40;

use Exporter qw<import>;
our @EXPORT_OK = qw<has_matching_brackets>;

use constant {
    PARENTHESIS    => 'parenthesis',
    SQUAREBRACKET  => 'squarebracket',
    CURLYBRACKET   => 'curlybracket',
};

sub has_matching_brackets ($text) {
    my @characters = split //, $text;

    my @stack;

    for my $char (@characters) {
        if ($char =~ m/\(/) {
            push @stack, PARENTHESIS;
        } elsif ($char =~ m/\[/) {
            push @stack, SQUAREBRACKET;
        } elsif ($char =~ m/{/) {
            push @stack, CURLYBRACKET;
        } elsif ($char =~ m/\)/) {
            my $previous_bracket = pop @stack || '';
            if ($previous_bracket ne PARENTHESIS) {
                return 0;
            }
        } elsif ($char =~ m/\]/) {
            my $previous_bracket = pop @stack || '';
            if ($previous_bracket ne SQUAREBRACKET) {
                return 0;
            }

        } elsif ($char =~ m/}/) {
            my $previous_bracket = pop @stack || '';
            if ($previous_bracket ne CURLYBRACKET) {
                return 0;
            }

        } else {
            next;
        }
    }

    if (scalar @stack == 0) {
        return 1;
    } else {
        return 0;
    }
}

1;
