export function isLeap(year: number) {
  if (year % 4 == 0) {
    if (year % 100 == 0 && year % 400 != 0) {
      return false
    } else {
      return true
    }
  } else {
    return false
  }
}
