# Sliding Window Algorithm in Go

This repository contains a practical implementation of the **Sliding Window** technique to solve a substring/subarray frequency problem.

## 🪟 What is a Sliding Window?

The **Sliding Window** is an algorithmic pattern used to perform operations on a specific subset of data (a "window") within a larger data structure (like an array or string). 

Instead of using nested loops to recalculate every possible subarray (which usually takes $O(n^2)$ time), we maintain two pointers—**Left ($l$)** and **Right ($r$)**—that "slide" across the data. This allows us to process the entire array in a single pass ($O(n)$ time complexity).



## 🚀 How it Works in this Implementation

The goal of this specific code is to find the **maximum length of a subarray** where each element appears **at most twice**.

1.  **Expand**: We move the `right` pointer to include a new element in the window.
2.  **Count**: We track the frequency of each element using a `map`.
3.  **Contract**: If an element's frequency exceeds 2 (the "invalid" state), we move the `left` pointer forward, removing elements until the window becomes valid again.
4.  **Update**: At each valid step, we calculate the window size ($r - l + 1$) and update our `maxLen`.

## 🛠️ Problem Types Solved by Sliding Window

* **Fixed Window:** Finding the average of all subarrays of size `K`.
* **Dynamic Window:** Finding the smallest subarray with a sum greater than `X`.
* **String/Pattern:** Finding the longest substring with unique characters or $K$ distinct characters.
* **Frequency:** Our current problem (longest subarray with at most $K$ repetitions).

---

## 📊 Execution Trace (Step-by-Step)

**Input:** `["b", "c", "b", "b", "b", "c", "b", "a"]`  
**Goal:** Max length where frequency $\le 2$.

| Step | Pointer `r` | Char | Action | Map State | `l` | Window | `maxLen` |
| :--- | :--- | :--- | :--- | :--- | :--- | :--- | :--- |
| 1 | 0 | "b" | Add | `{"b":1}` | 0 | `["b"]` | 1 |
| 2 | 1 | "c" | Add | `{"b":1, "c":1}` | 0 | `["b", "c"]` | 2 |
| 3 | 2 | "b" | Add | `{"b":2, "c":1}` | 0 | `["b", "c", "b"]` | 3 |
| 4 | 3 | "b" | **Invalid (>2)** | `{"b":3, "c":1}` | 0 | `["b", "c", "b", "b"]` | 3 |
| 5 | 3 | - | Shrink (L++) | `{"b":2, "c":1}` | 1 | `["c", "b", "b"]` | 3 |
| 6 | 4 | "b" | **Invalid (>2)** | `{"b":3, "c":1}` | 1 | `["c", "b", "b", "b"]` | 3 |
| 7 | 4 | - | Shrink (L++) | `{"b":3, "c":0}` | 2 | `["b", "b", "b"]` | 3 |
| 8 | 4 | - | Shrink (L++) | `{"b":2, "c":0}` | 3 | `["b", "b"]` | 3 |
| 9 | 5 | "c" | Add | `{"b":2, "c":1}` | 3 | `["b", "b", "c"]` | 3 |
| 10 | 6 | "b" | **Invalid (>2)** | `{"b":3, "c":1}` | 3 | `["b", "b", "c", "b"]` | 3 |
| 11 | 6 | - | Shrink (L++) | `{"b":2, "c":1}` | 4 | `["b", "c", "b"]` | 3 |
| 12 | 7 | "a" | Add | `{"b":2, "c":1, "a":1}` | 4 | `["b", "c", "b", "a"]` | **4** |

**Final Result: 4**

---

## 💻 Code Snippet

```go
func Sliding(s []string) int {
    l, r, maxLen := 0, 0, 0
    counter := make(map[string]int)

    for r < len(s) {
        charR := s[r]
        counter[charR]++

        for counter[charR] > 2 {
            charL := s[l]
            counter[charL]--
            l++
        }

        if (r-l)+1 > maxLen {
            maxLen = (r - l) + 1
        }
        r++
    }
    return maxLen
}


## ⚖️ Complexity Analysis

To understand why this approach is efficient, let's look at the "Big O" notation:

* **Time Complexity:** $O(n)$
    * Even though there is a `for` loop inside another `for` loop, it is still linear. This is because each element is visited at most twice: once by the `right` pointer (to add it) and once by the `left` pointer (to remove it).
* **Space Complexity:** $O(k)$
    * Where $k$ is the number of unique elements in the input. We store these in a map to track their frequencies. In the worst case (all elements are unique), it is $O(n)$.



## ⚠️ Edge Cases Handled

The current logic is robust enough to handle the following scenarios:

1.  **Empty Array:** Returns `0` immediately.
2.  **All Same Elements:** If the input is `["a", "a", "a", "a"]`, the window will correctly shrink to size `2`.
3.  **No Repetitions:** If every element is unique, the window never shrinks, and it returns the total length of the array.
4.  **Single Element:** Correcty returns `1`.

---

## 🛠️ How to Customize

If you want to change the allowed number of repetitions, simply change the constant in the inner loop condition:

```go
// Change '2' to 'K' to allow K repetitions
for counter[charR] > K {
    counter[s[l]]--
    l++
}