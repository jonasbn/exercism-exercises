//
// This is only a SKELETON file for the 'Triangle' exercise. It's been provided as a
// convenience to get you started writing code faster.
//

export class Triangle {
    constructor(a, b, c) {
        this.a = a;
        this.b = b;
        this.c = c;
    }

    kind() {

        if (this.a <= 0 || this.b <= 0 || this.c <= 0) {
            throw('Sides cannot be of length zero');
        }

        if (this.a + this.b < this.c
            || this.a + this.c < this.b
            || this.b + this.c < this.a) {
                throw('Triangle not equal');
            }

        if (this.a == this.b && this.a == this.c && this.b == this.c) {
            return 'equilateral';
        }

        if (this.a == this.b || this.a == this.c || this.b == this.c) {
            return 'isosceles';
        }

        if (this.a != this.b && this.a != this.c && this.b != this.c) {
            return 'scalene';
        }
    }
}
