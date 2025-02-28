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
    for(game of games) {
        const red = parseInt(game.match(/\d+(?= red)/) ?? 0) <= 12;
        const green = parseInt(game.match(/\d+(?= green)/) ?? 0) <= 13;
        const blue = parseInt(game.match(/\d+(?= blue)/) ?? 0) <= 14;

        if(!(red && green && blue)) {
            return;
        }
    };
    result += gameId;
});

console.log(result);