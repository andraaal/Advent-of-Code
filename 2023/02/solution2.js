"use strict";

const fs = require('fs');

const data = fs.readFileSync('./input.txt', 'utf8').split('\r\n');

let result = 0;

data.forEach(line => {
    const elements = line.split(': ');
    const gameId = parseInt(elements[0].slice(5));

    const games = elements[1].split('; ');
    console.log(games, gameId);
    let game;
    let minRed = 0;
    let minGreen = 0;
    let minBlue = 0;
    for(game of games) {
        minRed = Math.max(parseInt(game.match(/\d+(?= red)/) ?? 0), minRed);
        minGreen = Math.max(parseInt(game.match(/\d+(?= green)/) ?? 0), minGreen);
        minBlue = Math.max(parseInt(game.match(/\d+(?= blue)/) ?? 0), minBlue);
    };
    const power = minRed * minGreen * minBlue;
    console.log(power);
    result += power;
});

console.log(result);