'use strict';
/**
 * @param {number} num
 * @return {string}
 */
var convertToBase7 = function (num) {
    if (num === 0) return 0;
    return num < 0 ? '-' + convertPosToBase7(-num) : convertPosToBase7(num);
};

var convertPosToBase7 = function (num) {

    let res = '';
    while (num !== 0) {
        res = (num % 7) + res;
        num = Math.floor(num / 7);
    }

    return res;
}

console.log(convertToBase7(77));
console.log(convertToBase7(11));
console.log(convertToBase7(700));
console.log(convertToBase7(1));
console.log(convertToBase7(2));
console.log(convertToBase7(0));
console.log(convertToBase7(-1));
console.log(convertToBase7(-7));
console.log(convertToBase7(-77));