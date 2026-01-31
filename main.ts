const isEvenOrOdd = (n: number) => {
  let num = n % 2;

  if (num != 0) {
    return false;
  }

  return true;
};

console.log(isEvenOrOdd(10));
console.log(isEvenOrOdd(15));
console.log(isEvenOrOdd(40));
