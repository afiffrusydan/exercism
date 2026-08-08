export function decodedValue(colors: string[]) {
  let no1 = COLORS.indexOf(colors[0])
  let no2 = COLORS.indexOf(colors[1])
  return Number(""+no1+no2)
}
const COLORS =[
  "black",
  "brown",
  "red",
  "orange",
  "yellow",
  "green",
  "blue",
  "violet",
  "grey",
  "white",
]