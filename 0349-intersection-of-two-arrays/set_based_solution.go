package main

func intersection2(nums1 []int, nums2 []int) []int {
	if len(nums1) <= len(nums2) {
		set := BuildSet(nums1)
		return interSectionFromSet(nums2, set)
	}
	set := BuildSet(nums2)
	return interSectionFromSet(nums1, set)
}

func interSectionFromSet(nums []int, set *Set) []int {
	resp := make([]int, 0)
	current := NewSet()
	for _, num := range nums {
		if set.Has(num) && !current.Has(num) {
			resp = append(resp, num)
			current.Add(num)
		}
	}
	return resp
}

type Set struct {
	items map[int]struct{}
}

func NewSet() *Set {
	return &Set{
		items: make(map[int]struct{}),
	}
}

func BuildSet(nums []int) *Set {
	set := NewSet()
	for _, num := range nums {
		set.Add(num)
	}
	return set
}

func (s *Set) Has(x int) bool {
	_, exists := s.items[x]
	return exists
}

func (s *Set) Add(x int) {
	s.items[x] = struct{}{}
}
