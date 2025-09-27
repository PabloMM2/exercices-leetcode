function reverseString(s: string[]): void {
    let l : number = 0
    let r : number = s.length -1

    while(l < r) {
        let aux : string = s[l] 
        s[l] = s[r] 
        s[r] = aux
        l++
        r--
    }
};