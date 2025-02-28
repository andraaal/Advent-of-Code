"use strict";

const fs = require('fs');

let data = fs.readFileSync('./input.txt', 'utf8').replace(/\r\n/gm, "");

let result = 0;

const numbers = []
data.replace(/[0-9]+/g, (match, index) => {
    numbers.push([index, match]);
    return '.'.repeat(match[0].length);
});

let dataPair;

for (dataPair of numbers) {
    const index = dataPair[0];
    const number = dataPair[1];
    
    console.log(index, number);
    console.log(data.slice(index - 1 - 140, index + number.length + 1 - 140));
    console.log(data.slice(index - 1, index + number.length + 1));
    console.log(data.slice(index - 1 + 140, index + number.length + 1 + 140));

    //checking above number
    if (data.slice(index - 1 - 140, index + number.length + 1 - 140).replace(/\.|[0-9]/g, "")) {
        result += parseInt(number);
        continue;
    }

    //checking below number
    if (data.slice(index - 1 + 140, index + number.length + 1 + 140).replace(/\.|[0-9]/g, "")) {
        result += parseInt(number);
        continue;
    }

    if (data[index - 1].replace(/\.|[0-9]/g, "")) {
        result += parseInt(number);
        continue;
    }
    if (data[index + number.length].replace(/\.|[0-9]/g, "")) {
        result += parseInt(number);
        continue;
    }
};

console.log(result);
console.log(data);