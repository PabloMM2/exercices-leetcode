class Solution:
    def merge(self, intervals: List[List[int]]) -> List[List[int]]:

        if len(intervals) == 0:
            return intervals

        intervals.sort(key=lambda x: x[0])
        response : List[List[int]] = []
        start = intervals[0][0]
        end = intervals[0][1]

        for i in range(1, len(intervals)):
            if intervals[i][0] > end:
                response.append([start, end])
                start = intervals[i][0]
                end = intervals[i][1]
            elif intervals[i][1] > end:
                end = intervals[i][1]

        response.append([start, end])

        return response
        