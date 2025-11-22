def proteins(strand):
    
    proteins = []

    start = 0
    end = 3
    
    while end <= len(strand):
        codon = strand[start:end]
        protein = translate(codon)

        if protein == "STOP":
            break

        proteins.append(protein)
        start = end
        end += 3;

    return proteins

def translate(codon):

    protein = ""

    match codon:
        case "AUG":
            protein = "Methionine"
        case "UUU":
            protein = "Phenylalanine"
        case "UUC":
            protein = "Phenylalanine"
        case "UUA":
            protein = "Leucine"
        case "UUG":
            protein = "Leucine"
        case "UCU":
            protein = "Serine"
        case "UCC":
            protein = "Serine"
        case "UCA":
            protein = "Serine"
        case "UCG":
            protein = "Serine"
        case "UAU":
            protein = "Tyrosine"
        case "UAC":
            protein = "Tyrosine"
        case "UGU":
            protein = "Cysteine"
        case "UGC":
            protein = "Cysteine"
        case "UGG":
            protein = "Tryptophan"
        case "UAA":
            protein = "STOP"
        case "UAG":
            protein = "STOP"
        case "UGA":
            protein = "STOP"
        case _:
            protein = ""
    
    return protein