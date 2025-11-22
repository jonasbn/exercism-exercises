=begin
Write your code for the 'Roman Numerals' exercise in this file. Make the tests in
`roman_numerals_test.rb` pass.

To get started with TDD, see the `README.md` file in your
`ruby/roman-numerals` directory.
=end

module RomanNumerals
    def to_roman
        
        positional_sets = [
            {
                "primary"   => 'I',
                "secondary" => 'V',
                "tertiary"  => 'X',
            },
            {
                "primary"   => 'X',
                "secondary" => 'L',
                "tertiary"  => 'C',
            },
            {
                "primary"   => 'C',
                "secondary" => 'D',
                "tertiary"  => 'M',
            },
            {
                "primary" => 'M',
            },
        ];

        digits = self.to_s.chars.reverse

        complete_roman_numeral = '';

        i = 0
        while i <= digits.length
            (remainder, multiplier) = _calculate_remainder_and_multiplier(digits[i].to_i);

            roman_numeral = '';
            set = positional_sets[i];

            # The roman numeral does not require duplication and we can just append it
            if multiplier > 0 and remainder < 4
                roman_numeral = _postfix(roman_numeral, set["secondary"], 1);
            end

            if remainder == 4
                # The roman numeral requires subtraction, meaning we append the highest available
                # numeral in the set and prefix the required amount of the lowest available numeral
                # from the set the required amount of times
                if multiplier > 0
                    roman_numeral = _postfix(roman_numeral, set["tertiary"], 1);
                    roman_numeral = _prefix(roman_numeral, set["primary"], 5 - remainder);

                # The roman numeral requires subtraction, meaning we append the next highest available
                # numeral in the set and prefix the required amount of the lowest available numeral
                # from the set the required amount of times
                else
                    roman_numeral = _postfix(roman_numeral, set["secondary"], 1);
                    roman_numeral = _prefix(roman_numeral, set["primary"], 5 - remainder);
                end
            elsif remainder > 0

                # The roman numeral requires addition meaning we append the lowest available numeral
                # from the set the required amount of times
                roman_numeral = _postfix(roman_numeral, set["primary"], remainder);
            end

            complete_roman_numeral =  roman_numeral + complete_roman_numeral;
            i = i + 1;
        end

        return complete_roman_numeral;
    end

    def _prefix(roman_numeral, character, amount)

        prefix = character * amount;

        roman_numeral = prefix + roman_numeral;

        return roman_numeral;
    end

    def _postfix(roman_numeral, character, amount)

        postfix = character * amount;

        roman_numeral = roman_numeral + postfix

        return roman_numeral;
    end

    def _calculate_remainder_and_multiplier(number)

        remainder = number % 5;
        multiplier = number / 5;

        return remainder, multiplier;
    end                                                             
end
  
class Integer
    include RomanNumerals
end
  