import { throwStatement } from "@babel/types";

//
// This is only a SKELETON file for the 'Protein Translation' exercise. It's been provided as a
// convenience to get you started writing code faster.
//

export const translate = (rna) => {
    var protein_sequence = [];

    if (rna !== undefined) {
        var l = rna.length;

        while (l > 0) {
            let codon = rna.slice(0,3);
            l = l-3;
            rna = rna.slice(-(l));

            switch(codon) {
                case 'AUG':
                    protein_sequence.push('Methionine');
                    break;
                case 'UUU':
                    protein_sequence.push('Phenylalanine');
                    break;
                case 'UUC':
                    protein_sequence.push('Phenylalanine');
                    break;
                case 'UUA':
                    protein_sequence.push('Leucine');
                    break;
                case 'UUG':
                    protein_sequence.push('Leucine');
                    break;
                case 'UCU':
                    protein_sequence.push('Serine');
                    break;
                case 'UCC':
                    protein_sequence.push('Serine');
                    break;
                case 'UCA':
                    protein_sequence.push('Serine');
                    break;
                case 'UCG':
                    protein_sequence.push('Serine');
                    break;
                case 'UAU':
                    protein_sequence.push('Tyrosine');
                    break;
                case 'UAC':
                    protein_sequence.push('Tyrosine');
                    break;
                case 'UGU':
                    protein_sequence.push('Cysteine');
                    break;
                case 'UGC':
                    protein_sequence.push('Cysteine');
                    break;
                case 'UGG':
                    protein_sequence.push('Tryptophan');
                    break;
                case 'UAA': // STOP
                    return protein_sequence;
                case 'UAG': // STOP
                    return protein_sequence;
                case 'UGA': // STOP
                    return protein_sequence;
                default:
                    throw('Invalid codon');
            }
        }
    }

    return protein_sequence;
};
