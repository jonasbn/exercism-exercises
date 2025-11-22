//
// This is only a SKELETON file for the 'Meetup' exercise. It's been provided as a
// convenience to get you started writing code faster.
//

export const meetupDay = (year, month, dayname, nickname) => {

    var last_day_of_month = calculate_last_date_of_month(year, month);
    var array_of_days_in_month = [];

    for (let index = last_day_of_month; index > 0 ; index--) {
        var date = new Date(Date.UTC(year, month, index));
        var options = { weekday: 'long'};
        var name = new Intl.DateTimeFormat('en-US', options).format(date);

        array_of_days_in_month.unshift({ dayname: name, date: date });
    }

    // Find all days matching our dayname
    var days = array_of_days_in_month.filter(element => element.dayname == dayname);

    var index;

    // Handle aliases and order
    if (nickname == 'first' || nickname == '1st') {
        index = 0;
    } else if (nickname == 'second' || nickname == '2nd') {
        index = 1;
    } else if (nickname == 'third' || nickname == '3rd') {
        index = 2;
    } else if (nickname == 'fourth' || nickname == '4th') {
        index = 3;
    } else if (nickname == 'fifth' || nickname == '5th') {
        index = 4;
    } else if (nickname == 'last') {
        index = days.length - 1;
    }

    var day;

    // We just have to find a date based on order
    if (index != undefined) {
        day = days[index];

    // We have to find one of the made up day aliases
    } else {
        day = days.find(element => element.date.getDate() > 12 && element.date.getDate() < 20);
    }

    // Return just the date part from the structure
    return day.date;
}

function calculate_last_date_of_month(year, month) {
    // initiate a date a the first day of a given  month in a given year
    var date = new Date(Date.UTC(year, month, 1));

    // jump forward one month to the following month
    date.setMonth(date.getMonth() + 1);

    // jump back 24 hours so we end on the day before the first day of the following month
    date.setHours(date.getHours() - 24);

    return date.getDate();
}
