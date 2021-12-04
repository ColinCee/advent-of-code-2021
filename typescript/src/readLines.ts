// create a function that reads a file and returns an array of lines
import fs from 'fs';

export const readLines = (path: string): string[] => {
  return fs
    .readFileSync(path, 'utf-8')
    .split('\n')
    .filter((line) => line !== '\n');
};
