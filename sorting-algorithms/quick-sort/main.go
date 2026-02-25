package main

func quicksort(arr []int, left int, right int) {

	if left < right {
		//fmt.Println(arr[left : right+1])
		pi := partition(arr, left, right)

		quicksort(arr, left, pi-1)
		quicksort(arr, pi+1, right)

	}

}

func partition(arr []int, left int, rigth int) int {
	pivot := arr[rigth]

	i := left - 1

	for j := left; j < rigth; j++ {
		if arr[j] <= pivot {
			i++
			tmp := arr[i]
			arr[i] = arr[j]
			arr[j] = tmp

		}
		tmp := arr[i+1]
		arr[i+1] = arr[rigth]
		arr[rigth] = tmp

	}
	return i + 1
}

func main() {
	arr := []int{0, 3, 6, 7, 8, 4, 2, 1, 5}
	quicksort(arr, 0, len(arr)-1)

}
