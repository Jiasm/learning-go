function main (coins, amount) {
 const dp = new Array(amount + 1).fill(Infinity)

 dp[0] = 0

 for (let i = 1; i <= amount; i++) {
   for (const coin of coins) {
     if (i - coin >= 0) {
      dp[i] = Math.min(dp[i], dp[i - coin] + 1)
    }
   }
 }

 return dp[amount] === Infinity ? -1 : dp[amount]
}

console.log(main([1, 2, 5], 49))

function squrd (matrix) {
  let max = 0
  for (let i = 0, len = matrix.length; i < len; i++) {
    for (let j = 0, rowLen = matrix[i].length; j < rowLen; j++) {
      if (matrix[i][j] === 1 && i > 0 && j > 0) {
        matrix[i][j] = Math.min(matrix[i - 1][j], matrix[i - 1][j - 1], matrix[i][j - 1]) + 1
        max = Math.max(matrix[i][j], max)
      }
    }
  }

  return max ** 2
}

console.log(squrd([
  [1, 0, 1, 1, 1],
  [1, 0, 1, 0, 1],
  [1, 1, 0, 0, 1],
  [1, 0, 0, 1, 1],
]))

function thief (buildings) {
  const len = buildings.length
  const dp = new Array(len)

  dp[0] = buildings[0]
  dp[1] = Math.max(buildings[0], buildings[1])
  dp[2] = Math.max(buildings[0] + buildings[2], buildings[1])

  for (let i = 3; i < len; i++) {
    dp[i] = Math.max(buildings[i] + dp[i - 2], dp[i - 1])
  }

  console.log(dp)
  return Math.max(dp[len - 1], dp[len - 2])
}

console.log(thief([20, 20, 20, 50, 70]))

function subsets (nums) {
  let res = []
  let n = nums.length

  function back (path, i) {
    if (i <= n) {
      console.log('insert', path)
      res.push(path)
    }

    for (let j = i; j < n; j++) {
      path.push(nums[j])
      back([...path], j + 1)
      path.pop()
    }
  }

  back([], 0)

  return res
}

console.log(subsets([1, 2, 3]))

function maxSub (str) {
  let len = str.length
  if (len <= 1) return 1
  let max = 1

  for (let start = 0; start < len; start++) {
    const res = [str[start]]
    for (let end = start + 1; end < len; end++) {
      if (res.indexOf(str[end]) === -1) {
        res.push(str[end])
        max = Math.max(max, end - start + 1)
      } else {
        break
      }
    }
  }

  return max
}

console.log(maxSub('aabcdbd'))

function a () {
  for (let i = 0; i < 5; i++) {
    this.i = i
    setTimeout(function () {
      console.log(i)
    }, 0)
    console.log(this.i)
  }
}

console.log('=======')
a()

function wordPattern (pattern, str) {
  const patternMap = new Map()
  for (let index = 0; index < str.length; index++) {
    const p = pattern[index]
    const word = str[index]
    // init
    if (!patternMap.has(p)) {
      patternMap.set(p, word)
    } else {
      if (patternMap.get(p) !== word) {
        return false
      }
    }
  }

  return true
}

console.log(wordPattern('abba', 'dog cat cat dog'.split(' ')))
console.log(wordPattern('abba', 'dog cat cat cat'.split(' ')))

function moveZero (nums) {
  let i = 0
  let j = 0
  const len = nums.length
  for (; i < len; i++) {
    if (nums[i] !== 0) {
      [nums[i], nums[j]] = [nums[j], nums[i]]
      j++
    }
  }
  return nums
}

console.log(moveZero([0, 15, 0, 3, 12]))

function nimGame (count) {
  return count % 4 !== 0
}

console.log(nimGame(5))
console.log(nimGame(4))
console.log(nimGame(3))
console.log(nimGame(6))

function sortStack (stack, index = 0) {
  if (index === 0) {
    [stack[index], stack[stack.length - 1]] = [stack[stack.length - 1], stack[index]]
  }
  const current = stack[index]
  if (!current) return
  const left = index * 2 + 1
  const right = index * 2 + 2

  let leftVal = stack[left]
  let rightVal = stack[right]

  if (!leftVal && !rightVal) return

  if (!rightVal || leftVal > rightVal) {
    if (current < leftVal) {
      [stack[index], stack[left]] = [leftVal, current]
      // 重排后需要再次检测子集
      sortStack(stack, left)
    }
  } else if (!leftVal || rightVal > leftVal) {
    if (current < rightVal) {
      [stack[index], stack[right]] = [rightVal, current]
      // 重排后需要再次检测子集
      sortStack(stack, right)
    }
  }
}

function adjust (arr, index) {
  const left = index * 2 + 1
  const right = index * 2 + 2
  const current = arr[index]

  let leftVal = stack[left]
  let rightVal = stack[right]

  if (!leftVal && !rightVal) return

  if (!rightVal || leftVal > rightVal) {
    if (current < leftVal) {
      // 交换子父节点
      [stack[index], stack[left]] = [leftVal, current]
      // 继续配平下级
      adjust(stack, left)
    }
  } else if (!leftVal || leftVal > rightVal) {
    if (current < rightVal) {
      // 交换子父节点
      [stack[index], stack[right]] = [rightVal, current]
      // 继续配平下级
      adjust(stack, right)
    }
  }
}

function buildHeap (arr) {
  const heap = arr
  for (let i = heap.length - 1; i >= 0; i--) {
    adjust(heap, i)
  }

  return heap
}

const stack = [1,9,8,7,2,6,5,4,3]
const heap = buildHeap(stack)
console.log({heap})

function heapSort (heap) {
  const sorted = []

  while (heap.length) {
    const firstNode = heap.shift()
    sorted.push(firstNode)
    const lastNode = heap.pop()
    if (!lastNode) break
    heap.unshift(lastNode)
    adjust(heap)
  }

  return sorted
}
console.log(heapSort(heap))


// function deSortStack (stack, index = 0) {
//   if (index === 0) {
//     [stack[index], stack[stack.length - 1]] = [stack[stack.length - 1], stack[index]]
//   }
//   const current = stack[index]
//   if (!current) return
//   const left = index * 2 + 1
//   const right = index * 2 + 2

//   let leftVal = stack[left]
//   let rightVal = stack[right]

//   if (!leftVal && !rightVal) return

//   if (!rightVal || leftVal < rightVal) {
//     if (current > leftVal) {
//       [stack[index], stack[left]] = [leftVal, current]
//       // 重排后需要再次检测子集
//       sortStack(stack, left)
//     }
//   } else if (!leftVal || rightVal < leftVal) {
//     if (current > rightVal) {
//       [stack[index], stack[right]] = [rightVal, current]
//       // 重排后需要再次检测子集
//       sortStack(stack, right)
//     }
//   }
// }

// const stack2 = [2, 9,8,7,6,5,4,3,1]
// deSortStack(stack2)
// console.log({stack2})
