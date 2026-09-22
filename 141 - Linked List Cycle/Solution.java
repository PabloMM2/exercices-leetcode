public class Solution {
    public boolean hasCycle(ListNode head) {
        HashMap<ListNode, Boolean> nodes = new HashMap();
        while( head != null) {
            if(nodes.get(head) != null) {
                return true;
            }
            nodes.put(head, true);
            head = head.next;
        }
        return false;
    }
}