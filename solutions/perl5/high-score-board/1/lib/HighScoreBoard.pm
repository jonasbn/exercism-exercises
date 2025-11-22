package HighScoreBoard;

use v5.38;

our %Scores;

sub set_player_scores (%new_scores) {
    $Scores{$_} = $new_scores{$_} for keys %new_scores;

    return 1;
}

sub get_player_score ($player) {
    return $Scores{$player};
}

sub increase_player_scores (%additional_scores) {
    $Scores{$_} += $additional_scores{$_} for keys %additional_scores;

    return 1;
}

sub sort_players_by_name {

    # REF: https://perldoc.perl.org/functions/sort
    my @players = sort keys %Scores;

    return @players;
}

sub sort_players_by_score {

    # REF: https://perldoc.perl.org/functions/sort
    my @players = sort { $Scores{$b} <=> $Scores{$a} } keys %Scores;

    return @players;
}

sub delete_player ($player) {
    return delete $Scores{$player};
}
