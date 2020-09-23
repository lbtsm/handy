package slicehelper

func FindIdxOfIntSlice(value int, slice []int) int {
	if slice == nil {
		return -1
	}

	idx := -1
	for i, v := range slice {
		if v == value {
			return i
		}
	}

	return idx
}

func ExistOfIntSlice(value int, slice []int) bool {
	if slice == nil {
		return false
	}

	for _, v := range slice {
		if v == value {
			return true
		}
	}

	return false
}

func FindIdxOfInt8Slice(value int8, slice []int8) int {
	if slice == nil {
		return -1
	}

	idx := -1
	for i, v := range slice {
		if v == value {
			return i
		}
	}

	return idx
}

func ExistOfInt8Slice(value int8, slice []int8) bool {
	if slice == nil {
		return false
	}

	for _, v := range slice {
		if v == value {
			return true
		}
	}

	return false
}

func FindIdxOfInt16Slice(value int16, slice []int16) int {
	if slice == nil {
		return -1
	}

	idx := -1
	for i, v := range slice {
		if v == value {
			return i
		}
	}

	return idx
}

func ExistOfInt16Slice(value int16, slice []int16) bool {
	if slice == nil {
		return false
	}

	for _, v := range slice {
		if v == value {
			return true
		}
	}

	return false
}

func FindIdxOfInt32Slice(value int32, slice []int32) int {
	if slice == nil {
		return -1
	}

	idx := -1
	for i, v := range slice {
		if v == value {
			return i
		}
	}

	return idx
}

func ExistOfInt32Slice(value int32, slice []int32) bool {
	if slice == nil {
		return false
	}

	for _, v := range slice {
		if v == value {
			return true
		}
	}

	return false
}

func FindIdxOfInt64Slice(value int64, slice []int64) int {
	if slice == nil {
		return -1
	}

	idx := -1
	for i, v := range slice {
		if v == value {
			return i
		}
	}

	return idx
}

func ExistOfInt64Slice(value int64, slice []int64) bool {
	if slice == nil {
		return false
	}

	for _, v := range slice {
		if v == value {
			return true
		}
	}

	return false
}

func FindIdxOfStringSlice(value string, slice []string) int {
	if slice == nil {
		return -1
	}

	idx := -1
	for i, v := range slice {
		if v == value {
			return i
		}
	}

	return idx
}

func ExistOfStringSlice(value string, slice []string) bool {
	if slice == nil {
		return false
	}

	for _, v := range slice {
		if v == value {
			return true
		}
	}

	return false
}
