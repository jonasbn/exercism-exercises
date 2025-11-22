package KindergartenGarden;

use v5.40;

use Exporter qw<import>;
our @EXPORT_OK = qw<plants>;

sub plants ( $diagram, $student ) {

    my %plants = (
        C => 'clover',
        G => 'grass',
        R => 'radishes',
        V => 'violets',
    );

    my $position = _resolve_student($student);

    my @rows = split /\n/, $diagram;
    my @planted_student_plants = ();

    foreach my $row (@rows) {
        my @planted_plants = split //, $row;

        my $plant = $plants{$planted_plants[$position]};
        push @planted_student_plants, $plant;        
        $plant = $plants{$planted_plants[$position+1]};
        push @planted_student_plants, $plant;
    }

    return \@planted_student_plants;
}

sub _resolve_student {
	my $student = shift;

    my @children = qw(Alice Bob Charlie David Eve Fred Ginny Harriet Ileana Joseph Kincaid Larry);
    my $position = 0;

    foreach my $child (@children) {
        if ($child eq $student) {
            last;
        }
		$position += 1;
    }

	if ($position > 0) {
		$position *= 2;
	}

    return $position;
}

1;
