# Declare package 'Bob'
package Bob;

use v5.40;

use Exporter qw<import>;
our @EXPORT_OK = qw<hey>;

my $uppercase = qr/\p{Uppercase_Letter}/;
my $lowercase = qr/\p{Lowercase_Letter}/;
my $whitespace = qr/^\s*$/;
my $question = qr/\?+\s*$/;

sub hey ($msg) {
    if ($msg =~ $uppercase && $msg !~ $lowercase && $msg !~ $question) {
        return "Whoa, chill out!";
    } elsif ($msg =~ $uppercase && $msg !~ $lowercase && $msg =~ $question) {
        return "Calm down, I know what I'm doing!";
    } elsif ($msg =~ $question) {
        return "Sure.";
    } elsif ($msg =~ $whitespace) {
        return "Fine. Be that way!";
    } else {
        return "Whatever.";
    }
}

1;
