export function toRna(dna: string): string {
  let dnaA = dna.split("");
  let hasil = "";
  dnaA.forEach((huruf:string) =>{
    switch(huruf){
      case 'G': 
        hasil=hasil+'C'
        break;
      case 'C':
        hasil=hasil+'G'
        break;
      case 'T':
        hasil=hasil+'A'
        break;
      case 'A':
        hasil=hasil+'U'
        break;
      default:
        throw new Error("Invalid input DNA.");
    }
  })
  return hasil
}