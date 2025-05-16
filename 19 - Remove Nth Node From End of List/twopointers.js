var removeNthFromEnd = function(head, n) {
    let node = new ListNode(0, head)
    let fast = node
    let slow = node

    for(let i=0; i < n; i++) {
        fast = fast.next
    }

    while(fast.next != null) {
        fast = fast.next
        slow = slow.next
    }

    slow.next = slow.next.next

    return node.next
};