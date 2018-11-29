/**
 * @param {number} num
 * @return {boolean}
 */
var isUgly = function(num) {
    copy = num;
    if (num <= 0) return false;
    if (num ===1) return true;
    while(num!==1 && num%2==0) num/=2;
    while(num!==1 && num%3==0) num/=3;
    while(num!==1 && num%5==0) num/=5;
    return num===1;
};

arr = [-3,-1,0,1,2,3,4,5,6,7,6,8,14,22,23,455,4564,10000,900,810000,810001];
console.log(arr.map(isUgly));