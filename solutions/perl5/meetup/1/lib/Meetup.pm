package Meetup;

use v5.40;

use Time::Piece;
use Exporter qw<import>;
our @EXPORT_OK = qw<meetup>;

sub meetup ($desc) {

    # Example: 'Teenth Monday of May 2013'
    my ($parsed_descriptor, $parsed_weekday, $parsed_month, $parsed_year) = $desc =~ m/^(\w+) (\w+day) of (\w+) (\d+)$/;

    my $month_number = _resolve_month($parsed_month);
    my $complete_month = _initialize_month($month_number, $parsed_year);
    my $positioner = _resolve_descriptor($parsed_descriptor);
    my $reversed = 0;

    if ($positioner < 0) {
        @{$complete_month} = reverse @{$complete_month};
        $positioner = 1;
        $reversed = 1;
    }

    my $counter = 0;
    my $day_in_month;

    for (my $i = 0; $i < scalar @{$complete_month}; $i++) {
        my $day = $complete_month->[$i];

        if ($day eq $parsed_weekday) {
            $counter++;

            if ($positioner >= 13) {
                if (int($i / 12) >= 1) {
                    print STDERR "$i / 10 = ", int($i / 10), "\n";
                    $day_in_month = $i + 1;
                    last;
                }
            } else {
                if ($reversed) {
                    $day_in_month = scalar @{$complete_month} + 1 - ($i + 1);
                    last;
                } elsif ($positioner == $counter) {
                    $day_in_month = $i + 1;
                    last;
                }
            }
        }
    }

    return sprintf '%04d-%02d-%02d', $parsed_year, $month_number, $day_in_month;
}

sub _resolve_descriptor ($descriptor) {
    my %descriptors = (
        'First' => 1,
        'Second' => 2,
        'Third' => 3,
        'Fourth' => 4,
        'Fifth' => 5,
        'Teenth' => 13,
        'Last' => -1,
    );

    return $descriptors{$descriptor};
}

sub _resolve_month ($month) {
        my %months = (
	    'January' => 1, 
	    'February' => 2,
	    'March' => 3, 
	    'April' => 4, 
	    'May' => 5, 
	    'June' => 6, 
	    'July' => 7, 
	    'August' => 8, 
	    'September' => 9, 
	    'October' => 10, 
	    'November' => 11, 
	    'December' => 12, 
    );

    return $months{$month};
}

sub _initialize_month ($month, $year) {

    my $baseline_time_string = sprintf('01-%02d-%04d', $month, $year);
    my $baseline_time = Time::Piece->strptime($baseline_time_string, '%d-%m-%Y');
    my $month_last_day = $baseline_time->month_last_day;

    my @month = ();
    foreach my $day (1..$month_last_day) {
        my $time_string = sprintf('%02d-%02d-%04d', $day, $month, $year);
        my $time = Time::Piece->strptime($time_string, '%d-%m-%Y');
        push @month, $time->fullday;
    }
    return \@month;
}

1;
