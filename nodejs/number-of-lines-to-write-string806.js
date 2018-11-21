'use strict';
/**
 * @param {number[]} widths
 * @param {string} S
 * @return {number[]}
 */
var numberOfLines = function(widths, S) {
    let count = 0;
    let sum = 0;

    for (let i = 0; i < S.length; i++) {
        let val = widths[S.charCodeAt(i)-97];
        if (sum+val>100) {
            count++;
            sum = 0;
        }

        sum += val;
    }
    return [count+1,sum];
};


let widths = [4,10,10,10,10,10,10,10,10,10,10,10,10,10,10,10,10,10,10,10,10,10,10,10,10,10];
let S = "bbbcccdddaaa";
console.log(numberOfLines(widths,S));