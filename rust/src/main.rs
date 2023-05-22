fn main() {
    let out = get_input()
        .lines()
        .map(parse_line)
        .fold(Vec {x: 0, y: 0}, |mut acc, vec| {
            acc.x += vec.x;
            acc.y += vec.y;
            
            return acc;
    });

    print!("x: {}, y: {}", out.x, out.y);
}


fn get_input() -> &'static str {
    return "forward 5
down 5
forward 8
up 3
down 8
forward 2";
}

struct Vec {
    x: i32,
    y: i32,
}

fn parse_line(line: &str) -> Vec {
    let (dir, amount) = line.split_once(" ").expect("whitespace or go boom");
    let amount = str::parse::<i32>(amount).expect("number or go boom");

    if dir == "forward" {
        return Vec {x: amount, y: 0}
    } else if dir == "up" {
        return Vec {x: 0, y: -amount}
    }
    return Vec {x: 0, y: amount}
}
