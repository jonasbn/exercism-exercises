// Exercism assignment
// REF: https://exercism.io/my/solutions/17ad4fd5278741bf85b87802966faa6b

export const toRna = (dna) => {

    function translate(nucleotide) {
        switch (nucleotide) {
            case 'G':
                return 'C';
            case 'C':
                return 'G';
            case 'T':
                return 'A';
            case 'A':
                return 'U';
            default:
                break;
        }
    }

    return dna.split('').map(nucleotide => translate(nucleotide)).join('');
}
