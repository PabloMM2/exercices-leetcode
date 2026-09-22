class Solution:
    def hasCycle(self, head: Optional[ListNode]) -> bool:
        nodes = {}
        while head != None:
            if nodes.get(head):
                return 1
            
            nodes[head] = 1
            head = head.next
        
        return 0