class Solution:
    def maxProfit(self, prices: List[int]) -> int:
        max = 0
        min = prices[0]

        for price in prices:
            if price < min:
                min = price

            if (price - min) > max:
                max = (price - min)
        
        return max
        