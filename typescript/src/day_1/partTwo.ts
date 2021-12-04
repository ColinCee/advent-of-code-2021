import lines from './data.json';

let count = 0;
for (let i = 1; i < lines.length; i++) {
  const previousSum = lines[i - 1] + lines[i] + lines[i + 1];
  const currentSum = lines[i] + lines[i + 1] + lines[i + 2];

  if (currentSum > previousSum) {
    count++;
  }
}

console.log('count', count);
