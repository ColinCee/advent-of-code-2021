// count the number of times each bit appears in each position
interface BitCount {
  zeros: number;
  ones: number;
}

export const countBitsInEachPosition = (bits: string[]): BitCount[] => {
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

// get most common bits from bitCounts
export const getMostCommonBits = (
  bitCounts: BitCount[],
  equalBitResolution: 1 | 0 = 1,
): string[] => {
  return bitCounts.map((bitCount) => {
    if (bitCount.zeros > bitCount.ones) {
      return '0';
    } else if (bitCount.zeros < bitCount.ones) {
      return '1';
    } else {
      return equalBitResolution ? '1' : '0';
    }
  });
};

// get least common bits from bitCounts
export const getLeastCommonBits = (
  bitCounts: BitCount[],
  equalBitResolution: 1 | 0 = 1,
): string[] => {
  return bitCounts.map((bitCount) => {
    if (bitCount.zeros < bitCount.ones) {
      return '0';
    } else if (bitCount.zeros > bitCount.ones) {
      return '1';
    } else {
      return equalBitResolution ? '1' : '0';
    }
  });
};
