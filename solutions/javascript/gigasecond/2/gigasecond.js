
export const gigasecond = (date) => {

    // Adding 1 gigasecond in miliseconds
    // to adhere to Date constructor requirements

    return new Date(date.getTime() + 1e12);
};
