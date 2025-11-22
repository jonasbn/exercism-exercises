use v5.40;
use experimental qw<class>;

# REF:
# - https://fluca1978.github.io/2023/07/20/PerlClassSupport.html
# - https://andrewshitov.com/2023/07/06/built-in-classes-in-perl-5-38/
# - https://docs.mojolicious.org/perlclass

class Robot;

field $x :param;
field $y :param;
field $direction :param;

method direction ($d = undef) {
    if (defined $d) {
        $direction = $d;
    }
    return $direction;
}

method x ($x_param = undef) {
    if (defined $x_param) {
        $x = $x_param;
    }
    return $x;
}

method y ($y_param = undef) {
    if (defined $y_param) {
        $y = $y_param;
    }
    return $y;
}

method enact ($instruction_string) {

    my @instructions = split //, $instruction_string;

    for my $instruction (@instructions) {

        if ($instruction eq 'A') {
            if ($self->direction eq 'north') {
                $self->y($self->y + 1);
            } elsif ($self->direction eq 'east') {
                $self->x($self->x + 1);
            } elsif ($self->direction eq 'south') {
                $self->y($self->y - 1);
            } elsif ($self->direction eq 'west') {
                $self->x($self->x - 1);
            }
        } elsif ($instruction eq 'R') {
            if ($self->direction eq 'north') {
                $self->direction('east');
            } elsif ($self->direction eq 'east') {
                $self->direction('south');
            } elsif ($self->direction eq 'south') {
                $self->direction('west');
            } elsif ($self->direction eq 'west') {
                $self->direction('north');
            }
        } elsif ($instruction eq 'L') {
            if ($self->direction eq 'north') {
                $self->direction('west');
            } elsif ($self->direction eq 'west') {
                $self->direction('south');
            } elsif ($self->direction eq 'south') {
                $self->direction('east');
            } elsif ($self->direction eq 'east') {
                $self->direction('north');
            }
        }
    }
    return $self;
}

1;
