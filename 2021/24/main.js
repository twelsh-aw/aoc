const fs = require('fs')
const path = require('path')

let inputPath = path.resolve(__dirname, "input.txt");
let input = fs.readFileSync(inputPath).toString().split('\n')
let instructions = parseInput(input)
let debug = false
let loopDown = [9, 8, 7, 6, 5, 4, 3, 2, 1]
let loopUp = [1, 2, 3, 4, 5, 6, 7, 8, 9]
const maxNumber = 79997391969649 // from part 1
const minNumber = 16931171414113 // from part 2
// warning slow to run with these vals ~ 10 mins
// const maxNumber = 99999999999999 // if re-running part 1...
// const minNumber = 11111111111111 // if re-running part 2...

let p1 = puzzle1(loopDown)
let p2 = puzzle1(loopUp)
console.log(p1, p2)

function puzzle1(loopOrder) {
    let inputIndexes = instructions.filter(i => i.mode === "inp").map(i => instructions.indexOf(i));
    let dimsSeen = {}
    let startDims = {x: 0, y: 0, z: 0, w: 0}
    dimsSeen[getKey(startDims)] = 0
    for (let i = 0; i < 14; i++) {
        let instructionsToUse = getInstructionsToUse(instructions, inputIndexes, i)
        dimsSeen = getDimsSeen(dimsSeen, instructionsToUse, loopOrder)
        debug && console.log(i + 1, Object.keys(dimsSeen).length)
        if (Object.keys(dimsSeen).length > Math.pow(9, 14 - i + 1)) {
            // object getting too big, re-run on each piece we have so far
            debug && console.log('break')
            break
        }
    }

    let inverted = {}
    for (let key of Object.keys(dimsSeen)) {
        inverted[dimsSeen[key]] = key
    }

    let numTried = 0
    let keysOrdered = Object.keys(inverted).sort((a, b) => {
        if (loopOrder[0] === 1) {
            return a - b
        } else if (loopOrder[0] === 9) {
            return b -a
        } else {
            throw 'how to sort?'
        }
    });

    // can't fit everything into single object
    // run rest of algorithm 1 sequence at a time
    // could be smarter here and put in proper "batch" logic
    for (let key of keysOrdered) {
        let dimsSeen = {}
        dimsSeen[inverted[key]] = 0
        for (let i = key.length; i < 14; i++) {
            let instructionsToUse = getInstructionsToUse(instructions, inputIndexes, i)
            dimsSeen = getDimsSeen(dimsSeen, instructionsToUse, loopOrder)
        }

        let valid = Object.keys(dimsSeen).filter(k => parseKey(k).z === 0)
        if (valid.length) {
            let digits = Math.max(...valid.map(v => dimsSeen[v])).toString()
            return parseInt(key + digits)
        }
        debug && numTried++
        if (debug && numTried % 10000 === 0) {
            console.log(numTried, "out of", Object.keys(inverted).length, "(looking at", key, ")")
        }
    }

    throw 'did not find valid digits'
}

function getKey(dims) {
    return `${dims.x}_${dims.y}_${dims.z}_${dims.w}`
}

function parseKey(key) {
    let arr = key.split('_').map(Number)
    return {x: arr[0], y: arr[1], z: arr[2], w: arr[3]}
}

function getDigitUntilBound(numDigits, bound) {
    return parseInt(bound.toString().split('').slice(0, numDigits + 1).join(''))
}

function getNumDigits(dimsSeen) {
    let x = Object.values(dimsSeen)[0]
    if (x === 0) {
        return 0
    } else {
        return x.toString().length
    }
}

function getDimsSeen(dimsSeen, instructionsToUse, loopOrder) {
    let newDimsSeen = Object.create({})
    let numDigits = getNumDigits(dimsSeen)
    let maxNumForLoop = getDigitUntilBound(numDigits, maxNumber)
    let minNumForLoop = getDigitUntilBound(numDigits, minNumber)
    for (let dimsKey of Object.keys(dimsSeen)) {
        let num = dimsSeen[dimsKey];
        let dims = parseKey(dimsKey)
        // down for max i.e puzzle1, up for min = puzzle2
        for (let digit of loopOrder) {
            let nextNum = (num*10) + digit
            if (nextNum > maxNumForLoop || nextNum < minNumForLoop) {
                continue
            }

            let newDims = getNextValues({...dims}, instructionsToUse, digit)
            if (isNaN(newDims.z)) {
                continue
            }

            let newKey = getKey(newDims)
            if (newDimsSeen[newKey] === undefined) {
                newDimsSeen[newKey] = nextNum
            }
        }
    }

    return newDimsSeen
}

function getInstructionsToUse(instructions, inputIndexes, inputIndex) {
    let instructionsToUse = []
    let curInput = inputIndexes[inputIndex]
    let nextInput = inputIndexes[inputIndex + 1]
    if (nextInput === undefined) {
        nextInput = instructions.length
    }
    for (let i = curInput; i < nextInput; i++) {
        instructionsToUse.push(instructions[i])
    }

    return instructionsToUse
}

function digitsToNumber(digits) {
    return parseInt(digits.join(''))
}

function getNextValues(dims, instructions, digit) {
    for (let instruction of instructions) {
        switch (instruction.mode) {
            case "inp":
                dims[instruction.variable] = digit
                break
            case "add":
                dims[instruction.variable] += getVal(dims, instruction.val)
                break
            case "mul":
                dims[instruction.variable] *= getVal(dims, instruction.val)
                break
            case "div":
                if (getVal(dims, instruction.val) === 0) {
                    return {x: 0, y: 0, z: NaN, w: 0}
                }
                dims[instruction.variable] = Math.floor(dims[instruction.variable] / getVal(dims, instruction.val))
                break
            case "mod":
                if (dims[instruction.variable] < 0 || getVal(dims, instruction.val) <= 0) {
                    return {x: 0, y: 0, z: NaN, w: 0}
                }
                dims[instruction.variable] = dims[instruction.variable] % getVal(dims, instruction.val)
                break
            case "eql":
                if (dims[instruction.variable] === getVal(dims, instruction.val)) {
                    dims[instruction.variable] = 1
                } else {
                    dims[instruction.variable] = 0
                }
                break
            default:
                throw instruction.mode
        }
    }

    return dims
}

function getVal(dims, val) {
    if (dims[val] !== undefined) {
        return dims[val]
    } else {
        return parseInt(val)
    }
}

function parseInput(input) {
    let instructions = []
    for (let row of input) {
        let split = row.split(' ')
        let mode = split[0]
        let variable = split[1]
        let val = split[2]
        instructions.push({mode, variable, val})
    }

    return instructions
}
