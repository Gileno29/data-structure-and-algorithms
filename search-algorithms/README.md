# Binary Search and Variations

Binary search is an efficient search algorithm that uses the "divide and conquer" strategy.

You just can use a binary search in cases where the list or array is alrredy ordenated.

## How it works (Step by step)

1.  **Find the Middle**: The algorithm looks at the element in the center of the list.
2.  **Compare**:
    *   If the central element is the target, the search ends.
    *   If the target is smaller, the right half is discarded.
    *   If the target is larger, the left half is discarded.
3.  **Repeat**: Until the value is found or the list is empty.

## Interview Challenge: Search in Rotated Array

This algorithm is a "superset" of the traditional binary search. If the array is perfectly sorted (e.g., `[0, 1, 2, 4, 5, 6, 7]`), the condition `nums[low] <= nums[mid]` will always be true for the first half, and the code will always fall into the standard search logic. It is universal for sorted arrays.

### Dry Run Example

**Array**: `[4, 5, 6, 7, 0, 1, 2]` | **Target**: `0`

| Iteration | low | high | mid | nums[mid] | Applied Logic |
| :--- | :-: | :--: | :-: | :---: | :--- |
| 1 | 0 | 6 | 3 | 7 | Left side is sorted (4 <= 7). Target 0 is not between 4 and 7. Move `low` to `mid + 1`. |
| 2 | 4 | 6 | 5 | 1 | Right side is sorted (1 <= 2). Target 0 is not between 1 and 2. Move `high` to `mid - 1`. |
| 3 | 4 | 4 | 4 | 0 | `nums[mid] == target`. Return index 4. |

## Go Implementation

```go
package main

import "fmt"

func SearchRotated(nums []int, target int) int {
    low := 0
    high := len(nums) - 1

    for low <= high {
        mid := low + (high-low)/2

        if nums[mid] == target {
            return mid
        }

        // If the array were not rotated, this IF would always be true.
        if nums[low] <= nums[mid] { // Left side is sorted
            if target >= nums[low] && target < nums[mid] {
                high = mid - 1
            } else {
                low = mid + 1
            }
        } else { // Right side is sorted
            if target > nums[mid] && target <= nums[high] {
                low = mid + 1
            } else {
                high = mid - 1
            }
        }
    }
    return -1
}

func main() {
    rotatedArray := []int{4, 5, 6, 7, 0, 1, 2}
    normalArray := []int{0, 1, 2, 4, 5, 6, 7}
    target := 0

    fmt.Printf("Rotated - Index of %d: %d\n", target, SearchRotated(rotatedArray, target))
    fmt.Printf("Normal  - Index of %d: %d\n", target, SearchRotated(normalArray, target))
}
```

## Key Points for the Interview

1.  **Universality**: The code handles both cases (rotated or not).
2.  **Integer Division**: In Go, `5 / 2` results in `2`, which correctly handles calculating the middle index in lists with an odd or even number of elements.


## Complexty

The time complexy of binary search is O(log n) and the espace complexy of binary search is O(1)