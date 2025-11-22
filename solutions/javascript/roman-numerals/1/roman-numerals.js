
export const toRoman = (number) => {

    number = number.toString();
    let complete_roman_numeral = [];
    let number_as_list = number.split('').reverse();

    for (var i = 0; i < number_as_list.length; i++) {

        let primary   = '';
        let secondary = '';
        let tertiary  = '';

        switch (i) {
            case 0:
                primary   = 'I';
                secondary = 'V';
                tertiary  = 'X';
                break;
            case 1:
                primary   = 'X';
                secondary = 'L';
                tertiary  = 'C';
                break;
            case 2:
                primary   = 'C';
                secondary = 'D';
                tertiary  = 'M'
                break;
            case 3:
                primary   =   'M';
                break;
        }

        let roman_numeral = [];

        let number = number_as_list[i];
        let remainder = number % 5;
        let multiplier = Math.floor(number / 5);

        if (multiplier > 0 && remainder < 4) {
            roman_numeral.push(secondary);
        }
        if (remainder == 4) {
            if (multiplier > 0) {
                roman_numeral.push(tertiary);
                for (let j = 5-remainder; j > 0; j--) {
                    roman_numeral.unshift(primary);
                }
            } else {
                roman_numeral.push(secondary);
                for (let k = 5-remainder; k > 0; k--) {
                    roman_numeral.unshift(primary);
                }
            }
        } else if (remainder > 0) {
            for (let l = remainder; l > 0; l--) {
                roman_numeral.push(primary);
            }
        }

        complete_roman_numeral.unshift(roman_numeral.join(''));
    }
    return complete_roman_numeral.join('');
};
