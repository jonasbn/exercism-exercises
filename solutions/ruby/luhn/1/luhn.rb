module Luhn
  def self.valid?(s)

    s.gsub!(/\s/, '')

    a = s.split("").reverse()

    sum = 0

    if a.length <= 1
      return false
    end

    position = 0
    for number in a
      if number.match(/\D/)
        return false
      end

      digit = number.to_i
      product = 0

      if position % 2 != 0
        product = digit * 2

        if product > 9
          product -= 9
        end
      else
        product = digit
      end

      sum = sum + product
      position += 1
    end

    if sum % 10 == 0
      return true
    else
      return false
    end

  end
end
