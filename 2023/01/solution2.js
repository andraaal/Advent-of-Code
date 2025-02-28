"use strict";

const fs = require('fs');

const data = fs.readFileSync('./input.txt', 'utf8').split('\r\n');

let result = 0;

const map = {
    one: 1,
    two: 2,
    three: 3,
    four: 4,
    five: 5,
    six: 6,
    seven: 7,
    eight: 8,
    nine: 9
}

data.forEach(line => {
    let matches = line.matchAll(/(?=(one|two|three|four|five|six|seven|eight|nine|[1-9]))/g);
    let match = Array.from(matches, (m) => m[1]);
    console.log(line);
    console.log(match);
    let resultWithString = match.at(0) + match.at(-1);
    let resultWithoutString = resultWithString.replace(/one|two|three|four|five|six|seven|eight|nine/g, m => {
        return map[m];
    });
    result += parseInt(resultWithoutString);

});

console.log(result);