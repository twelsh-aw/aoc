const fs = require('fs')
const path = require('path')

let inputPath = path.resolve(__dirname, "input.txt");
let input = fs.readFileSync(inputPath).toString().split("\n");

let p1 = puzzle1(input)
let p2 = puzzle2(input)
console.log(p1, p2)

function puzzle1(inputs) {
    let res = getGammaEpsilon(inputs)
    let gammaDec = parseInt(res.gamma, 2)
    let epsilonDec = parseInt(res.epsilon, 2)
    return gammaDec * epsilonDec
}

function puzzle2(inputs) {
    let oxy = getOxyCo2(inputs, 'gamma')
    let co2 = getOxyCo2(inputs, 'epsilon')
    return oxy * co2
}

function getOxyCo2(inputs, gammaOrEpsilon) {
    const numDigits = 12
    let inputsForLoop = [...inputs]
    for (let j = 0; j < numDigits; j++) {
        let nextInputs = []
        let res = getGammaEpsilon(inputsForLoop)
        for (let row of inputsForLoop) {
            if (row.charAt(j) === res[gammaOrEpsilon].charAt(j)) {
                nextInputs.push(row)
            }
        }

        if (nextInputs.length === 1) {
            return parseInt(nextInputs[0], 2)
        }

        inputsForLoop = [...nextInputs]
    }
}

function getGammaEpsilon(inputs) {
    const numDigits = 12
    let digitsFreqs = {}
    for (let row of inputs) {
        for (let j = 0; j < numDigits; j++) {
            let digitVal = row.charAt(j)
            if (!digitsFreqs[j]) {
                digitsFreqs[j] = {}
            }

            if (!digitsFreqs[j][digitVal]) {
                digitsFreqs[j][digitVal] = 0
            }
            digitsFreqs[j][digitVal]++
        }
    }

    let gamma = ''
    let epsilon = ''
    for (let digit of Object.keys(digitsFreqs)) {
        let freqs = digitsFreqs[digit]
        if (freqs['0'] > freqs['1']) {
            gamma += '0'
            epsilon += '1'
        } else {
            gamma += '1'
            epsilon += '0'
        }
    }

    return {gamma, epsilon}
}
