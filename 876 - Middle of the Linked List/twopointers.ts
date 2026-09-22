// Definition for singly-linked list.

class ListNode {
    val: number
    next: ListNode | null
    constructor(val?: number, next?: ListNode | null) {
        this.val = (val===undefined ? 0 : val)
        this.next = (next===undefined ? null : next)
    }
}
 

function middleNode(head: ListNode | null) : ListNode | null {
    let start : ListNode | null = new ListNode(0, head);
    let slow : ListNode | null = start;
    let fast : ListNode | null = start;

    while(fast != null) {
        slow = slow.next;
        fast = fast.next == null ? fast.next : fast.next.next;
    }

    return slow
};