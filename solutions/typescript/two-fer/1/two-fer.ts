/**
 * This stub is provided to make it straightforward to get started.
 */

export function twoFer(Name?: string): string {
  // ^                 ^   ^ this is called a return type; it's the type of the
  // ^                 ^     value that is returned from this function
  // ^                 ^
  // ^                 parameters go here
  // ^
  // allows the tests to import this function and call it
  // <-- Your code goes here. You may remove all the commentary in this file.
  if (typeof Name === 'undefined') {
    Name = "you"
  }
  return "One for "+Name+", one for me."
}
