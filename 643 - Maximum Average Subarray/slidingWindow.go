func findMaxAverage(nums []int, k int) float64 {
    sum := 0.0
    for i:=0; i < k; i++ {
        sum += float64(nums[i])
    }

    max_avg := sum / float64(k)

    for j:= k; j < len(nums); j++ {
        sum += float64(nums[j])
        sum -= float64(nums[j-k])

        avg := sum / float64(k)


        if avg > max_avg {
            max_avg = avg
        }
    }

    return max_avg
}