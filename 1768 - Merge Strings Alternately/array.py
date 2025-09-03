class Solution:
    def mergeAlternately(self, word1: str, word2: str) -> str:

        longer = max(len(word1), len(word2))
        merged = ""
        for i in range(longer):

            if len(word1) > i :
                merged += word1[i]
            
            if len(word2) > i :
                merged += word2[i]
                    
        return merged
