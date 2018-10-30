'use strict';
/**
 * @param {number} num
 * @return {number}
 */
var findComplement = function(num) {    
    let length = Math.floor(1+Math.log2(num));
    return ((1<<length)-1) ^ num;
};
console.log(findComplement(1<<5));

