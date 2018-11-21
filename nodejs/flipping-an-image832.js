const { performance } = require('perf_hooks');
/**
 * @param {number[][]} A
 * @return {number[][]}
 */
var flipAndInvertImage = function (A) {
    'use strict';

    for (let i = 0, L = A.length; i < L; i++) {
        for (let j = 0, X = A[i].length;j < X-1-j; j++) {
            let tmp = A[i][j];
            A[i][j] = A[i][X-1-j];
            A[i][X-1-j] = tmp;
        }

        for (let j = 0, X = A[i].length;j < X; j++) {
            A[i][j] = -A[i][j]+1;
        }
    }
    return A;
    /*
    A.forEach((a, b, c) => {
        for (let i = 0, L=a.length, X = L - i - 1; i < X; i++) {
            let tmp = a[i];
            a[i] = a[X];
            a[X] = tmp;
            X = L - i - 1;
        }
        a.forEach((vv, i, cc) => {
            a[i] = -a[i] + 1;
        });
    });
    return A;
    */
};

const t1 = performance.now();
for (let i = 0; i < 100000000; i++) {

    flipAndInvertImage([[1, 1, 1], [1, 1, 0], [1, 0, 0]]);
}
console.log(t1, performance.now() - t1);
// flipAndInvertImage([1,1,1,1,0,0,1]);