import { readLines } from '../readLines';
import path from 'path';

// const lines = [
//     'forward 5',
//     'down 5',
//     'forward 8',
//     'up 3',
//     'down 8',
//     'forward 2',
//   ];
const lines = readLines(path.join(__dirname, './input.txt'));

// increase depth by the number when string starts with down
// decrease depth by the number when string starts with up
// increase position by the number when string starts with forward
let position = 0;
let depth = 0;

lines.forEach((line) => {
  const [direction, number] = line.split(' ');
  switch (direction) {
    case 'up':
      depth -= Number(number);
      break;
    case 'down':
      depth += Number(number);
      break;
    case 'forward':
      position += Number(number);
      break;
    default:
      console.log('error');
  }
});

console.log('position: ', position);
console.log('depth: ', depth);
console.log('final', depth * position);
