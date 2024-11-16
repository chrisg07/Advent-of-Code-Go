## Go Advent of Code

### Structure

Each year expects a top level directory, e.g. `2021`

Each day's solution for a year expects a child directory, e.g. `2021/Day9`

Each solution's directory requires an `input` directory with two files, `example.txt` and `complete.txt`. 

`.txt` files are ignored due to AoC's preference that puzzle inputs not be shared publicly.

### Running Tests

To execute the tests for a particular day's solution:
- Add the required `inputs` (`example.txt` and `complete.txt`)
- Navigate to that day's directory, e.g. `2021/Day9`
- Run `go test`

## Solutions with visualizations
d
### 2021 

#### Day 9

Display basins: `go test -timeout 30s -run ^TestDay9PartB2021Complete$ github.com/chrisg07/Advent-of-Code-Go/2021/Day9 -v`