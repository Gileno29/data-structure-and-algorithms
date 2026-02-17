# First Unique Character using Hash Maps (Go)

This repository features an efficient solution for finding the first non-repeating character in a string using a **Hash Map** (the `map` type in Go).

## 🗺️ What is a Hash Map?

A **Hash Map** is a data structure that implements an associative array abstract data type, a structure that can map keys to values. It uses a **hash function** to compute an index into an array of buckets or slots, from which the desired value can be found.



## 🚀 How it Works in this Implementation

1.  **Character Tracking**: We convert the string into a slice of `runes` to handle Unicode safely.
2.  **The Map Structure**: We use a map `map[string][]int` where:
    * **Key**: The character (as a string).
    * **Value**: A slice `[]int{index, frequency}`.
3.  **Frequency Count**: We loop through the string. If the character is new, we store its index and set frequency to 1. If it exists, we increment the frequency.
4.  **Order Preservation**: Since Go maps are unordered, we iterate through the **original string** a second time. We check each character's frequency in the map and return the index of the first one with a frequency of 1.

---

## 📊 Execution Trace (Step-by-Step)

**Input:** `"leetcode"`  
**Target:** Find the index of the first unique letter.

| Step | Char | Action | Map State `{"char": [index, count]}` |
| :--- | :--- | :--- | :--- |
| 1 | 'l' | New | `{"l": [0, 1]}` |
| 2 | 'e' | New | `{"l": [0, 1], "e": [1, 1]}` |
| 3 | 'e' | Exists | `{"l": [0, 1], "e": [1, 2]}` |
| 4 | 't' | New | `{"l": [0, 1], "e": [1, 2], "t": [3, 1]}` |
| 5 | 'c' | New | `{"l": [0, 1], "e": [1, 2], "t": [3, 1], "c": [4, 1]}` |
| ... | ... | ... | ... |

**Second Pass (Iteration over "leetcode"):**
1. Is `d["l"][1] == 1`? **Yes!** -> **Return Index 0**.

---

## ⚖️ Complexity Analysis

### Time Complexity: $O(n)$
* **First Pass:** We iterate through the string of length $n$ to populate the map ($O(n)$).
* **Second Pass:** We iterate through the string again to find the first unique character ($O(n)$).
* **Total:** $O(n) + O(n) = O(2n)$, which simplifies to **$O(n)$**.

### Space Complexity: $O(k)$
* We store counts in a map. In the worst case, $k$ is the number of unique characters in the alphabet. For standard ASCII, $k \leq 256$. Since $k$ is constant relative to the input size, it is **$O(1)$** in terms of alphabet size, or **$O(k)$** for Unicode.

---

## 🛠️ Common Problems Solved with Hash Maps

Hash maps are the "Swiss Army Knife" of data structures. Common use cases include:

* **Frequency Counting:** Tracking occurrences of words or characters (as seen here).
* **Two Sum Problem:** Finding two numbers in an array that add up to a specific target.
* **Caching/Memoization:** Storing results of expensive function calls to avoid re-calculation.
* **Grouping/Anagrams:** Grouping strings that contain the same characters.
* **Data Indexing:** Building an index to quickly retrieve records by a specific field (e.g., ID).



---

## 💻 Code Snippet

```go
func FirstUniqChar(s string) int {
    characters := []rune(s)
    d := make(map[string][]int) // map[char] -> [first_index, count]

    for index, value := range characters {
        charStr := string(value)
        if _, ok := d[charStr]; !ok {
            d[charStr] = []int{index, 1}
        } else {
            d[charStr][1]++
        }
    }

    for _, value := range characters {
        charStr := string(value)
        if d[charStr][1] == 1 {
            return d[charStr][0]
        }
    }

    return -1
}


## ⚡ Space Optimization: Array vs. Hash Map

While a `map` is flexible for any Unicode character, it carries significant memory overhead. If we know the input only contains **lowercase English letters**, we can use an array of size 26.

### Why use an array?
1.  **Memory**: A map stores metadata, pointers, and hash buckets. An array is a contiguous block of memory.
2.  **Speed**: No hashing is required. We calculate the index using a simple subtraction: `char - 'a'`.
3.  **Predictability**: The space is always exactly 26 integers, regardless of the string length.



### Optimized Implementation

```go
func FirstUniqChar(s string) int {
    // Array to store the count of each character (a-z)
    var counts [26]int

    // First pass: Count frequencies
    for _, char := range s {
        counts[char-'a']++
    }

    // Second pass: Find the first character with a count of 1
    for i, char := range s {
        if counts[char-'a'] == 1 {
            return i
        }
    }

    return -1
}