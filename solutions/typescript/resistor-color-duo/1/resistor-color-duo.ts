export function decodedValue(arr:string[]):number {
  const colorMap:Record<string, number> = {
  black: 0,
  brown: 1,
  red: 2,
  orange: 3,
  yellow: 4,
  green: 5,
  blue: 6,
  violet: 7,
  grey: 8,
  white: 9,
}
  let num:number = Number(`${colorMap[arr[0]]}${colorMap[arr[1]]}`)
  
  
  return num
}
