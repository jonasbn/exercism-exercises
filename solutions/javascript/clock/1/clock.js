//
// This is only a SKELETON file for the 'Clock' exercise. It's been provided as a
// convenience to get you started writing code faster.
//

const hours_in_day = 24;
const minutes_in_hour = 60;
const minutes_in_day = hours_in_day * minutes_in_hour;

export class Clock {

  constructor(hours = 0, minutes = 0) {

    // if hours exceed a day we calculate the number of hours exceeding 24 hours
    if (hours >= hours_in_day) {
        hours = trim_hours_to_day(hours);
    }

    // if negaive hours exceed a day we calculate the number of hours exceeding 24 hours and subtract
    if (hours <= -(hours_in_day)) {
        hours = hours_in_day + trim_hours_to_day(hours);
    }

    // if negative hours between below 24 hours, we calculate the remainder
    if (hours < 0) {
        hours = hours_in_day + hours;
    }

    // we calculate how many minutes from hours and add the specified minutes (trimmed)
    this.minutes = hours * minutes_in_hour + trim_minutes_to_day(minutes);
  }

  toString() {
    let hours = this.calculateHours();
    let minutes = this.minutes % minutes_in_hour;

    return leftFillNumber(hours) + ':' + leftFillNumber(minutes);
  }

  plus(minutes) {
    this.minutes += trim_minutes_to_day(minutes);
    return this;
  }

  minus(minutes) {
    this.minutes -= trim_minutes_to_day(minutes);
    return this;
  }

  equals(clock) {
    return this.toString() === clock.toString();
  }

  calculateHours() {
    // if we are negative, we subtract from minutes in a day
    if (this.minutes < 0) {
        this.minutes += minutes_in_day;
    }

    // calculate hours
    let hours = Math.floor(this.minutes / minutes_in_hour);

    // If we exceed the number of hours in a day, we trim to 24 hours
    if (hours >= hours_in_day) {
        hours = trim_hours_to_day(hours);
    }
    return hours;
  }
}

// this trims the amount of hours so only 24-hours are represented
function trim_hours_to_day(hours) {
    return hours % hours_in_day;
}

// this trims the amount of minutes so only 24-hours are represented
function trim_minutes_to_day(minutes) {
    return minutes % minutes_in_day;
}

// this prefixes with zero up to length of 2
function leftFillNumber(number, targetLength = 2) {
    return number.toString().padStart(targetLength, 0);
}
