class Solution:
    def isSubsequence(self, s: str, t: str) -> bool:
        reg = ".*"

        for i in s:
            reg += i+".*"
        
        if re.match(reg, t):
            return True

        return False
        