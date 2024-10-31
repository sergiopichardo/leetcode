/*
Write a function called `same`, which accepts two arrays. 
The function should return true if every value in the array has it corresponding value squared in the second array
The frequency of values must be the same.
*/

/*
INPUT
array_1: number[]
array_2: number[]

OUTPUT
return_value: true | false (boolean) 

RULES 
- returns true if: 
    - every value in the array has its corresponding value squared in the second array 
- the frequency of values must be the same

Assumptions
- the values in array 1 have a corresponding squared value in array 2
- arrays will be of at least size 1

EXAMPLES 

T1
a1: [2, 2, 2]
a2: [4, 4, 4]
rv: true 

T2
a1: [1, 2, 3, 4]
a2: [1, 4, 9, 16, 16]
rv: false - because the arrays don't have the same frequency of values

T3
a1: [1, 2, 3, 9]
a2: [1, 4, 9, 81]
rv: true 

T4
a1: [0]
a2: [2]
rv: true

T5
a1: [1]
a2: [1]
rv: true

T6
a1: [5]
a2: [10]
rv: false

MENTAL MODEL 
- we need to check if both arrays have the same length
- We need to go over each value in array one, square it, and check if the value at the same position in array two is squared
    - if all the values are squared, then return true 
    - if a value is not squared then return false 

DATA STRUCTURE 
- array (iterate over array_1)

ALGORITHM 
- if array_1 and array_2 DO NOT have the same length
    - return false 
- iterate over array_1 
    - if the current value (squared) is identical to the value at the same position in array_2 ARE NOT the same
        - return false 
- return true 
*/

const same_v1 = (array_1: number[], array_2: number[]): Boolean => {
  if (array_1.length !== array_2.length) {
    return false;
  }

  for (let index = 0; index < array_1.length; index += 1) {
    const currentValue = array_1[index];
    const valueSquared = currentValue ** 2;
    if (valueSquared !== array_2[index]) {
      return false;
    }
  }

  return true;
};

const array_1_t1 = [2, 2, 2];
const array_2_t1 = [4, 4, 4];
console.log(same_v1(array_1_t1, array_2_t1)); // true

const array_1_t2 = [1, 2, 3, 4];
const array_2_t2 = [1, 4, 9, 16, 16];
console.log(same_v1(array_1_t2, array_2_t2)); // false

const array_1_t3 = [1, 2, 3, 9];
const array_2_t3 = [1, 4, 9, 81];
console.log(same_v1(array_1_t3, array_2_t3)); // true

const array_1_t4 = [0];
const array_2_t4 = [0];
console.log(same_v1(array_1_t4, array_2_t4)); // true

const array_1_t5 = [1];
const array_2_t5 = [1];
console.log(same_v1(array_1_t5, array_2_t5)); // true

const array_1_t6 = [1, 5, 10];
const array_2_t6 = [1, 25, 20];
console.log(same_v1(array_1_t6, array_2_t6)); // false
