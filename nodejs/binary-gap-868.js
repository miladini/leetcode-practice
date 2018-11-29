/**
 * @param {number} N
 * @return {number}
 */
'use strict';
var binaryGap = function(N) {
    if ((N&(N-1))===0) {
        return 0;
    }
    let count = 0;
    let max = -1;
    let first = -1;
    while(N!==0) {
        if (N&1===1) {
            if (first===-1) {                
                first = count;
            } else {
                max = Math.max(max,count-first);
                first = count;
            }
            
        }
        count++;
        N>>>=1;        
    }
    return max;
};

console.log(binaryGap(1));
console.log(binaryGap(2));
console.log(binaryGap(4));
console.log(binaryGap(8));
console.log(binaryGap(22));
console.log(binaryGap(5));
console.log(binaryGap(6));
console.log(binaryGap(8));
console.log(binaryGap(536870912));
console.log(binaryGap(536870913));