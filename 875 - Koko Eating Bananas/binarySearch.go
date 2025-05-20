func minEatingSpeed(piles []int, h int) int {
    l, r := 1, Max(piles)
    for l < r {
        m := int((l+r)/2)
        min_h := 0
        for _, pile := range piles {
            min_h += int(math.Ceil(float64(pile)/ float64(m)))
        }

        if min_h <= h {
            r = m
        } else {
            l = m + 1
        }
    }
    return r
}

func Max(piles []int) int {
    max := 0
    for _, p := range piles {
        if p > max {
            max = p
        }
    }
    return max
}