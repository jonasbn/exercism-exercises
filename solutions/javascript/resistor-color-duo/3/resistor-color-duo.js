const MAX_NUMBER_OF_COLORS = 2;

export const decodedValue = (colors) => {

    const values = colors.slice(0, MAX_NUMBER_OF_COLORS).map(color => COLORS.indexOf(color));

    return Number(values.join(''));
};

export const COLORS = ["black","brown","red","orange","yellow","green","blue","violet","grey","white"];
