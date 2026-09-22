class Solution {
    public ListNode middleNode(ListNode head) {
        Map<Integer, ListNode> values = new HashMap<>();
        int counter = 1;

        while(head != null) {
            values.put(counter, head);
            head = head.next;
            counter++;
        }

        int index = (int) Math.round((double)counter/2);
        return values.get(index);
    }
}