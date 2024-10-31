/*
Write a function called `same`, which accepts two arrays. 
The function should return true if every value in the array has it corresponding value squared in the second array
The frequency of values must be the same.

The order does not matter!
*/

/*
INPUT
array_1: number[]
array_2: number[]

OUTPUT
returns --> boolean  

RULES (Explicit, Implicit, Assumptions, Boundaries, Edge Cases, Terms, Formulas)
- @R1: returns true if every value in `array_1` has a corresponding value 'squared' in `array_2`
- @R2: The frequency of values must be the same 

- (Note): The order of values does not matter
- (formula): squared --> Math.pow(n, 2) or n ** 2

EXLORATION and TESTS
T1 --> satiesfies @R1 and @R2
array_1: [1, 2, 3, 4]
array_2: [4, 9, 1, 16]
result: true 

T2 --> Does not satiesfy @R1 and @R2
array_1: [1, 2, 3, 4]
array_2: [4, 9, 1, 16, 16]
result: true 


MENTAL MODEL
{
  1: 1, 
  2: 1, 
  3: 1,
  4: 2, ---> 4 ** 2 = 16, if array_2[16] also has a count of 2, continue. Otherwise, return false
}

{
  4: 1,
  9: 1,
  1: 1
  16: 2,
}


DATA STRUCTURE 
- Object (count frequencies)

ALGORITHMS
- if the size of array_1 and array_2 are not equal
  - return false 

- set `array_1_frequencies` to (FUNCTION: frequencies in each array)
- set `array_2_frequencies` to (FUNCTION: frequencies in each array)

- for each value in array_1_frequencies 
  - set current_value_squared to (value ** 2) 
  - if array_1_frecuencies[current_value_squared] does not have a matching value in array_2_frequencies
    - return false 
- return true 

COMPLEXITY 
- N: the number of items in array 1
- M: the number of items in array 2
- time: O(N + M)
- space: O(max(N, M))

REFACTOR 

*/

const getFrequencies = (items: number[]) => {
  let result: { [key: number]: number } = {};
  for (let index = 0; index < items.length; index += 1) {
    const currentValue = items[index];
    if (!(currentValue in result)) result[currentValue] = 0;
    result[currentValue] += 1;
  }

  return result;
};

const same = (array_1: number[], array_2: number[]): Boolean => {
  if (array_1.length !== array_2.length) {
    return false;
  }

  const array1Freq = getFrequencies(array_1);
  const array2Freq = getFrequencies(array_2);

  for (let array1Key in array1Freq) {
    const value1Squared = Number(array1Key) ** 2;

    // has matching squared values
    if (!(value1Squared in array2Freq)) {
      return false;
    }

    // has the same frequencies
    if (array1Freq[array1Key] !== array2Freq[Number(array1Key) ** 2]) {
      return false;
    }
  }

  return true;
};

const array_01_t1 = [1, 2, 3, 4];
const array_02_t1 = [4, 9, 1, 16];
console.log(same(array_01_t1, array_02_t1)); // true

const array_01_t2 = [1, 2, 3, 4];
const array_02_t2 = [4, 9, 1, 16, 16];
console.log(same(array_01_t2, array_02_t2)); // false

const array_01_t3 = [1, 2, 3, 4];
const array_02_t3 = [4, 9, 1, 4];
console.log(same(array_01_t3, array_02_t3)); // false
