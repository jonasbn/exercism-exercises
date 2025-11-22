package MatchingBrackets;

use v5.40;

use Exporter qw<import>;
our @EXPORT_OK = qw<has_matching_brackets>;

sub has_matching_brackets ($text) {
    my @characters = split //, $text;

    my @stack;

    for my $char (@characters) {
        if ($char =~ m/\(/) {
            push @stack, 'parenthesis';
        } elsif ($char =~ m/\[/) {
            push @stack, 'squarebracket';
        } elsif ($char =~ m/{/) {
            push @stack, 'curlybracket';
        } elsif ($char =~ m/\)/) {
            my $previous_bracket = pop @stack || '';
            if ($previous_bracket ne 'parenthesis') {
                return 0;
            }
        } elsif ($char =~ m/\]/) {
            my $previous_bracket = pop @stack || '';
            if ($previous_bracket ne 'squarebracket') {
                return 0;
            }

        } elsif ($char =~ m/}/) {
            my $previous_bracket = pop @stack || '';
            if ($previous_bracket ne 'curlybracket') {
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
