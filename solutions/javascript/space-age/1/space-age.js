// REF: https://exercism.io/my/solutions/41c45603beea4ec2b9387c72c3448d96

const seconds_in_earth_year = 60 * 60 * 24 * 365.25;

export const PLANETS = {
	mercury: 0.2408467,
	venus:   0.61519726,
	earth:   1,
	mars:    1.8808158,
	jupiter: 11.862615,
	saturn:  29.447498,
	uranus:  84.016846,
	neptune: 164.79132,
};

export const age = (planet, age) => {
    return Math.round(age / seconds_in_earth_year / PLANETS[planet] * 100) / 100;
};
