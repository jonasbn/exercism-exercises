
export const gigasecond = (date) => {

    // Adding 1 gigasecond in miliseconds
    // to adhere to Date constructor requirements

    const epoch = Date.UTC(
        date.getUTCFullYear(),
        date.getUTCMonth(),
        date.getUTCDate(),
        date.getUTCHours(),
        date.getUTCMinutes(),
        date.getUTCSeconds()
    ) + 1e12;

    return new Date(epoch);
};
