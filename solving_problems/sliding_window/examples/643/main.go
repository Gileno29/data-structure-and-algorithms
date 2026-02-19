package main

// Maximum Average Subarray I (LeetCode 643) -
//  Fácil Problema: Dado um array nums e um inteiro k, encontre o sub-array contíguo de tamanho k que possui a maior média.

func FoudHighAvg(arr []int, k int) float64 {

	sum := 0

	for i := 0; i < k; i++ {
		sum += arr[i]
	}

	maxSum := sum

	for i := k; i < len(arr); i++ {
		sum = arr[i] - arr[i-k]

		if sum > maxSum {
			maxSum = sum
		}

	}

	return float64(maxSum) / float64(k)

}

func main() {
	arr := []int{10, 20, 11, 15, 22, 13}

	FoudHighAvg(arr, 3)

}
