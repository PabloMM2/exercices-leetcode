function merge(intervals: number[][]): number[][] {

    if(intervals.length == 0) {
        return intervals
    }

    intervals.sort((a, b) => a[0] - b[0])

    let response : number[][] = []
    let start : number = intervals[0][0]
    let end : number = intervals[0][1]

    for(let i = 1; i < intervals.length; i++) {
        if (intervals[i][0] > end) {
            response.push([start, end])
            start = intervals[i][0]
            end = intervals[i][1]
        } else if (intervals[i][1] > end) {
            end = intervals[i][1]
        }
    }

    response.push([start, end])
    
    return response
};