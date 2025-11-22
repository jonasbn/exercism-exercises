use v5.40;
use experimental qw<class>;
use Data::Dumper;

class GradeSchool;
 
field %roster = ();

method add ( $student, $grade ) {
    if (not exists $roster{$student}) {
        $roster{$student} = $grade;
    }
}

method roster ( $grade = undef ) {

    my @sorted = ();
    my @students = ();

    if ($grade) {
        foreach my $student (keys %roster) {
            if ($roster{$student} == $grade) {
                push @students, $student;
            }
        }
        @sorted = sort @students;

    } else {

        @sorted = sort { $roster{$a} <=> $roster{$b}
                                     ||
                                  $a cmp $b
        } keys %roster;
    }

    return \@sorted;
}

1;
