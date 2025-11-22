export const isPangram = str => {
  var characters = str.toString()
  .split('');

  const counts = new Map();

  characters.forEach(element => {
    if (isLetter(element)) {
      counts.set(element.toLowerCase(), element);
    }
  });

  if (counts.size == 26) {
    return true
  }

  return false
};

 
function isLetter(charVal)
{
  if( charVal.toUpperCase() != charVal.toLowerCase() )
    return true;
  else
    return false;
} 
 