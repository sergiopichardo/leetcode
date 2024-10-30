/*
Write a function called `same`, which accepts two arrays. 
The function should return true if every value in the array has it corresponding value squared in the second array
The frequency of values must be the same.

The order does not matter!
*/

const same = (array_1: number[], array_2: number[]): Boolean => {
  return true;
};

const array_1_t1 = [1, 2, 3];
const array_2_t1 = [9, 1, 4];
console.log(same(array_1_t1, array_2_t1)); // true

const array_1_t2 = [1, 2, 3, 4];
const array_2_t2 = [4, 1, 16, 16, 9];
console.log(same(array_1_t2, array_2_t2)); // false

const array_1_t3 = [1, 2, 3, 9];
const array_2_t3 = [1, 81, 4, 9];
console.log(same(array_1_t3, array_2_t3)); // true

const array_1_t4 = [0];
const array_2_t4 = [0];
console.log(same(array_1_t4, array_2_t4)); // true

const array_1_t6 = [];
const array_2_t6 = [];
console.log(same(array_1_t6, array_2_t6)); // false
