var middleNode = function(head) {

    let values = new Map();
    let current = head;
    let count = 0;

    while(current != null) {
        values.set(count, current)
        count++
        current = current.next
    }

    let middle = parseInt(count / 2)
    return values.get(middle)
};