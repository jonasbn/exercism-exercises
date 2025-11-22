package Accumulate;

use v5.38;

use Exporter qw<import>;
our @EXPORT_OK = qw<accumulate>;

sub accumulate ( $list, $func ) {
    my @new_list;
    foreach my $item ( @{$list} ) {
        push @new_list, $func->($item);
    }
    return \@new_list;
}
