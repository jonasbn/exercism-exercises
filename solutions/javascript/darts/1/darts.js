
export const solve = (x, y) => {

    let distance = calculate_distance(x, y);

    if (distance <= 1) {
        return 10;
    } else if (distance <= 5) {
        return 5;
    } else if (distance <= 10) {
        return 1;
    }
    return 0;
};

function calculate_distance(x, y) {
    // REF: https://en.wikipedia.org/wiki/Distance
    // REF: https://byjus.com/maths/distance-between-points/
    return Math.sqrt(Math.pow(x, 2) + Math.pow(y, 2));
}
