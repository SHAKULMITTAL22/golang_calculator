


const {add,divide}=require('./math');
 test('adds 2+3 to equal 5',()=>{
    const actual =add(2,3);
    const expected=5;
    expect(actual).toBe(expected);
 
 });
 

 


test('divides 10 by 2 to equal 5', () => {
  const actual = divide(10, 2);
  const expected = 5;
  expect(actual).toBe(expected);
});

test('throws error when dividing by zero', () => {
  expect(() => divide(10, 0)).toThrow("Cannot divide by zero");
});