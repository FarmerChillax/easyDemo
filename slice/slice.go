package gslice

// 计算两个 slice 的交集
func Intersection[T comparable](src1, src2 []T) []T {
	tmp := make(map[T]struct{})
	result := []T{}
	for _, item := range src1 {
		tmp[item] = struct{}{}
	}
	for _, item := range src2 {
		if _, ok := tmp[item]; ok {
			result = append(result, item)
		}
	}
	return result
}

// 判断两个切片是否存在交集
func IsIntersection[T comparable](src, target []T) bool {
	tmp := make(map[T]struct{})
	for _, item := range src {
		tmp[item] = struct{}{}
	}

	for _, item := range target {
		if _, ok := tmp[item]; ok {
			return true
		}
	}
	return false
}

func SliceToSet[T comparable](list ...[]T) []T {
	tmp := make(map[T]struct{})
	result := []T{}
	src := []T{}
	for _, item := range list {
		src = append(src, item...)
	}

	for _, item := range src {
		tmp[item] = struct{}{}
	}

	for key := range tmp {
		result = append(result, key)
	}

	return result
}
