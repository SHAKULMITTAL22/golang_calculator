


const add=require('./math');
 test('adds 2+3 to equal 5',()=>{
    const actual =add(2,3);
    const expected=5;
    expect(actual).toBe(expected);
 
 });
 console.log('All tests passed!');