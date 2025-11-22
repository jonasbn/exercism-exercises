package Acronym;

use v5.40;

use Data::Dumper;

use Exporter qw<import>;
our @EXPORT_OK = qw<abbreviate>;

sub abbreviate ($phrase) {
    $phrase =~ s/_//g;
    my @letters = $phrase =~ m/\b{wb}([A-Za-z])/g;

    my @uppercased_letters = map(uc, @letters);

    return join '', @uppercased_letters;
}

1;
