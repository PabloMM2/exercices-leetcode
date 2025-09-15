function productExceptSelf(nums: number[]): number[] {

    let responses : number[] = [];
    let result = 1;
    let left : number[] = [];
    let right : number[] = [];

    nums.forEach(function(element, index) {
        result *= element
        left.push(result)
    })

    result = 1
    for (let i = (nums.length -1); i >=0; i--) {
        result *= nums[i]
        right.unshift(result)
    }

    for (let j = 0; j < nums.length; j++) {
        if (j == 0) {
            responses.push(right[j + 1])
        } else if (j == nums.length -1) {
            responses.push(left[j -1])
        } else {
            responses.push(left[j -1] * right[j+1])
        }
    }

    return responses
};