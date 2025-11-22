package AtbashCipher;

use v5.38;
our $VERSION = '0.01';

use Exporter qw<import>;
our @EXPORT_OK = qw<encode_atbash decode_atbash>;

my $plain = 'abcdefghijklmnopqrstuvwxyz';
my $cipher = reverse $plain;

sub encode_atbash ($phrase) {

    my $encoded_phrase = q();

    foreach my $character (split //, $phrase) {

        if ($character =~ m/[\., ]/) {
            next;
        };

        my $encoded_character = '';

        if ($character =~ m/[\d]/) {
            $encoded_character = $character;
        } else {
            my $position = index $plain, lc $character;

            $encoded_character = substr $cipher, $position, 1;
        }

        $encoded_phrase .= $encoded_character;
    }

    my $encoded_and_formatted_phrase = _format_cipher($encoded_phrase);

    return $encoded_and_formatted_phrase;
}

sub decode_atbash ($phrase) {

    my $decoded_phrase = q();

    foreach my $character (split //, $phrase) {

        my $decoded_character = '';

        if ($character =~ m/[\., ]/) {
            next;
        };

        if ($character =~ m/[\d]/) {
            $decoded_character = $character;
        } else {
            my $position = index $cipher, lc $character;

            $decoded_character = substr $plain, $position, 1;
        }

        $decoded_phrase .= $decoded_character;
    }

    return $decoded_phrase;

}

sub _format_cipher ($phrase) {

    my $group_length = 5;
    my $formatted_phrase = '';

    my $position = 0;
    while (my $group = substr $phrase, $position, $group_length) {

        if ($position + $group_length > length $phrase) {
            $group_length = length($phrase) - $position;
        }

        $formatted_phrase .= $group;
        $formatted_phrase .= ' ';

        $position += $group_length;
    }

    chop $formatted_phrase;

    return $formatted_phrase;
}

1;
