/**
 * @param {number[][]} mat
 * @param {number} r
 * @param {number} c
 * @return {number[][]}
 */
const matrixReshape = function(mat, r, c) { 
  const matrixRows = mat.length; 
  const matrixCols = mat[0].length; 

  if (matrixRows * matrixCols !== r * c) {
    return mat; 
  } 

  let col = 0; 
  let row = 0; 

  const resultArray = (new Array(r)).fill().map(() => Array(c).fill(0))

  for (let i = 0; i < mat.length; i++) {
    for (let j = 0; j < mat[0].length; j++) {
      if (col === c) {
        row++; 
        col = 0; 
      }
      resultArray[row][col] = mat[i][j]; 
      col++; 
    }
  }

  return resultArray; 

};



console.log(matrixReshape([[1, 2], [3, 4]], 1, 4)); // [[1, 2, 3, 4]]
console.log(matrixReshape([[1, 2], [3, 4]], 1, 3)); // [[1, 2], [3, 4]]

console.log(matrixReshape([[1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12]], 4, 3)); // [[1, 2, 3], [4, 5, 6], [7, 8, 9], [10, 11, 12]]