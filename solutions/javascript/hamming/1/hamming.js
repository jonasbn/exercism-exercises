export const compute = (left_strand, right_strand) => {


    if (left_strand == "" && right_strand.length > 0) {
        throw new Error("left strand must not be empty");
    }

    if (right_strand == "" && left_strand.length > 0) {
        throw new Error("right strand must not be empty");
    }

    if (left_strand.length != right_strand.length) {
        throw new Error("left and right strands must be of equal length");
    }

    let distance = 0;
    let position = 0;

    left_strand.split('').forEach(character => {
        if (character != right_strand.charAt(position)) {
            distance++;
        }
        position++;
    });

    return distance;
};
