package utils

import (
	"fmt"
	"sort"
)

type StringArray []string

func (s StringArray) Contains(item string) bool {

	if len(s) == 0 {
		return false
	}

	i, j := 0, len(s)-1
	for i <= j {
		if s[i] == item || s[j] == item {
			return true
		}
		i++
		j--
	}
	return false
}

func InterSection(n1 []int, n2 []int) bool {

	if len(n1) == 0 || len(n2) == 0 {
		return false
	}

	sort.Ints(n1)
	sort.Ints(n2)
	s1, s2, l1, l2 := 0, 0, len(n1), len(n2)
	var data []int
	m := map[int]bool{}
	for s1 < l1 && s2 < l2 {
		if m[n1[s1]] {
			s1++
			continue
		}
		if m[n2[s2]] {
			s2++
			continue
		}
		if n1[s1] == n2[s2] {
			data = append(data, n1[s1])
			m[n1[s1]] = true
			s2++
			s1++
		} else if n1[s1] > n2[s2] {
			s2++
		} else {
			s1++
		}
	}
	return len(data) > 0
}

func ArrPage(pageNumber, pageSize, totalCount int) (startIndex int, endIndex int) {
	totalPage := 0
	if totalCount%pageSize == 0 {
		totalPage = totalCount / pageSize
	} else {
		totalPage = totalCount/pageSize + 1
	}
	fmt.Println(totalPage)

	if pageNumber < 1 {
		pageNumber = 1
	}

	startIndex = (pageNumber - 1) * pageSize
	endIndex = startIndex + pageSize

	if endIndex >= totalCount {
		endIndex = totalCount
	}

	if startIndex >= totalCount {
		startIndex = 0
		endIndex = 0
	}

	return startIndex, endIndex
}
