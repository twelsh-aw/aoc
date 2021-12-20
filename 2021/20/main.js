const fs = require('fs')
const path = require('path')

let inputPath = path.resolve(__dirname, "input.txt");
let input = fs.readFileSync(inputPath).toString().split('\n')
let parsed = parseInput(input)

let p1 = puzzle1()
let p2 = puzzle2()
console.log(p1, p2)

function puzzle1() {
    let image = parsed.image
    for (let i = 0; i < 2; i++) {
        image = enhance(image, i)
    }

    return countLit(image)
}

function puzzle2() {
    let image = parsed.image
    for (let i = 0; i < 50; i++) {
        image = enhance(image, i)
    }

    return countLit(image)
}

function countLit(image) {
    let lit = 0
    for (let row = 0; row < image.length; row++) {
        for (let col = 0; col < image[0].length; col++) {
            if (image[row][col] === '#') {
                lit++
            }
        }
    }

    return lit
}

function enhance(image, step) {
    let expanded = expandImage(image, step)
    let newImage = []
    for (let row = 0; row < expanded.length; row++) {
        let rowValues = []
        for (let col = 0; col < expanded[0].length; col++) {
            let binaryValue = getBinaryValue(expanded, row, col, step)
            let enhancedValue = parsed.alg[binaryValue]
            rowValues.push(enhancedValue)
        }
        newImage.push(rowValues)
    }

    return newImage
}

function getBinaryValue(expanded, row, col, step) {
    let binaryString = ''
    if (expanded[row - 1]) {
        binaryString += convertValue(expanded[row - 1][col - 1], step)
        binaryString += convertValue(expanded[row - 1][col], step)
        binaryString += convertValue(expanded[row - 1][col + 1], step)
    } else {
        binaryString += convertValue(undefined, step)
        binaryString += convertValue(undefined, step)
        binaryString += convertValue(undefined, step)
    }

    binaryString += convertValue(expanded[row][col - 1], step)
    binaryString += convertValue(expanded[row][col], step)
    binaryString += convertValue(expanded[row][col + 1], step)

    if (expanded[row + 1]) {
        binaryString += convertValue(expanded[row + 1][col - 1], step)
        binaryString += convertValue(expanded[row + 1][col], step)
        binaryString += convertValue(expanded[row + 1][col + 1], step)
    } else {
        binaryString += convertValue(undefined, step)
        binaryString += convertValue(undefined, step)
        binaryString += convertValue(undefined, step)
    }

    if (binaryString.length !== 9) {
        throw binaryString
    }

    return parseInt(binaryString, 2)
}

function convertValue(value, step) {
    if (value === '.') {
        return '0'
    } else if (value === '#') {
        return '1'
    } else if (value === undefined) {
        return (step % 2).toString()
    } else {
        throw value
    }
}

// pad with extra 2 rows/cols of . since infiinite
function expandImage(image, step) {
    let newImage = []
    let cols = image[0].length
    let defaultVal = (step % 2) ? '#' : '.'
    let emptyRows = new Array(cols + 2).fill(defaultVal)
    newImage.push([...emptyRows])
    for (let row of image) {
        newImage.push([defaultVal, ...row, defaultVal])
    }
    newImage.push([...emptyRows])
    return newImage
}

function parseInput(input) {
    let alg = input[0].split('')
    let image = []
    for (let row of input.slice(2)) {
        image.push(row.split(''))
    }

    return {alg, image}
}
