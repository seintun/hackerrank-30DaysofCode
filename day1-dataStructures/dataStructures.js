function main() {
  var i = 4;
  var d = 4.0;
  var s = "HackerRank ";
  // Declare second integer, double, and String variables.
  let integer = parseInt(readLine());
  let double = parseFloat(readLine());
  let string = readLine();
  // Read and save an integer, double, and String to your variables.
  // Print the sum of both integer variables on a new line.
  console.log(i + integer);
  // Print the sum of the double variables on a new line.
  console.log((d + double).toFixed(1));
  // Concatenate and print the String variables on a new line
  // The 's' variable above should be printed first.
  console.log(s + string);
}
