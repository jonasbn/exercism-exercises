package Isogram;

use v5.40;

use Exporter qw<import>;
our @EXPORT_OK = qw<is_isogram>;

sub is_isogram ($phrase) {
    my @characters = grep { $_ =~ /\w/ } split //, lc $phrase;
    my %seen;

    for my $char (@characters) {
        return 0 if $seen{$char}++;
    }

    return 1;
}

1;
