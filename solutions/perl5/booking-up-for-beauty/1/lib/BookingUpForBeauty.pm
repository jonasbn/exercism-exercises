package BookingUpForBeauty;

use v5.38;

# Suggested datetime modules you can use:
#use Time::Piece;
#use DateTime::Tiny;

# Recommended, commented out for portability.
use Readonly;

use Exporter ('import');
our @EXPORT_OK = qw(appointment_has_passed is_afternoon_appointment describe_appointment);

my $STRPTIME_FORMAT = '%Y-%m-%d' . 'T' . '%H:%M:%S';
Readonly::Scalar $STRPTIME_FORMAT => $STRPTIME_FORMAT;

# Private subroutines conventionally start with an underscore.
# It isn't necessary, but provided for convenience.
sub _parse_datetime ($date_string) {
    my $t = Time::Piece->strptime($date_string, $STRPTIME_FORMAT);
    return $t;
}

sub appointment_has_passed ($date_string) {
    my $time = _parse_datetime($date_string);
    return $time->epoch < time() ? 1 : 0;
}

sub is_afternoon_appointment ($date_string) {
    my $time = _parse_datetime($date_string);
    
    if ($time->hour >= 12 && $time->hour < 18) {
        return 1;
    } else {
        return 0;
    }
}

sub describe_appointment ($date_string) {
    my $time = _parse_datetime($date_string);

    my $day_indicator = 'AM';
    my $hour = $time->hour;

    if ($time->hour >= 12) {
        $hour -= 12;
        $day_indicator = 'PM';
    } elsif ($time->hour == 0) {
        $hour = 12;
    }

    return sprintf('You have an appointment on %02d/%02d/%04d %d:%02d %s', $time->mon, $time->mday, $time->year, $hour, $time->min, $day_indicator);
}
