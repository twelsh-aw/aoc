const fs = require('fs')
const path = require('path')

let inputPath = path.resolve(__dirname, "input.txt");
let input = fs.readFileSync(inputPath).toString().split("\n");

let p1 = puzzle1(input)
let p2 = puzzle2(input)
console.log(p1, p2)

function puzzle1(inputs) {
    let numIncreases = 0
    for (let i = 0; i < inputs.length - 1; i++) {
        if (parseFloat(inputs[i+1]) - parseFloat(inputs[i]) > 0) {
            numIncreases++
        }
    }

    return numIncreases
}

function puzzle2(inputs) {
    let sums = []
    for (let i = 0; i < inputs.length - 2; i++) {
        let sum = parseFloat(inputs[i+2]) + parseFloat(inputs[i+1]) + parseFloat(inputs[i])
        sums.push(sum)
    }

    return puzzle1(sums)
}