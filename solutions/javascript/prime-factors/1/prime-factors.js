
export const primeFactors = (number, divisor = 2) => {
    var primefactors = [];

    while (number > 1) {
        var remainder = number % divisor;
        if (remainder === 0) {
            primefactors.push(divisor);
            number = number / divisor;
        } else {
            divisor = divisor + 1;
        }
    }

    return primefactors;
};
