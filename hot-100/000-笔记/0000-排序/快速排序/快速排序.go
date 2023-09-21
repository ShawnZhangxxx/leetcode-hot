package main

func quickSort(arr []int, low, high int) {
	if low < high {
		pi := partition(arr, low, high)
		quickSort(arr, low, pi-1)
		quickSort(arr, pi+1, high)
	}
}

func partition(arr []int, low, high int) int {
	pivot := arr[high] //选最大的为pivot , //TODO 可以用荷兰国旗问题优化, 两遍往里压,等于的数在中间,经典的随机快排重复值只搞定一个
	i := low - 1
	for j := low; j < high; j++ {
		if arr[j] <= pivot {
			i++
			arr[i], arr[j] = arr[j], arr[i]
		}
	}
	arr[i+1], arr[high] = arr[high], arr[i+1]
	return i + 1
}

func main() {
	partition([]int{1,2,3,1},0,3)
	//arr := []int{7, 2, 6,1, 6, 8, 5, 5,3,2,2,2, 4}
	//fmt.Println("Before sorting:", arr)
	//quickSort(arr, 0, len(arr)-1)
	//fmt.Println("After sorting:", arr)
}