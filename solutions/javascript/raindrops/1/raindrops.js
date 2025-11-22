export const convert = (number) => {
    let rain = "";

    if (number % 3 == 0) {
        rain += "Pling"
    }

    if (number % 5 == 0) {
        rain += "Plang"
    }

    if (number % 7 == 0) {
        rain += "Plong"
    }

    if (rain.length == 0) {
        rain = number.toString();
    }

    return rain;
};
