const fs = require('fs')
const path = require('path')

let inputPath = path.resolve(__dirname, "input.txt");
let input = fs.readFileSync(inputPath).toString();
let parsed = parseInput(input)

let p1 = puzzle1()
let p2 = puzzle2()
console.log(p1, p2)

// dumb logic, but not changing
function puzzle1() {
    let xVMin = Math.ceil((Math.sqrt(1 + 8 * parsed.xMin) - 1) / 2)
    let x = xVMin
    let y = 0
    let numMisses = 0
    let maxYs = []
    while (true) {
        let xPos = 0
        let yPos = 0
        let step = 1
        let yVals = []
        let inTarget = false
        while (true) {
            xPos += Math.max(0, x - (step - 1))
            yPos += y - (step - 1)
            yVals.push(yPos)

            if (xPos >= parsed.xMin && xPos <= parsed.xMax && yPos >= parsed.yMin && yPos <= parsed.yMax) {
                inTarget = true
                break
            }

            if (xPos < parsed.xMin && yPos < parsed.yMin) {
                break
            }

            if (xPos > parsed.xMax || yPos < parsed.yMin) {
                break
            }

            step++
        }

        if (inTarget) {
            maxYs.push(Math.max(...yVals))
        } else {
            numMisses++
        }

        if (numMisses > 100) {
            return Math.max(...maxYs)
        }

        y++
    }
}

function puzzle2() {
    let x = 0
    let velocities = []
    while (true) {
        if (x > parsed.xMax) {
            break
        }

        for (let y = 0; true; y++) {
            let obj = trackTrajectory(x, y)
            if (obj.inTarget) {
                velocities.push([x, y])
            } else if (obj.breakX) {
                break
            }
        }

        for (let y = -1; true; y--) {
            let obj = trackTrajectory(x, y)
            if (obj.inTarget) {
                velocities.push([x, y])
            } else if (y < parsed.yMin) {
                break
            }
        }

        x++
    }

    //console.table(velocities)
    return velocities.length
}

function trackTrajectory(xv, yv) {
    let step = 0
    let x = 0
    let y = 0
    let inTarget = false
    let breakX = false
    while (true) {
        // console.log(x, y, step)
        let xd = Math.max(0, xv - step)
        if (xd === 0 && x < parsed.xMin) {
            breakX = true
            return {inTarget, breakX}
        }

        let yd = yv - step
        x += xd
        y += yd

        if (xd === 0 && y < parsed.yMin && step > (parsed.xMax - parsed.yMin)) {
            breakX = true
            return {inTarget, breakX}
        }

        if (y < parsed.yMin) {
            return {inTarget, breakX}
        }

        if (x >= parsed.xMin && x <= parsed.xMax && y >= parsed.yMin && y <= parsed.yMax) {
            inTarget = true
            return {inTarget, breakX}
        }

        step++
    }
}

function parseInput(input) {
    let data = input.split('target area: ')[1]
    let split = data.split(', ')
    let x = split[0]
    let xMin = parseInt(x.split('..')[0].split('x=')[1])
    let xMax = parseInt(x.split('..')[1])
    let y = split[1]
    let yMin = parseInt(y.split('..')[0].split('y=')[1])
    let yMax = parseInt(y.split('..')[1])
    return {xMin, xMax, yMin, yMax}
}

