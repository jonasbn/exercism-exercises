package RunLengthEncoding;

use v5.40;

use Exporter qw<import>;
our @EXPORT_OK = qw<encode decode>;
use Data::Dumper;

sub encode ($string) {
    my $encoded_string = '';

    my @characters = split //, $string;

    my $count = 1;
    for (my $i = 0; $i < scalar @characters; $i++) {

        my $character = $characters[$i];
        my $next_character = $characters[$i+1] // '';
        
        if ($next_character eq $character) {
            $count++;
            next;
        } else {
            if ($count > 1) {
                $encoded_string .= $count;
            }

            $encoded_string .= $character;
            $count = 1;
        }
    }

    return $encoded_string
}

sub decode ($string) {
    my $decoded_string = '';

    my @sections = $string =~ m/([[:alpha:]]{1}|\d+|\s*)\n?/g;
    my $count = 1; 

    foreach my $section (@sections) {

        if ($section =~ m/\d/) {
            $count = $section;
            next;
        } else {
            $decoded_string .= $section x $count;
        }
        $count = 1; 
    }

    return $decoded_string;
}

1;
