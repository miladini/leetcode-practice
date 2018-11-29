/**
 * @param {string} S
 * @return {string}
 */
'use strict'
var reverseOnlyLetters = function(S) {
    let arr = S.split('');
    let left  = 0;
    let length = S.length;
    let right = length-1;

    while(left<length && !(arr[left].match(/[A-Za-z]/) )) left++;
    while(right >= 0 && !(arr[right].match(/[A-Za-z]/) )) right--;
    while(left < right) {
        let tmp = arr[left];
        arr[left] = arr[right];
        arr[right] = tmp;
        left++;
        right--;
        while(left<length && !(arr[left].match(/[A-Za-z]/) )) left++;
        while(right >= 0 && !(arr[right].match(/[A-Za-z]/) )) right--;
    }
    return arr.join('');
};


console.log(reverseOnlyLetters("S%x^l*a!m"));
