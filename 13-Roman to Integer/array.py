class Solution:
    def romanToInt(self, s: str) -> int:
        rm = {"I": 1, "V": 5, "X": 10, "L": 50, "C": 100, "D": 500, "M": 1000}
        total = rm[s[len(s) -1]]
        last = s[len(s)-1]

        for i in range(len(s)-2, -1, -1):
            if rm[s[i]] >= rm[last]:
                total += rm[s[i]]
            else:
                total -= rm[s[i]]
            
            last = s[i]

        return total
        