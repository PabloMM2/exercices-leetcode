class Solution:
    def isSubsequence(self, s: str, t: str) -> bool:

        last = 0
        counter = 0

        for i in s:
            for j in range(last, len(t)):
                last += 1
                if i == t[j]:
                    counter += 1
                    break

        return counter == len(s)
        