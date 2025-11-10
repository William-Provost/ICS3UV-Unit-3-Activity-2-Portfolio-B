/**
 * @author William Provost
 * @version 1.0.0
 * @date 2025-11-10
 * @fileoverview This program shows formatting of numbers.
 */

// variables
const PI: number = 3.141593;      // constant variable for Pi
const age: number = 8.3;          // age is set to 8 years old
const studentName: string = "Johnny"; // String object set to Johnny

// output
console.log(`${studentName} is ${Math.round(age)} years old.`);
console.log(`The value of PI is ${PI.toFixed(2)}`);
console.log(`PI = ${PI.toFixed(6)}`);
console.log(`${PI.toFixed(3)}`);

console.log("\nDone.");
