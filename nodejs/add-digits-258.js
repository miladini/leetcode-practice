/**
 * @param {number} num
 * @return {number}
 */
var addDigits = function(num) {
    'use strict';
    let numS = num.toString();
    while (numS.length > 1) {
        let sum = 0;
        for (let i = 0; i < numS.length; i++) {
            sum += parseInt(numS.substring(i,i+1));
        }
        numS = sum.toString();
    }
    return numS;
};

console.log(addDigits(8712387128));
