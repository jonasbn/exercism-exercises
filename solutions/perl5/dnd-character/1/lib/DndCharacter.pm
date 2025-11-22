use v5.40;
use experimental qw<class>;

use POSIX qw/floor/;

class DndCharacter {

    field $strength :reader;
    field $dexterity :reader;
    field $constitution :reader;
    field $intelligence :reader;
    field $wisdom :reader;
    field $charisma :reader;

    ADJUST {
        $strength     = ability();
        $dexterity    = ability();
        $constitution = ability();
        $intelligence = ability();
        $wisdom       = ability();
        $charisma     = ability();
    }

    sub hitpoints {
        my $self = shift;
        return 10 + DndCharacter->modifier($self->constitution);
    }

    sub ability {
        my $self = shift;

    	my @throws = ();

    	for my $i (1 .. 4) {
    		push @throws, DndCharacter->_throw_die();
    	}

    	my @sorted = sort { $a <=> $b } @throws;

    	my $ability = $sorted[0] + $sorted[1] + $sorted[2];

    	return $ability
    }

    sub _throw_die {
    	return int(rand(6)) + 1;
    }

    sub modifier {
        my $class = shift;
        my $score = shift;
        return floor(($score - 10) / 2);
    }
}
1;
