const fs = require('fs')
const path = require('path')

let inputPath = path.resolve(__dirname, "input.txt");
let input = fs.readFileSync(inputPath).toString().split(",").map(Number);

let p1 = puzzle1()
let p2 = puzzle2()
console.log(p1, p2)

function puzzle1() {
    let minFuel = Math.pow(2, 32)
    for (let i = 0; i < Math.max(...input); i++) {
        let fuel = input.map(p => Math.abs(p -i)).reduce((a, b) => a + b)
        if (fuel < minFuel) {
            minFuel = fuel
        }
    }

    return minFuel
}

function puzzle2() {
    let minFuel = Math.pow(2, 32)
    for (let i = 0; i < Math.max(...input); i++) {
        let fuel = input.map(p => {
            let d = Math.abs(p - i)
            return d * (d+1) / 2
        }).reduce((a, b) => a + b)
        if (fuel < minFuel) {
            minFuel = fuel
        }
    }

    return minFuel
}

