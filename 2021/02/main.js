const fs = require('fs')
const path = require('path')

let inputPath = path.resolve(__dirname, "input.txt");
let input = fs.readFileSync(inputPath).toString().split("\n");

let p1 = puzzle1()
let p2 = puzzle2()
console.log(p1)
console.log(p2)

function puzzle1() {
    let pos = [0,0]
    for (let line of input) {
        let splitLine = line.split(" ")
        let direction = splitLine[0]
        let value = parseFloat(splitLine[1])
        if (direction === 'forward') {
            pos[0] += value
        } else if (direction === 'down') {
            pos[1] += value
        } else if (direction === 'up') {
            pos[1] -= value
        } else {
            throw 'unknown direction: ' + direction
        }
    }

    return pos[0] * pos[1]
}

function puzzle2() {
    let pos = [0,0,0] // horiz, depth, aim
    for (let line of input) {
        let splitLine = line.split(" ")
        let direction = splitLine[0]
        let value = parseFloat(splitLine[1])
        if (direction === 'forward') {
            pos[0] += value
            pos[1] += pos[2] * value
        } else if (direction === 'down') {
            pos[2] += value
        } else if (direction === 'up') {
            pos[2] -= value
        } else {
            throw 'unknown direction: ' + direction
        }
    }

    return pos[0] * pos[1]
}
