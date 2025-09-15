func firstBadVersion(n int) int {
	l, r := 1, n

	for l <= r {
		if isBadVersion(l) {
			return l
		}
		m := int((l + r) / 2)
		if !isBadVersion(m) {
			l = m + 1
		} else {
			r = m
		}
	}

	return 1
}

//--------- Second version ---------

func firstBadVersion(n int) int {
    l := 1, r := n
    for l < r {
        m := int((l+r)/2)

        if isBadVersion(m) {
            if !isBadVersion(m -1){
                return m
            }
            r = m -1
        } else {
            l = m +1
        }
    }
    return l
}