/**
 * @param {number[]} nums
 * @return {number}
 */
var singleNumber = function(nums) {
    let xor = 0;

    for (let x of nums) {
        xor ^=  x;
    }

    return xor;
};

console.log(singleNumber([2,3,4,6, 2,3,4]));

