=begin
Write your code for the 'Atbash Cipher' exercise in this file. Make the tests in
`atbash_cipher_test.rb` pass.

To get started with TDD, see the `README.md` file in your
`ruby/atbash-cipher` directory.
=end

module Atbash

    def self.encode(phrase)
        plain = 'abcdefghijklmnopqrstuvwxyz'
        cipher = plain.reverse

        encoded_phrase = ''

        for c in phrase.chars do
            if c.match(/[\., ]/)
                next
            end

            if c.match(/\d+/)
                e = c
            else
                c.downcase!
                p = plain.index(c)
                e = cipher[p]
            end
            encoded_phrase.concat(e)
        end
        return format(encoded_phrase)
    end

    def self.format(phrase)
        p = 0
        l = 5
        formatted_phrase = ''

        while p < phrase.length
            s = phrase[p, l]
            p = p + l
            formatted_phrase.concat(s)
            if p < phrase.length
                formatted_phrase.concat(" ")
            end
        end
        return formatted_phrase
    end

    def self.decode(phrase)
        plain = 'abcdefghijklmnopqrstuvwxyz'
        cipher = plain.reverse

        decoded_phrase = ''

        for c in phrase.chars do
            if c.match(/[\., ]/)
                next
            end

            if c.match(/\d+/)
                d = c
            else
                c.downcase!
                p = cipher.index(c)
                d = plain[p]
            end
            decoded_phrase.concat(d)
        end

        return decoded_phrase
    end
end
