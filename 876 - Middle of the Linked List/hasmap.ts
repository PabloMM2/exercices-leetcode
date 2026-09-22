// Definition for singly-linked list.

class ListNode {
    val: number
    next: ListNode | null
    constructor(val?: number, next?: ListNode | null) {
        this.val = (val===undefined ? 0 : val)
        this.next = (next===undefined ? null : next)
    }
}

function middleNode(head: ListNode | null): ListNode | null {
    let values = new Map<number, ListNode | null>();
    let counter : number = 1;

    while(head != null) {
        values.set(counter, head)
        head = head.next;
        counter++;
    }

    var index : number = Math.round(counter/2);
    return values.get(index);
};