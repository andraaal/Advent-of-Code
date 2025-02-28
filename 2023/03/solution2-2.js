"use strict";

const fs = require('fs');

let data = fs.readFileSync('./input.txt', 'utf8').replace(/\r\n/gm, "");

let result = 0;

/*const numbers = []
data.replace(/[0-9]+/g, (match, index) => {
    numbers.push([index, match]);
    return '.'.repeat(match[0].length);
});*/

const gears = [];
let gear;

data.replace(/\*/gm, (_match, index) => {
    gears.push(index);
    return '*';
})

for (gear of gears) {
    
    console.log(gear);
    console.log(data.slice(gear - 1 - 140, gear + 2 - 140));
    console.log(data.slice(gear - 1, gear + 2));
    console.log(data.slice(gear - 1 + 140, gear + 2 + 140));

    //checking above number
    if (data.slice(gear - 1 - 140, gear + 1 - 140).replace(/\.|[0-9]/g, "")) {
        result += parseInt('*');
    }

    //checking below number
    if (data.slice(gear - 240, gear + 1 + 140).replace(/\.|[0-9]/g, "")) {
        result += parseInt('*');
    }

    if (data[gear - 1].replace(/\.|[0-9]/g, "")) {
        result += parseInt('*');
    }
    if (data[gear + 1].replace(/\.|[0-9]/g, "")) {
        result += parseInt('*');
    }
};

console.log(result);
console.log(data);

function findNumber(index) {
    let resultNumber = [data[index]];
    let i = index;
    while(data[i - 1].match(/[0-9]/)) {
        resultNumber.unshift(data[i - 1]);
        i--;
    }
    i = index;
    while(data[i + 1].match(/[0-9]/)) {
        resultNumber.push(data[i + 1]);
        i++;
    }
    return [i === index, resultNumber.join("")];
}