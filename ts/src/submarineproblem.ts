/* https://adventofcode.com/2021/day/2
problem on /submarine.md
*/

function getInput(): string {
  return `forward 5
down 5
forward 8
up 3
down 8
forward 2`;
}

function parseLine(line: string): [number, number] {
  const [direction, a] = line.split(" ");
  const amount = +a;

  if (direction === "forward") {
    return [amount, 0];
  } else if (direction === "up") {
    return [0, -amount];
  }

  return [0, amount];

}

const out = getInput()
  .split("\n")
  .map(x => parseLine(x))
  .reduce((acc, amount) => {
    acc[0] += amount[0];
    acc[1] += amount[1];
    return acc;
}, [0, 0]);

console.log(out, out[0] * out[1]);