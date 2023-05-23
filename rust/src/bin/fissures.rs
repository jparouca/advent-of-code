use std::str::FromStr;
use anyhow::{Result, anyhow};

fn get_input() -> &'static str {
    return "0,9 -> 5,9
8,0 -> 0,8
9,4 -> 3,4
2,2 -> 2,1
7,0 -> 7,4
6,4 -> 2,0
0,9 -> 2,9
3,4 -> 1,4
0,0 -> 8,8
5,5 -> 8,2"
}
#[derive(Debug)]
struct Vect {
    x: i32,
    y: i32,
}

#[derive(Debug)]
struct Line {
    start: Vect,
    end: Vect,
}

impl FromStr for Vect {
    type Err = anyhow::Error;
    fn from_str(s: &str) -> Result<Self>{
        let result = s.split_once(",");
        if result.is_none() {

            return Err(anyhow!("no string floats?? boom"));
        }
        let (x, y) = result.unwrap();
        let x: i32 = str::parse(x)?;
        let y: i32 = str::parse(y)?;

         return Ok(Vect {
            x, y
        });
    }
}

impl FromStr for Line{
    type Err = anyhow::Error;
    fn from_str(s: &str) -> Result<Self>{
        let result = s.split_once(" -> ");
        if result.is_none() {

            return Err(anyhow!("no arrows? you guess. boom"));
        }
        let (start, end) = result.unwrap();
        let start: Vect = str::parse(start)?;
        let end: Vect = str::parse(end)?;

        return Ok(Line {
            start, end
        })
    }
}

impl Line {
    fn is_hor_or_ver(&self) -> bool {
        return self.start.x == self.end.x || self.start.y == self.end.y;
    }
}


fn main() {
    let lines = get_input()
        .lines()
        .flat_map(|x| str::parse(x))
        .filter(|x: &Line| x.is_hor_or_ver())
        .collect::<Vec<Line>>();

    println!("{:?}", lines);

}
