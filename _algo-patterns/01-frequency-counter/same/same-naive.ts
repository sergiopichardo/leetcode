const sameNaive = (array_1: number[], array_2: number[]) => {
  if (array_1.length !== array_2.length) {
    return false;
  }

  for (let index = 0; index < array_1.length; index += 1) {
    const currentValue = array_1[index];
    const squaredValue = Math.pow(currentValue, 2);
    let array2TargetIndex = array_2.indexOf(squaredValue);

    if (array2TargetIndex === -1) {
      return false;
    }

    array_2.splice(array2TargetIndex, 1);
  }

  return true;
};

console.log(sameNaive([1, 2, 3, 4], [4, 9, 1, 16])); // true
console.log(sameNaive([1, 2, 3, 4], [4, 9, 1, 16, 16])); // false
console.log(sameNaive([1, 2, 3, 4], [4, 9, 1, 4])); // false
