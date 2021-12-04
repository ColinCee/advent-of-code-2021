import path from 'path';
import { readLines } from '../readLines';

// const bits = [
//   '00100',
//   '11110',
//   '10110',
//   '10111',
//   '10101',
//   '01111',
//   '00111',
//   '11100',
//   '10000',
//   '11001',
//   '00010',
//   '01010',
const bits = readLines(path.join(__dirname, './input.txt'));

// count the number of times each bit appears in each position
interface BitCount {
  zeros: number;
  ones: number;
}

const countBits = (bits: string[]): BitCount[] => {
  const bitCounts: BitCount[] = bits[0]
    .split('')
    .map(() => ({ zeros: 0, ones: 0 }));

  bits.forEach((number) => {
    number.split('').forEach((bit, index) => {
      if (bit === '0') {
        bitCounts[index].zeros++;
      } else {
        bitCounts[index].ones++;
      }
    });
  });

  return bitCounts;
};

const bitCounts = countBits(bits);
// get most common bits from bitCounts
const gammaRate = bitCounts.map((bitCount) => {
  if (bitCount.ones > bitCount.zeros) {
    return '1';
  } else {
    return '0';
  }
});

console.log(gammaRate.join(''));

// get least common bits from bitCounts
const epsilonRate = bitCounts.map((bitCount) => {
  if (bitCount.ones < bitCount.zeros) {
    return '1';
  } else {
    return '0';
  }
});

console.log(epsilonRate.join(''));

// powerConsumption is the base 10 value of gammaRate multiplied by base 10 value of epsilonRate
const powerConsumption =
  parseInt(gammaRate.join(''), 2) * parseInt(epsilonRate.join(''), 2);

console.log(powerConsumption);
