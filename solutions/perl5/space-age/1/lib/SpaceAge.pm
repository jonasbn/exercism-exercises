package SpaceAge;

use v5.40;

use Exporter qw<import>;
our @EXPORT_OK = qw<age_on_planet>;

use constant SECONDS_IN_EARTH_YEAR => 31557600;

my %planets = (
    'Mercury' => 0.2408467,
    'Venus'   => 0.61519726,
    'Earth'   => 1,
    'Mars'    => 1.8808158,
    'Jupiter' => 11.862615,
    'Saturn'  => 29.447498,
    'Uranus'  => 84.016846,
    'Neptune' => 164.79132,    
);

sub age_on_planet ( $planet, $seconds ) {
    die "not a planet" unless exists $planets{$planet};
    # Use sprintf to round instead of POSIX
    # REF: https://stackoverflow.com/questions/178539/how-do-you-round-a-floating-point-number-in-perl
    return sprintf('%.2f', ((($seconds / SECONDS_IN_EARTH_YEAR)) / $planets{$planet}));
}

1;
