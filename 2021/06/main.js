const fs = require('fs')
const path = require('path')

let inputPath = path.resolve(__dirname, "input.txt");
let initialFish = fs.readFileSync(inputPath).toString().split(",");

let p1 = puzzle1(80)
let p2 = puzzle2()
console.log(p1, p2)

function puzzle1(days) {
    let fishDict = parseInput(initialFish)
    for (let day = 0; day < days; day++) {
        let nextDict = {}
        for (let num of Object.keys(fishDict).sort()) {
            if (num === '0') {
                nextDict[6] = fishDict[0]
                nextDict[8] = fishDict[0]
            } else if (num === '7' && nextDict[6]) {
                nextDict[6] += fishDict[7]
            } else {
                nextDict[num - 1] = fishDict[num]
            }
        }

        fishDict = {...nextDict}
    }

    return Object.values(fishDict).reduce((a, b) => a + b)
}

function puzzle2() {
    return puzzle1(256)
}

function parseInput(initialFish) {
    let grouped = {}
    for (let fish of initialFish) {
        if (!grouped[fish]) {
            grouped[fish] = 0
        }

        grouped[fish]++
    }

    return grouped
}
