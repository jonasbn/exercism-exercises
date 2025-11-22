
export const transform = (oldData) => {
    let newData = new Object()

    for (let [key, value] of Object.entries(oldData)) {
        value.forEach(element => {
            newData[element.toLowerCase()] = parseInt(key);
        });
    }

    return newData;
};
