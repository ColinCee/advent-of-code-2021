import { readLines } from '../readLines';
import path from 'path';

// const lines = [
//   'forward 5',
//   'down 5',
//   'forward 8',
//   'up 3',
//   'down 8',
//   'forward 2',
// ];
const lines = readLines(path.join(__dirname, './input.txt'));

// for each line
// increase aim by the number when string starts with down
// decrease aim by the number when string starts with up
// increase position by the number when string starts with forward
// increase depth by aim multiplied by number when string starts with forward
let position = 0;
let depth = 0;
let aim = 0;

lines.forEach((line) => {
  const [direction, number] = line.split(' ');
  const value = Number(number);
  if (direction === 'down') {
    aim += value;
  } else if (direction === 'up') {
    aim -= value;
  } else if (direction === 'forward') {
    position += value;
    depth += aim * value;
  }
});

console.log('position: ', position);
console.log('depth: ', depth);
console.log('final', depth * position);
