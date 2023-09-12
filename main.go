package main

import (
	"fmt"
)

func output_array(array []int) {
	fmt.Println(array)
}
func bubble_sort(array []int, sz int) {
	for i := 0; i < sz-1 ; i++ {
		for j := 0; j < sz-1-i; j++ {
			if array[j] > array[j+1] {
				array[j], array[j+1] = array[j+1], array[j]
			}
		}
	}
}
func select_sort(arr []int, sz int) {
	var i int = 0
	var j int = 0
	for i = 0; i < sz; i++ {
		var min int = i
		for j=i+1; j < sz; j++ {
			if arr[j] < arr[min] {
				min = j
			}
		}
		arr[i], arr[min] = arr[min], arr[i]
	}
}
func insert_sort(arr []int, sz int) {
	sortNum := 0
	i := 0
	j := 0
	for i=0; i<sz; i++ {
		sortNum = arr[i]
		// 找插入的地方
		for j=i-1; j>=0; j-- {
			if arr[j] < sortNum {
				break
			}
			arr[j+1] = arr[j]
		}
		arr[j+1] = sortNum
	}
}
func shell_sort_in_group(arr []int, sz int, pos int, step int) {
	sortNum := 0
	i := 0
	j := 0
	for i=pos+step; i<sz; i+=step {
		// pos + step 就是该组的第一个元素的位置
		// 例如：当pos==0的时候，该组的第一个元素就是整个序列的第一个元素
		// 当pos==1的时候，该组第一个元素就是这个那个序列的第二个元素
		sortNum = arr[i]
		for j=i-step; j>=0; j-=step {
			if arr[j] < sortNum {
				break
			}
			arr[j+step] = arr[j]
		}
		arr[j+step] = sortNum
	}
}
func shell_sort(arr []int, sz int) {
	i := 0
	step := 0 // 希尔间隔
	for step = sz/2; step>0; step/=2 {
		// 步长为step，即共有istep个组，对每一组都执行插入排序
		for i=0; i<step; i++ {
			// 每一组的组排序就是插入排序的稍微修改的版本
			shell_sort_in_group(arr, sz, i, step)
		}
	}
}
func quick_sort_part_sort1(arr []int, begin int, end int) int {
	// Hoare思想
	left := begin
	right := end
	keyi := left
	for left < right {
		// 右边走，找小
		for left < right && arr[right] >= arr[keyi] {
			right--
		}
		// 左边走，找大
		for left < right && arr[left] <= arr[keyi] {
			left++
		}
		arr[left], arr[right] = arr[right], arr[left]
	}
	arr[keyi], arr[right] = arr[right], arr[keyi]
	keyi = left
	return keyi
}
func quick_sort_part_sort2(arr []int, begin int, end int) int {
	// 挖坑法
	key := arr[begin]
	piti := begin
	for begin < end {
		// 右边找小的，填到左边的坑里面去，这个位置形成新的坑
		for begin < end && arr[end] >= key {
			end--
		}
		arr[piti] = arr[end]
		piti = end

		// 左边找大的，填到右边的坑里面去，这个位置形成新的坑
		for begin < end && arr[begin] <= key {
			begin++
		}
		arr[piti] = arr[begin]
		piti = begin;
	}
	arr[piti] = key
	return piti
}
func quick_sort_by_interval(arr []int, begin int, end int) {
	if begin >= end {
		return // 区间不存在
	}
	keyi := quick_sort_part_sort1(arr, begin, end)
	quick_sort_by_interval(arr, begin ,keyi-1)
	quick_sort_by_interval(arr, keyi+1, end)
}
func quick_sort(arr []int, sz int) {
	begin := 0
	end := sz - 1
	quick_sort_by_interval(arr, begin ,end)
}
func heap_sort_adjust_down(arr []int, size int, parent int) {
	child := parent*2+1
	for child < size {
		// 选出左右孩子中小/大的那个
		if child + 1 < size && arr[child+1] > arr[child] {
			child++
		}
		// 孩子和父亲比较
		if arr[child] > arr[parent] {
			arr[child], arr[parent] = arr[parent], arr[child]
			parent = child
			child = parent*2 + 1
		} else {
			break
		}
	}
}
func heap_sort(arr []int, sz int) {
	for i:=(sz-1-1)/2; i>=0; i-- {
		heap_sort_adjust_down(arr, sz, i)
	}
	// 
	end := sz-1
	for end > 0 {
		arr[0],arr[end] = arr[end],arr[0]
		heap_sort_adjust_down(arr, end, 0)
		end--
	}
}
func merge_sort_by_group(arr []int, begin int, end int, tmp_array []int){
	if begin >= end {
		return
	}
	mid := (begin + end) / 2
	// [begin,mid][mid+1,end]分治递归，让子区间有序
	merge_sort_by_group(arr, begin, mid, tmp_array)
	merge_sort_by_group(arr, mid+1, end, tmp_array)
	// 归并[begin,mid][mid+1,end]
	begin1 := begin; end1 := mid;
	begin2 := mid+1; end2 := end;
	i := begin1
	for begin1 <= end1 && begin2 <= end2 {
		if arr[begin1] < arr[begin2] {
			tmp_array[i] = arr[begin1]
			i++
			begin1++
		} else {
			tmp_array[i] = arr[begin2]
			i++
			begin2++
		}
	}
	for begin1 <= end1 {
		tmp_array[i] = arr[begin1]
		i++
		begin1++
	}
	for begin2 <= end2 {
		tmp_array[i] = arr[begin2]
		i++
		begin2++
	}
	//把归并后的数据拷贝回原来的数组
	copy(arr[begin:end+1], tmp_array[begin:end+1])
}
func merge_sort(arr []int, sz int) {
	tmp_array := make([]int, sz) // 创建一个大小为sz的整数数组并初始化为0
	merge_sort_by_group(arr, 0, sz-1, tmp_array)
}

var (
	K int = 3 // 基数排序需要的全局变量
	RADIX int = 10
	queue [][]int
)
func radix_sort_queue_pop(qu []int) []int {
    if len(qu) == 0 {
        return qu // 如果数组为空，不做任何操作
    }
    // 删除第一个元素
    qu = qu[1:]
    return qu
}
func radix_sort_queue_push(qu []int, data int) []int {
	qu = append(qu, data)
	return qu
}
func radix_sort_get_key(value int, k int) int {
	key := 0
	for k >= 0 {
		key = value % 10
		value /= 10
		k--
	}
	return key
}
func radix_sort_distribute(arr []int, left int, right int, k int){
	// k表示是第几次分发数据
	for i:=left; i<right; i++ {
		key := radix_sort_get_key(arr[i], k)
		queue[key] = radix_sort_queue_push(queue[key], arr[i])
	}
}
func radix_sort_collect(arr []int) {
	k := 0
	for i:=0; i < RADIX; i++ {
		for len(queue[i]) != 0 {
			arr[k] = queue[i][0] // 先进先出
			k++
			queue[i] = radix_sort_queue_pop(queue[i])
		}
	}
}
func radix_sort_by_group(arr []int, left int, right int) {
	for i:=0; i<K; i++ {
		// 分发数据
		radix_sort_distribute(arr, left, right, i)
		// 回收数据
		radix_sort_collect(arr)
	}
}
func radix_sort(arr []int, sz int) {
	// 初始化队列
	queue = make([][]int, RADIX)
	left := 0
	right := sz;
	radix_sort_by_group(arr, left, right)
}


func TestSorts() {
	a := []int{1, 3, 5, 7, 9, 2, 4, 6, 8, 10, 11, 13, 12, 100, 14}
	len := len(a)
	output_array(a)
	radix_sort(a, len)
	output_array(a)
}
func others() {
	a := 10
	b := 5
	c := (a + b)/2
	fmt.Println(c)
}
func main() {
	TestSorts()
}