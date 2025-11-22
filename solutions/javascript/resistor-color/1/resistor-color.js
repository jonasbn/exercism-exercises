
export const colorCode = (color) => {
    // Using linear search
    for (var i = 0; i < COLORS.length; i++) {
        if (color === COLORS[i]) {
            return i;
        }
    }
};

export const COLORS = ["black","brown","red","orange","yellow","green","blue","violet","grey","white"];
