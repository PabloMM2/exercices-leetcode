/**
 * Definition for singly-linked list.
 * class ListNode {
 *     val: number
 *     next: ListNode | null
 *     constructor(val?: number, next?: ListNode | null) {
 *         this.val = (val===undefined ? 0 : val)
 *         this.next = (next===undefined ? null : next)
 *     }
 * }
 */

function removeNthFromEnd(head: ListNode | null, n: number): ListNode | null {
    let node : ListNode = new ListNode(0, head);
    let slow : ListNode = node
    let fast : ListNode = node

    for(let i = 0; i < n; i++) {
        fast = fast.next
    }

    while(fast.next != null) {
        fast = fast.next
        slow = slow.next
    }

    slow.next = slow.next.next;

    return node.next;
};