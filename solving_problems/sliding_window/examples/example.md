📚 Lista de Exercícios: Sliding Window (Janela Deslizante)
1. Maximum Average Subarray I (LeetCode 643) - Fácil
Problema: Dado um array nums e um inteiro k, encontre o sub-array contíguo de tamanho k que possui a maior média.

🐍 Python
```Python
def findMaxAverage(nums, k):
    curr_sum = sum(nums[:k])
    max_sum = curr_sum
    for i in range(k, len(nums)):
        curr_sum += nums[i] - nums[i-k]
        max_sum = max(max_sum, curr_sum)
    return max_sum / k
```
🐹 Go

```Go
func findMaxAverage(nums []int, k int) float64 {
    sum := 0
    for i := 0; i < k; i++ { sum += nums[i] }
    maxSum := sum
    for i := k; i < len(nums); i++ {
        sum += nums[i] - nums[i-k]
        if sum > maxSum { maxSum = sum }
    }
    return float64(maxSum) / float64(k)
}
```
2. Longest Substring Without Repeating Characters (LeetCode 3) - Médio
Problema: Encontre o comprimento da maior substring sem caracteres repetidos.

🐍 Python
```Python
def lengthOfLongestSubstring(s):
    char_map = {}
    left = max_len = 0
    for right, char in enumerate(s):
        if char in char_map and char_map[char] >= left:
            left = char_map[char] + 1
        char_map[char] = right
        max_len = max(max_len, right - left + 1)
    return max_len
```
🐹 Go
```Go
func lengthOfLongestSubstring(s string) int {
    charMap := make(map[byte]int)
    left, maxLen := 0, 0
    for right := 0; right < len(s); right++ {
        if idx, found := charMap[s[right]]; found && idx >= left {
            left = idx + 1
        }
        charMap[s[right]] = right
        if dist := right - left + 1; dist > maxLen { maxLen = dist }
    }
    return maxLen
}
```

3. Max Consecutive Ones III (LeetCode 1004) - Médio
Problema: Dado um array binário e um inteiro k, retorne o número máximo de 1s consecutivos se puder inverter no máximo k zeros.

🐍 Python
```Python
def longestOnes(nums, k):
    left = 0
    for right in range(len(nums)):
        if nums[right] == 0: k -= 1
        if k < 0:
            if nums[left] == 0: k += 1
            left += 1
    return len(nums) - left
```
🐹 Go

```Go
func longestOnes(nums []int, k int) int {
    left := 0
    for right := 0; right < len(nums); right++ {
        if nums[right] == 0 { k-- }
        if k < 0 {
            if nums[left] == 0 { k++ }
            left++
        }
    }
    return len(nums) - left
}
```

4. Minimum Size Subarray Sum (LeetCode 209) - Médio
Problema: Dado um array de inteiros positivos e um target, encontre o comprimento mínimo de um subarray cuja soma seja maior ou igual ao target.

🐍 Python
```Python
def minSubArrayLen(target, nums):
    left = total = 0
    res = float('inf')
    for right in range(len(nums)):
        total += nums[right]
        while total >= target:
            res = min(res, right - left + 1)
            total -= nums[left]
            left += 1
    return res if res != float('inf') else 0
```
🐹 Go

 ```   Go
    func minSubArrayLen(target int, nums []int) int {
        left, total, res := 0, 0, len(nums)+1
        for right := 0; right < len(nums); right++ {
            total += nums[right]
            for total >= target {
                if dist := right - left + 1; dist < res { res = dist }
                total -= nums[left]
                left++
            }
        }
        if res > len(nums) { return 0 }
        return res
    }
```

5. Permutation in String (LeetCode 567) - Médio
Problema: Dadas duas strings s1 e s2, retorne true se s2 contiver uma permutação de s1.

🐍 Python
```Python
    from collections import Counter
    def checkInclusion(s1, s2):
        cnt1, w, k = Counter(s1), len(s1), len(s1)
        for i in range(len(s2)):
            if s2[i] in cnt1: 
                if cnt1[s2[i]] > 0: k -= 1
                cnt1[s2[i]] -= 1
            if i >= w:
                if s2[i-w] in cnt1:
                    if cnt1[s2[i-w]] >= 0: k += 1
                    cnt1[s2[i-w]] += 1
            if k == 0: return True
        return False
```
🐹 Go
```Go
    func checkInclusion(s1 string, s2 string) bool {
        if len(s1) > len(s2) { return false }
        var count [26]int
        for i := 0; i < len(s1); i++ {
            count[s1[i]-'a']++
            count[s2[i]-'a']--
        }
        if isZero(count) { return true }
        for i := len(s1); i < len(s2); i++ {
            count[s2[i]-'a']--
            count[s2[i-len(s1)]-'a']++
            if isZero(count) { return true }
        }
        return false
    }
    func isZero(arr [26]int) bool {
        for _, v := range arr { if v != 0 { return false } }
        return true
    }
```

Dica Final: Se precisar de ajuda para entender a lógica de algum desses especificamente, é só me chamar! Quer que eu explique o passo a passo de algum deles?