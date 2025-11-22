// Exercism assignment
// REF: https://exercism.io/my/solutions/17ad4fd5278741bf85b87802966faa6b

export const TRANSCRIPTIONS = { G: 'C', C: 'G', T: 'A', A: 'U' };

export const toRna = (dna) => {

    return dna.split('').map(nucleotide => TRANSCRIPTIONS[nucleotide]).join('');
}
