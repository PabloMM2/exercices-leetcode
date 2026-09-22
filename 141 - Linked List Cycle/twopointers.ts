function hasCycle(head: ListNode | null): boolean {
    let slow : ListNode = head;
    let fast : ListNode = head;
    while(fast != null && fast.next != null){
        slow = slow.next;
        fast = fast.next.next;
        if( slow == fast ) {
            return true;
        }
    }
    return false;
};