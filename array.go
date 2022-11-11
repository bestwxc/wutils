package wutils

// GetArrayIndexOf 获取元素在数组中的序号
/*
func GetArrayIndexOf(a []any, item any) int {
	if len(a) <= 0 {
		return -1
	}
	if item == nil {
		return -1
	}
	for i := 0; i < len(a); i++ {
		if a[i] == item {
			return i
		}
	}
	return -1
}
*/

// IsByteArrayEqual 判断数组是否相等
func IsByteArrayEqual(a1 []byte, a2 []byte) bool {
	if len(a1) != len(a2) {
		return false
	} else if len(a1) == 0 {
		if a1 == nil && a2 == nil {
			return true
		}
		if a1 != nil && a2 != nil {
			return true
		}
		return false
	} else {
		for i := 0; i < len(a1); i++ {
			if a1[i] != a2[i] {
				return false
			}
		}
		return true
	}
}
