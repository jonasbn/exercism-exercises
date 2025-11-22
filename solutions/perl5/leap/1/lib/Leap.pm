# Declare package 'Leap'
package Leap;

use v5.38;

use constant {
    TRUE => 1,
    FALSE => 0,
};

use Exporter qw<import>;
our @EXPORT_OK = qw<is_leap_year>;

sub is_leap_year ($year) {
    if ($year % 4 == 0) {
        if ($year % 100 == 0) {
            if ($year % 400 == 0) {
                return TRUE;
            } else {
                return FALSE;
            }
        } else {
            return TRUE;
        }
    } else {
        return FALSE;
    }
}
