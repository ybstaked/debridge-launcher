/**
 * Function for getting part of documents
 * @param array
 * @param pageSize
 * @param pageNumber begin from 0
 */
export function paginate(array: any, pageSize: number, pageNumber: number) {
  const skip = pageNumber * pageSize;
  const end = Math.min((pageNumber + 1) * pageSize, array.length);
  return array.slice(skip, end);
}
