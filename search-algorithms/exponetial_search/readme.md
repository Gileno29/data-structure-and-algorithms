# Exponential Search Algorithm in Go

This repository contains an implementation of the **Exponential Search** algorithm, a highly efficient technique for searching sorted arrays, especially those of infinite size or when the target element is near the beginning.

## 📈 What is Exponential Search?

Exponential Search works in two distinct phases:
1.  **Range Finding**: It rapidly finds a range where the target might exist by doubling the index (`1, 2, 4, 8, 16...`) until it finds an element greater than or equal to the target.
2.  **Binary Search**: Once the range is narrowed down, it performs a standard Binary Search within that specific block.



## 🚀 How it Works in this Implementation

1.  **Initial Check**: It first checks if the first element is the target to avoid unnecessary loops.
2.  **The "Jump" Phase**: It starts at index 1 and multiplies the index by 2 (`right *= 2`) as long as the value at that index is less than the target.
3.  **The "Clamp" Phase**: Using `min(right, n-1)`, we ensure that our search range never exceeds the actual boundaries of the array (avoiding "Out of Bounds" errors).
4.  **Binary Search**: It calls `BinarySearch` on the slice between the previous power of 2 (`right/2`) and the current `right`.

---

## 📊 Execution Trace (Step-by-Step)

**Input Array:** `[1, 3, 5, 7, 9, 11, 13, 15, 17, 19]`  
**Target:** `15`  
**Array Length ($n$):** `10`

| Step | Action | `right` Index | `arr[right]` | Logic |
| :--- | :--- | :--- | :--- | :--- |
| 1 | Start | 1 | 3 | $3 < 15$, so double `right` |
| 2 | Jump | 2 | 5 | $5 < 15$, so double `right` |
| 3 | Jump | 4 | 9 | $9 < 15$, so double `right` |
| 4 | Jump | 8 | 17 | $17 > 15$, Stop Jumping! |
| 5 | Range | **[4, 8]** | - | Binary Search between indices 4 and 8 |

**Final Phase: Binary Search on `[9, 11, 13, 15, 17]`** * Midpoint finds `13` $\rightarrow$ search right.
* Next midpoint finds `15` $\rightarrow$ **Return Index 7**.

---

## ⚖️ Complexity Analysis

### Time Complexity
* **Best Case:** $O(1)$ if the target is the first element.
* **Average/Worst Case:** $O(\log i)$, where $i$ is the index of the target. 
    * *Note:* This is often better than standard Binary Search ($O(\log n)$) if the target is located in the early part of the array.

### Space Complexity
* **$O(1)$**: This implementation is iterative (even the binary search part), meaning it uses a constant amount of extra memory regardless of input size.



## 🛠️ Key Advantage: The "Infinite" Search
Exponential Search is particularly useful for **unbounded or infinite arrays** (like a live stream of data or a very large file) where you don't know the total length ($n$) in advance. You can find the upper bound without ever knowing how large the array actually is.

---

## 💻 Code Snippet

```go
func ExponentialSearch(arr []int, target int) int {
    if arr[0] == target {
        return 0
    }

    n := len(arr)
    right := 1
    for right < n && arr[right] < target {
        right *= 2
    }

    // Binary search within the discovered range
    // min(right, n-1) prevents index out of bounds
    return BinarySearch(arr, target, right/2, min(right, n-1))
}