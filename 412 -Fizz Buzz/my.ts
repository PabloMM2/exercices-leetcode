function fizzBuzz(n: number): string[] {
    let response : string[] = []

    for(let i = 1; i <= n; i++) {
        if( i % 5 == 0 && i % 3 == 0) {
            response.push("FizzBuzz")
        } else if( i % 5 == 0) {
            response.push("Buzz")
        } else if( i % 3 == 0) {
            response.push("Fizz")
        } else {
            response.push(`${i}`)
        }
    }
    return response
};