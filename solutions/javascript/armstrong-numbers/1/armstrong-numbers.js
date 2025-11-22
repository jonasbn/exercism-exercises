//
// This is only a SKELETON file for the 'Armstrong numbers' exercise. It's been provided as a
// convenience to get you started writing code faster.
//

export const validate = (number) => {
    var ar = number.toString().split('');

    var l = ar.length;
    var sum = 0;

    for (var i = 0; i < l; i++) {
        sum += ar[i] ** l;
    };

    if (number == sum) {
        return true;
    } else {
        return  false;
    }
};
