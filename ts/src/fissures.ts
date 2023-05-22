/* Input
0,9 -> 5,9
8,0 -> 0,8
9,4 -> 3,4
2,2 -> 2,1
7,0 -> 7,4
6,4 -> 2,0
0,9 -> 2,9
3,4 -> 1,4
0,0 -> 8,8
5,5 -> 8,2
*/

function getInput(): string {
  return `0,9 -> 5,9
8,0 -> 0,8
9,4 -> 3,4
2,2 -> 2,1
7,0 -> 7,4
6,4 -> 2,0
0,9 -> 2,9
3,4 -> 1,4
0,0 -> 8,8
5,5 -> 8,2`;
}

type Vec = {
  x: number;
  y: number;
}

type Line = {
  start: Vec,
  end: Vec,
}

function isHorOrVert(line: Line) {
  return line.start.x === line.end.x || line.start.y === line.end.y;
}

function parseVec(v: string) {
  const [x, y] = v.split(",");
  
  return {
    x: +x,
    y: +y,
  }
}

function parseLine(l: string) {
  const [start, end] = l.split(" -> ");

  return {
    start: parseVec(start),
    end: parseVec(end)
  }
};

const lines = getInput()
  .split("\n")
  .map(x => parseLine(x))
  .filter(isHorOrVert)

console.log(lines)
