func topKFrequent(nums []int, k int) []int {

    var values = make(map[int]int)

    for _, val := range nums {
        values[val]++
    }

    response := []int{}
  
    for len(response) < k {
        max := 0
        big := 0      
        for key, val := range values {
            if val > max {
                max = val
                big = key
            }
        }

        response = append(response, big)
        delete(values, big)
    }

    return response
}

