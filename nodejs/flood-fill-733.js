'use strict';
/**
 * @param {number[][]} image
 * @param {number} sr
 * @param {number} sc
 * @param {number} newColor
 * @return {number[][]}
 */
var floodFill = function(image, sr, sc, newColor) {
    let visited = [];
    for (let i = 0; i < image.length; i++) {
        visited.push(new Array(image[0].length).fill(false));
    }
    dfs(image,visited,sr,sc,newColor,image[sr][sc]);
    return image;
};

var dfs  = function(image,visited,sr,sc,newColor,sColor) {
    image[sr][sc] = newColor;
    visited[sr][sc] = true;
    if (sr-1>=0 && !visited[sr-1][sc] && image[sr-1][sc]==sColor) {
        dfs(image,visited,sr-1,sc,newColor,sColor);
    }
    if (sr+1<image.length && !visited[sr+1][sc] && image[sr+1][sc]==sColor) {
        dfs(image,visited,sr+1,sc,newColor,sColor);
    }
    if (sc-1>=0 && !visited[sr][sc-1] && image[sr][sc-1]==sColor) {
        dfs(image,visited,sr,sc-1,newColor,sColor);
    }
    if (sc+1<image[0].length && !visited[sr][sc+1] && image[sr][sc+1]==sColor) {
        dfs(image,visited,sr,sc+1,newColor,sColor);
    }
}

// let image = [[1,1,1],[1,1,0],[1,0,1]];
let image = [[1]];
let sr = 0;
let sc = 0;
let newColor = 2;

floodFill(image, sr, sc, newColor);
console.log(image);