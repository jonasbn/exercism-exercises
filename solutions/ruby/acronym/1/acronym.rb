class Acronym
  def self.abbreviate(phrase)
    return phrase.scan(/\b\w/).join.upcase
  end
end
