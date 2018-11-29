/**
 * @param {number[]} candies
 * @return {number}
 */

 'use strict';
var distributeCandies = function(candies) {
    let set = new Set();
    for (let v of candies) set.add(v);
    return Math.min(candies.length/2, set.size);
};

let x = distributeCandies([1,1,2,3])
console.log(x)

