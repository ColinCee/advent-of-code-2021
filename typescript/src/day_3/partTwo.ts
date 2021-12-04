import path from 'path';
import { readLines } from '../readLines';

// const bitsList = [
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
// ];

const bitsList = readLines(path.join(__dirname, './input.txt'));

const countBitsAtIndex = (bitsList: string[], index: number) => {
  const bitCounts = {
    zero: 0,
    ones: 0,
  };
  bitsList.forEach((bits) => {
    if (bits[index] === '0') {
      bitCounts.zero++;
    } else if (bits[index] === '1') {
      bitCounts.ones++;
    }
  });

  return bitCounts;
};

const findMostCommonBitAtIndex = (
  bitsList: string[],
  index: number,
): number => {
  const bitCounts = countBitsAtIndex(bitsList, index);
  return bitCounts.ones >= bitCounts.zero ? 1 : 0;
};

const findLeastCommonBitAtIndex = (
  bitsList: string[],
  index: number,
): number => {
  const bitCounts = countBitsAtIndex(bitsList, index);
  return bitCounts.zero <= bitCounts.ones ? 0 : 1;
};

const findRating = (type: 'oxygen' | 'co2') => {
  let filteredBitList = [...bitsList];

  for (let i = 0; i < bitsList[0].length; i++) {
    if (filteredBitList.length === 1) {
      break;
    }
    const bitCriteria =
      type === 'oxygen'
        ? findMostCommonBitAtIndex(filteredBitList, i)
        : findLeastCommonBitAtIndex(filteredBitList, i);

    filteredBitList = filteredBitList.filter(
      (bit) => bit[i] === bitCriteria.toString(),
    );
    console.log('index', i, filteredBitList);
  }

  return filteredBitList[0];
};

// // multiply oxygenRating and co2Rating and convert to base 10 to get life support rating
const oxygenRating = parseInt(findRating('oxygen'), 2);
console.log('oxygenRating', oxygenRating);
const co2Rating = parseInt(findRating('co2'), 2);
console.log('co2Rating', co2Rating);
const lifeSupportRating = oxygenRating * co2Rating;
console.log('lifeSupportRating: ', lifeSupportRating);
