def roman(number):
        
    positional_sets = [
        {
            'primary': 'I',
            'secondary': 'V',
            'tertiary': 'X',
        },
        {
            'primary': 'X',
            'secondary': 'L',
            'tertiary': 'C',
        },
        {
            'primary': 'C',
            'secondary': 'D',
            'tertiary': 'M',
        },
        {
            'primary': 'M',
        },
    ]

    digits = [int(i) for i in str(number)]
    digits.reverse()

    complete_roman_numeral = ''
    position = 0

    for number in digits:
        (remainder, multiplier) = _calculate_remainder_and_multiplier(number)

        roman_numeral = ''
        set = positional_sets[position]
        
        # The roman numeral does not require duplication and we can just append it
        if multiplier > 0 and remainder < 4:
            roman_numeral = _postfix(roman_numeral, set['secondary'], 1)

        if remainder == 4:
            # The roman numeral requires subtraction, meaning we append the highest available
            # numeral in the set and prefix the required amount of the lowest available numeral
            # from the set the required amount of times
            if multiplier > 0:
                roman_numeral = _postfix(roman_numeral, set['tertiary'], 1)
                roman_numeral = _prefix(roman_numeral, set['primary'], 5 - remainder)
            # The roman numeral requires subtraction, meaning we append the next highest available
            # numeral in the set and prefix the required amount of the lowest available numeral
            # from the set the required amount of times
            else:
                roman_numeral = _postfix(roman_numeral, set['secondary'], 1)
                roman_numeral = _prefix(roman_numeral, set['primary'], 5 - remainder)
            
        elif remainder > 0:
            # The roman numeral requires addition meaning we append the lowest available numeral
            # from the set the required amount of times
            roman_numeral = _postfix(roman_numeral, set['primary'], remainder)

        complete_roman_numeral =  roman_numeral + complete_roman_numeral
        
        position = position + 1

    return complete_roman_numeral

def _prefix(roman_numeral, character, amount):
    prefix = character * amount
    roman_numeral = prefix + roman_numeral
    
    return roman_numeral

def _postfix(roman_numeral, character, amount):
    postfix = character * amount
    roman_numeral = roman_numeral + postfix
    
    return roman_numeral

def _calculate_remainder_and_multiplier(number):
    remainder = number % 5
    multiplier = int(number / 5)

    return remainder, multiplier
