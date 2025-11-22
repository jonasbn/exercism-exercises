module Translation
  def self.of_rna(strand)
    proteins = []

    position = 0
    codon_length = 3

    while position <= strand.length()

      codon = strand[position,codon_length]
      protein = self.translate(codon)

      if protein != ""
        if protein == "STOP"
          break
        else
          proteins.push(protein)
        end
      end
      position += codon_length
    end
    return proteins
  end

  def self.translate(codon)

    protein = ""
    case codon
    when ""
      protein = ""
    when "AUG"
      protein = "Methionine"
    when "UUU", "UUC"
      protein = "Phenylalanine"
    when "UUA", "UUG"
      protein = "Leucine"
    when "UCU", "UCC", "UCA", "UCG"
      protein = "Serine"
    when "UAU", "UAC"
      protein = "Tyrosine"
    when "UGU", "UGC"
      protein = "Cysteine"
    when "UGG"
      protein = "Tryptophan"
    when "UAA", "UAG", "UGA"
      protein = "STOP"
    else
      raise InvalidCodonError.new "InvalidCodonError"
    end

    return protein
  end
end

class InvalidCodonError < StandardError
end
