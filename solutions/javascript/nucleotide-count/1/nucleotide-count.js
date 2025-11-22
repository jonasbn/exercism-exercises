
export class NucleotideCounts {
  static parse(strand) {

    const re = /[^ACGT]/i;
    const invalid_nucleotide_match = strand.match(re);

    if (invalid_nucleotide_match) {
        throw new Error("Invalid nucleotide in strand");
    }

    var a = strand.split("A").length - 1
    var c = strand.split("C").length - 1
    var g = strand.split("G").length - 1
    var t = strand.split("T").length - 1

    return [a, c, g, t].join(' ');
  }
}
