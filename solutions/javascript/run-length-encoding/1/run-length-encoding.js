
//
// This is only a SKELETON file for the 'Run Length Encoding' exercise. It's been provided as a
// convenience to get you started writing code faster.
//

export const encode = (str) => {

    var encoded_string = '';
    var buffer         = '';
    var last_character = '';

    str.toString().split('').forEach(character => {
        if (character != last_character) {
            // First character handling
            if (last_character == '') {
                buffer = character;
            } else {
                encoded_string = encode_buffer(buffer, encoded_string);
                buffer = character;
            }
        } else {
            buffer = buffer + character;
        }
        last_character = character;
    });

    // Last character handling
    encoded_string = encode_buffer(buffer, encoded_string);

    return encoded_string;
};

function encode_buffer(buffer, encoded_string) {
    var buffer_length = buffer.length;
    if (buffer_length > 1) {
        encoded_string = encoded_string + buffer_length + buffer.substr(0, 1);
    } else {
        encoded_string = encoded_string + buffer;
    }

    return encoded_string;
}

export const decode = (str) => {

    var decoded_string = '';
    var number = '';

    str.toString().split('').forEach(character => {

        var number_regex = /^[0-9]$/;
        // Note the inclusion of whitespace
        var letter_regex = /^[A-Z ]$/i;

        if (character.match(number_regex)) {
            number = number + character;
            character = '';

        } else if (character.match(letter_regex)) {

            if (number > 0) {
                while (number > 0) {
                    decoded_string = decoded_string + character;
                    number = number - 1;
                }
            } else {
                decoded_string = decoded_string + character;
            }
        }
    });

    return decoded_string;
};
