package app

// BinarySearch : 二分搜尋法 O(log n)
func BinarySearch(list interface{}) {

}

// QuickSort : 快速排序法 O(nlog n)
type QuickSort interface {
	Qsort()
}

// QSIntSlice : 快排 []int
type QSIntSlice struct{}

// Qsort :
func (QSIntSlice) Qsort(list []int) {

}

// QuickSort : 快速排序法 O(nlog n)
// func QuickSort(list ...interface{}) interface{} {
// 	if len(list) <= 1 {
// 		return list
// 	}
// 	leftList := make()
// 	return list
// }
