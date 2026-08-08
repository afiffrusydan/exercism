export function decodedResistorValue(colors: string[]): string {
  let nilai: number = 0;
  let i: number = 0;
  let satuan: string = "";
  colors.forEach((item) => {
    const itemId : number = COLORS.indexOf(item);
    switch(i){
      case 0:
        nilai = itemId;
        break;
      case 1:
        nilai = Number(""+nilai+itemId);
        break;
      case 2:
        const modulo: number = itemId%3;
        let nol: string = "";
        for (let j = 0; j < modulo; j++) {
          nol += "0";
        }
        switch((itemId-modulo)/3){
          case 0:
            if((nilai*(10**modulo))%1000===0 && nilai > 0){
              nilai = nilai*(10**modulo)/1000
              satuan = " kiloohms";
            } else {
              satuan = nol+ " ohms";
            }
            break;
          case 1:
            satuan = nol+ " kiloohms";
            break;
          case 2:
            satuan = nol+ " megaohms";
            break;
          case 3:
            satuan = nol+ " gigaohms";
            break;
          default:
            break;
        }
        break;
      default:
        break;
    }
    i++;
  })
  return nilai + satuan;
}
const COLORS = [
  'black',
  'brown',
  'red',
  'orange',
  'yellow',
  'green',
  'blue',
  'violet',
  'grey',
  'white',
];