double calculeazaTotal({pret, cantitate = 1}) {
  return pret * cantitate;
}

String afiseazaAdresa([String? adresa]) {
  adresa ??= "necunoscut";
  return "Adresa: $adresa"; 
}


void main() {
 print(calculeazaTotal(pret: 10.5));
 print(calculeazaTotal(pret: 10.5, cantitate: 3));
 print(afiseazaAdresa("Strada Florilor 5"));
 print(afiseazaAdresa(null));
}

