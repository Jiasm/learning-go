function compare (arr) {
  if (arr.length <= 1) {
    return arr
  }
  const mid = Math.floor(arr.length / 2) 
  const left = compare(arr.slice(0, mid))
  const right = compare(arr.slice(mid))

  const result = []

  for (let leftIndex = 0, rightIndex = 0; leftIndex < left.length || rightIndex < right.length; ) {
    if (leftIndex >= left.length) {
      result.push(right[rightIndex++])
    } else if (rightIndex >= right.length) {
      result.push(left[leftIndex++])
    } else {
      if (left[leftIndex] < right[rightIndex]) {
        result.push(left[leftIndex++])
      } else {
        result.push(right[rightIndex++])
      }
    }
  }

  return result
}

console.log(compare([1, 8, 5, 3, 4, 7, 2]))

function cuttingRope (number) {
  if (number === 2 || number === 3) return number - 1

  const a = number % 3
  const b = parseInt(number / 3)

  if (a === 0) {
    return 3 ** b
  } else if (a === 1) {
    return 2 * 2 * (3 ** (b - 1))
  } else {
    return 2 * (3 ** b)
  }
}

console.log(cuttingRope(11))