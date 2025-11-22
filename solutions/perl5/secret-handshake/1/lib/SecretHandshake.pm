package SecretHandshake;

use v5.40;

use Data::Dumper;
use Exporter qw<import>;
our @EXPORT_OK = qw<handshake>;

my %actions = (
  'wink'            => 1 << 0,
  'double blink'    => 1 << 1,
  'close your eyes' => 1 << 2,
  'jump'            => 1 << 3,
  'reverse'         => 1 << 4,
);

my @actions_order = ('wink', 'double blink', 'close your eyes', 'jump', 'reverse');

sub handshake ($number) {
    my $bit_field = int $number;
    
    my @handshake = ();

    for my $action (@actions_order) {
        if ($actions{ $action } & $bit_field) {
            if ($action eq 'reverse') {
                @handshake = reverse @handshake;
            } else {
                push @handshake, $action;
            }
        }
    }
    return \@handshake;
}

1;
