import { paginate } from './paginate';

describe('paginate', () => {
  const pageSize = 5;

  it('Test1', () => {
    const input = [1, 2, 3, 4, 5, 6, 7, 8, 9, 10];
    expect(paginate(input, pageSize, 0)).toEqual([1, 2, 3, 4, 5]);
    expect(paginate(input, pageSize, 1)).toEqual([6, 7, 8, 9, 10]);
  });

  it('Test2', () => {
    const input = [1, 2, 3, 4];
    expect(paginate(input, pageSize, 0)).toEqual([1, 2, 3, 4]);
    expect(paginate(input, pageSize, 1)).toEqual([]);
  });

  it('Test3', () => {
    const input = [1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11];
    expect(paginate(input, pageSize, 0)).toEqual([1, 2, 3, 4, 5]);
    expect(paginate(input, pageSize, 1)).toEqual([6, 7, 8, 9, 10]);
    expect(paginate(input, pageSize, 2)).toEqual([11]);
  });

  it('Test4', () => {
    const input = [1, 2, 3, 4, 5, 6];
    expect(paginate(input, pageSize, 0)).toEqual([1, 2, 3, 4, 5]);
    expect(paginate(input, pageSize, 1)).toEqual([6]);
  });

  it('Test5', () => {
    const input = [1, 2, 3, 4, 5];
    expect(paginate(input, pageSize, 0)).toEqual([1, 2, 3, 4, 5]);
    expect(paginate(input, pageSize, 1)).toEqual([]);
  });
});
