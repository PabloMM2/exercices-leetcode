var copyRandomList = function(head) {
    let values = new Map();
    let current = head
    while(current != null) {
        let n = new Node(current.val, null, null)
        values.set(current, n)
        current = current.next
    }

    current = head
    while(current != null) {
        let n = values.get(current)
        n.next = values.get(current.next) || null
        n.random = values.get(current.random) || null
        current = current.next
    }
    
    return values.get(head)
};