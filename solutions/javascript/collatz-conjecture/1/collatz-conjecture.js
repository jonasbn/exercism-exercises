export const steps = (n, s = 0) => {
    if (n <= 0) { throw "Only positive numbers are allowed"; }
    if (n == 1) { return s; }
    if (n % 2 == 0) {
        n = n / 2;
    } else {
        n = n * 3 + 1;
    }
    return steps(n, ++s);
};
