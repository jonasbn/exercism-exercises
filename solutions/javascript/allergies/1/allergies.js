
export const ALLERGEN_SCORES = {
    1: "eggs",
    2: "peanuts",
    4: "shellfish",
    8: "strawberries",
    16: "tomatoes",
    32: "chocolate",
    64: "pollen",
    128: "cats",
};

export class Allergies {
    constructor(score) {
        this.score = score;
        this.allergens = []

        Object.keys(ALLERGEN_SCORES).sort(function(a, b) { return b - a }).forEach(score => {
            if (this.score >= score) {
                var remainder = this.score % score
                this.score = remainder
                this.allergens.push(ALLERGEN_SCORES[score])
            }
        });
    }

    list() {
        return this.allergens.reverse()
    }

    allergicTo(allergen) {
        if (this.allergens.includes(allergen)) {
            return true
        }

        return false
    }
}
