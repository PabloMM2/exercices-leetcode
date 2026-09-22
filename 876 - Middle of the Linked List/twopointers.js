// Definition for singly-linked list.
function ListNode(val, next) {
    this.val = (val===undefined ? 0 : val)
    this.next = (next===undefined ? null : next)
}

/**
 * @param {ListNode} head
 * @return {ListNode}
 */
var middleNode = function(head) {
    let start = new ListNode(0, head);
    let slow = start;
    let fast = start;

    while(fast != null){
        slow = slow.next;
        fast = fast.next == null ? fast.next : fast.next.next;
    }
    
    return slow
};