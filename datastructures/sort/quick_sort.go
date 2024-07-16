package sort

func QuickSort(arr []int, left, right int) {
	if left >= right {
		return
	}
	i, j := left, right
	privot := arr[i] //privot就是我们单趟选择的分界值，一般我们是选择左右边界值，不选择中间值
	for i < j {
		//每次找到大于key或者是小于key的值就将 i, j 对应的值进行交换
		for i < j && arr[j] >= privot {
			j--
		}
		arr[i] = arr[j]
		for i < j && arr[i] <= privot {
			i++
		}
		arr[j] = arr[i]
	}
	//当for循环退出时，此时i的位置就是key值在排序后应该在的位置
	arr[i] = privot
	QuickSort(arr, left, i-1)  //递归将key左边的数组进行排序
	QuickSort(arr, i+1, right) ////递归将key右边的数组进行排序
}
