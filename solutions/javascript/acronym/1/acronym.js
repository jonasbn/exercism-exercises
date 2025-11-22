//
// This is only a SKELETON file for the 'Acronym' exercise. It's been provided as a
// convenience to get you started writing code faster.
//

export const parse = (str) => {

  const re = /\s+|-/;
  var acronym = '';
  str = str.replace(/_/g, '');

  str.toString().split(re).forEach(character => {
      acronym = acronym + character.charAt(0).toUpperCase();
  });

  return acronym;
}
