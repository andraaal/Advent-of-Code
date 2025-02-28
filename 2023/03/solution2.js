"345"*"234";
"use strict";

const fs = require('fs');

let data = fs.readFileSync('./input.txt', 'utf8').replace(/\r\n/gm, "");

let result = 0;

const gears = [];
let startIndex = 0;
let index = 0;
while((index = data.indexOf('*', startIndex)) > -1) {
    gears.push(index);
    startIndex = index + 1;
}

console.log(gears);

let gearIndex;

for (gearIndex of gears) {
    const index = gearIndex;

    console.log(index);
    console.log(data.slice(index - 1 - 140, index + 2 - 140));
    console.log(data.slice(index - 1, index + 2));
    console.log(data.slice(index - 1 + 140, index + 2 + 140));

    
    let matchResults = [];


    matchResults = matchResults.concat(data.slice(index - 1 - 140, index + 2 - 140).match(/[0-9]*/g));
    
    matchResults = matchResults.concat(data.slice(index - 1 + 140, index + 2 + 140).match(/\[0-9]*/g));

    matchResults = matchResults.concat(data[index - 1].match(/\[0-9]*/g));

    matchResults = matchResults.concat(data[index + 1].match(/\[0-9]*/g));

    console.log(matchResults);
    if (matchResults.length === 2) {
        result += parseInt(matchResults[0]) * parseInt(matchResults[1]);
    }
};

console.log(result);
console.log(data);