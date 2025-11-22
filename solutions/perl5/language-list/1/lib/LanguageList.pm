package LanguageList;

use v5.38;

our @Languages;

sub add_language ($language) {
    push @Languages, $language;
}

sub remove_language () {
    delete $Languages[-1];
}

sub first_language () {
    return $Languages[0];
}

sub last_language () {
    return $Languages[-1];
}

sub get_languages (@elements) {

    my @l;
    # $[ is the index of the first element in an array, by default 0
    # change is deprecated since perl 5.30
    # REF: https://perldoc.perl.org/perlvar#$%5B     
    foreach my $element (@elements) {
        push @l, $Languages[$element-1];
    }
    
    return @l;
}

sub has_language ($language) {
    return grep { $_ eq $language } @Languages;
}
