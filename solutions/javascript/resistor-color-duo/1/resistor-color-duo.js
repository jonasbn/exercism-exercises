const MAX_NUMBER_OF_COLORS = 2;

export const decodedValue = (colors) => {
    let value = [];

    for (let i = 0; i < MAX_NUMBER_OF_COLORS; i++) {
        value.push(COLORS.indexOf(colors[i]));
    };

    return parseInt(value.join(''));
};

export const COLORS = ["black","brown","red","orange","yellow","green","blue","violet","grey","white"];
