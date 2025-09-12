class Solution:
    def productExceptSelf(self, nums: List[int]) -> List[int]:
        response: list[int] = []
        result = 1
        left: list[int] = []
        right: list[int] = []

        for num in nums:
            result *= num
            left.append(result)
        
        result = 1

        for i in range(len(nums) -1, -1, -1):
            result *= nums[i]
            right.insert(0,result)

        for i in range(len(nums)):
            if i == 0:
                response.append(right[i+1])
            elif i == len(nums) -1:
                response.append(left[i-1])
            else:
                response.append(left[i-1] * right[i+1])

        return response
        