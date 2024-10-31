/*
Given two strings, write a function to determine if the second string is an anagram of the first. An anagram is a word, phrase, or name formed by rearranging the letters of another, sucha as *cinema*, formed from *iceman*. 
*/

/*
INPUT
- sourceSequence: string
- targetSequence: string

OUTPUT
- result: boolean 

RULES 
- @Explicit_Rule_1: determine if the second string (target) is an anagram of the first string (source)
- (Anagram): a word, phrase or name formed by rearranging letter of another 
    - e.g.: "cinema" can be formed from "iceman"
- @Implicit_Rule_1: both strings must have the same length
- @Implicit_Rule_2: empty strings are considered anagram of each other
- @Implicit_Rule_3: both strings must have the same frequency of characters
- @Assumption_1: valid arguments will always be provided. 
    - We'll never be given an unexpected value such as a number, null or undefined


EXPLORATION / TESTS
T1
ss: "abcde"
ts: "edcba"
@returns true 

T2
ss: "abcde"
ts: ""
@returns false

NAIVE SOLUTION 
// time: O(N+M)
// space: O(N+M) 

DATA STRUCTURE
- object (count character)


ALGORITHM 
- if the lengths of `sourceSquence` and `targetSequence` are not equal 
    - return false 
- if both `sourceSequence` and `targetSequence` are empty 
    - return true 
- set `sourceCounter` = (FUNC: count frequencies)
- for each letter in targetSquence 
    - if sourceCounter[letter] is truthy
        - sourceCounter[letter] -= 1 
- if value in each `sourceCounter` property is zero 
    - return true 
- Otherwise
    - return false

COMPLEXITY 
- time: O(N)
- space: O(N)
*/

// const validAnagram = (sourceSequence: string, targetSequence: string) => {
//   const sortedSourceSequence = sourceSequence.split("").sort().join("");
//   const sortedTargetSequence = sourceSequence.split("").sort().join("");
//   return sortedSourceSequence === sortedTargetSequence;
// };

const validAnagram = (sourceSequence: string, targetSequence: string) => {
  if (sourceSequence.length !== targetSequence.length) {
    return false;
  }

  if (sourceSequence.length === 0 && targetSequence.length === 0) {
    return true;
  }

  const sourceCounter = sourceSequence
    .split("")
    .reduce((acc: Record<string, number>, curr: string) => {
      // return acc[curr] = acc[curr] ? ++acc[curr] : 1, acc;
      acc[curr] = (acc[curr] || 0) + 1;
      return acc;
    }, {});

  for (let letter of targetSequence) {
    if (!sourceCounter[letter]) return false;
    sourceCounter[letter] -= 1;
  }

  for (let letter of targetSequence) {
    if (sourceCounter[letter] !== 0) {
      return false;
    }
  }

  return true;
};

console.log(validAnagram("", "")); // true
console.log(validAnagram("qwerty", "qeywrt")); // true
console.log(validAnagram("anagram", "nagaram")); // true
console.log(validAnagram("texttwisttime", "timetwisttext")); // true

console.log(validAnagram("aaz", "zza")); // false
console.log(validAnagram("rat", "car")); // false
console.log(validAnagram("awesome", "awesom")); // false
