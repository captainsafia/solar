package solar

import "testing"

func TestIndexOf(t *testing.T) {
    index := IndexOf([]int{1, 2, 3, 4}, 3)
    if index != 2 {
        t.Error("Should have returned an index of 2.")
    }
}

func TestUnion(t *testing.T) {
    union := Union([]int{1, 2, 3}, []int{4, 5})
    if len(union) != 5 {
        t.Error("Should have returned a length of 5.")
    }
}

func TestIntersection(t *testing.T) {
    intersection := Intersection([]int{1, 2, 3}, []int{1, 2})
    if len(intersection) != 2 {
        t.Error("Should have returned a length of 2.")
    }
}

func TestDifference(t *testing.T) {
    difference := Difference([]int{1, 2, 3}, []int{1, 2})
    if len(difference) != 1 {
        t.Error("Should have returned a length of 1.")
    }
}

func TestReplace(t *testing.T) {
    replaced := Replace([]int{1, 2, 2, 2, 3}, 2, 4)
    for index, element := range replaced {
        if element == 2 {
            t.Error("Should have replaced element at index", index)
        }
    }
}

func TestPop(t *testing.T) {
    elements, popped := Pop([]int{1, 2, 3, 4})
    if popped != 1 && len(elements) != 3 {
        t.Error("Did not pop the first element from the list")
    }
}
