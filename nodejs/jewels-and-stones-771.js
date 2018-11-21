'use strict';
/**
 * @param {string} J
 * @param {string} S
 * @return {number}
 */
var numJewelsInStones = function(J, S) {
    let jSet = new Set(J);
    let intersectionSet = [...S].map(x=>jSet.has(x)?1:0).reduce((acc,val)=> acc+val,0);
    return intersectionSet;
};

console.log(numJewelsInStones("fafaxp","fffffffff"));