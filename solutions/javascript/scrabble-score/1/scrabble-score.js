export const SCORES = {
    A: 1,
    B: 3,
    C: 3,
    D: 2,
    E: 1,
    F: 4,
    G: 2,
    H: 4,
    I: 1,
    J: 8,
    K: 5,
    L: 1,
    M: 3,
    N: 1,
    O: 1,
    P: 3,
    Q: 10,
    R: 1,
    S: 1,
    T: 1,
    U: 1,
    V: 4,
    W: 4,
    X: 8,
    Y: 4,
    Z: 10
};

export const score = (word, double_word = false, triple_word = false, double_character_positions = [], triple_character_positions = []) => {
    let accumulated_score= 0
    let position = 0;

    word.split('').forEach(character => {
        let score = SCORES[character.toUpperCase()];

        if (double_character_positions && double_character_positions.includes(position)) {
            score = double(score)
        }
        if (triple_character_positions && triple_character_positions.includes(position)) {
            score = triple(score)
        }

        accumulated_score += score;
        position++;
    });

    if (double_word) {
        accumulated_score = double(accumulated_score)
    }

    if (triple_word) {
        accumulated_score = triple(accumulated_score)
    }

    return accumulated_score
};

function double(score) {
    return score *= 2
}

function triple(score) {
    return score *= 3
}
