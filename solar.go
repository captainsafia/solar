package solar

func IndexOf(a []int, item int) int {
    for index, element := range a {
        if element == item {
            return index
        }
    }
    return -1
}

func Intersection(a []int, b []int) []int {
    intersection := make([]int, 0)
    for _, element := range a {
        if IndexOf(b, element) != -1 {
            intersection = append(intersection, element)
        }
    }
    return intersection
}

func Union(a []int, b []int) []int {
    union := append(a, b...)
    return union
}

func Difference(a []int, b []int) []int {
    difference := make([]int, 0)
    for _, element := range a {
        if IndexOf(b, element) == -1 {
            difference = append(difference, element)
        }
    }
    return difference
}

func Replace(a []int, value int, replacement int) []int {
    b := make([]int, len(a))
    copy(b, a)
    for index, element := range a {
        if element == value {
            b[index] = replacement
        }
    }
    return b
}

func Pop(a []int) ([]int, int) {
    b := make([]int, len(a))
    copy(b, a)
    element := b[0]
    return append(b[:0], b[1:]...), element
}
