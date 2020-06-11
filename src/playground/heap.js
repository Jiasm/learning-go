function buildHeap (arr) {
  const heap = []
  for (let i = 0, len = arr.length; i < len; i++) {
    const node = arr[i]
    heap.push(node)
    upAdjust(heap, i)
  }

  return heap
}

function upAdjust (heap, index) {
  if (index === 0) return
  const parentIndex = Math.ceil(index / 2) - 1

  if (heap[index] > heap[parentIndex]) {
    [heap[index], heap[parentIndex]] = [heap[parentIndex], heap[index]]
    upAdjust(heap, parentIndex)
  }
}

function sortHeap (heap) {
  const sorted = []
  for (let i = 0, len = heap.length; i < len; i++) {
    const node = heap.shift()
    sorted.push(node)
    const last = heap.pop()
    heap.unshift(last)
    downAdjust(heap, 0)
  }

  return sorted
}

function downAdjust (heap, index) {
  const leftIndex = index * 2 + 1
  const rightIndex = index * 2 + 2

  const current = heap[index]

  if (!current) return

  const left = heap[leftIndex]
  const right = heap[rightIndex]

  if (!left && !right) return

  if (!right || left > right) {
    if (current < left) {
      [heap[index], heap[leftIndex]] = [heap[leftIndex], heap[index]]
      // 检测子节点是否需要重新配平
      downAdjust(heap, left)
    }
  } else if (!left || right > left) {
    if (current < right) {
      [heap[index], heap[rightIndex]] = [heap[rightIndex], heap[index]]
      // 检测子节点是否需要重新配平
      downAdjust(heap, right)
    }
  }
}

const heap = buildHeap([1, 7, 5, 2, 9, 8, 3, 4, 6])
console.log(heap)

console.log(sortHeap(heap))

function quickSort (arr, start, end) {
  console.log({start, end})
  if (start < end) {
    const index = getPartation(arr, start, end)

    quickSort(arr, start, index - 1)
    quickSort(arr, index + 1, end)

    return arr
  }
}

function getPartation (arr, start, end) {
  const tmp = arr[start]

  let low = start
  let high = end

  while (low < high) {

    while (low < high && arr[high] >= tmp) {
      high--
    }

    arr[low] = arr[high]

    while (low < high && arr[low] <= tmp) {
      low++
    }
  
    arr[high] = arr[low]
  }
  arr[low] = tmp

  return low
}

console.log(quickSort([1, 7, 2, 5, 9, 4], 0, 5))