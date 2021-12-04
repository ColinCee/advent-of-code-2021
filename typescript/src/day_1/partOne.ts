import lines from './data.json';

let count = 0;
for (let i = 1; i < lines.length; i++) {
  if (lines[i] > lines[i - 1]) {
    count++;
  }
}

console.log('count', count);
