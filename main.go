package main

import (
	"log"
	"net/http"
	"os"
	"encoding/json"

	"github.com/gorilla/mux"
)

// ---------------------------------------------------------------------------------------
// Data
// ---------------------------------------------------------------------------------------

type Kpital struct {
	Country string `json:"Country"`
	City    string `json:"City"`
}

type Kpitals []Kpital

// ---------------------------------------------------------------------------------------
// Endpoints
// ---------------------------------------------------------------------------------------

// "/kpitals/all"
func kpitals(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)

	json.NewEncoder(w).Encode(allKpitals())
}


// "/kpitals/city/{country}"
func CityHandler(w http.ResponseWriter, r *http.Request) {
	vars    := mux.Vars(r)
	country := vars["country"]

	w.WriteHeader(http.StatusOK)

	for _, v := range allKpitals() {
    	if v.Country == country {
        	json.NewEncoder(w).Encode(v.City)
    	}
	}
}

// "/kpitals/country/{city}"
func CountryHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	city := vars["city"]

	w.WriteHeader(http.StatusOK)
	
	for _, v := range allKpitals() {
    	if v.City == city {
        	json.NewEncoder(w).Encode(v.Country)
    	}
	}
}

func homepage(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(allKpitals())	
}


// ---------------------------------------------------------------------------------------
// Rooting
// ---------------------------------------------------------------------------------------

func handleRequests() {
	r := mux.NewRouter().StrictSlash(true)

	r.HandleFunc("/", homepage).Methods("GET")
	r.HandleFunc("/kpitals/all", kpitals).Methods("GET")
	r.HandleFunc("/kpitals/country/{city}", CountryHandler).Methods("GET")
	r.HandleFunc("/kpitals/city/{country}", CityHandler).Methods("GET")
	http.Handle("/", r)
}

// ---------------------------------------------------------------------------------------
// -
// ---------------------------------------------------------------------------------------

func main() {
	log.Message("Hello")
	port := os.Getenv("PORT")

	if port == "" {
		log.Fatal("$PORT must be set")
	}


	handleRequests()
}

// ---------------------------------------------------------------------------------------
// LDB
// ---------------------------------------------------------------------------------------

func allKpitals() Kpitals {
	return Kpitals{
		Kpital{City: "Abou Dabi", Country: "Émirats arabes unis"},
		Kpital{City: "Abuja", Country: "Nigeria"},
		Kpital{City: "Accra", Country: "Ghana"},
		Kpital{City: "Achgabat", Country: "Turkménistan"},
		Kpital{City: "Addis-Abeba", Country: "Éthiopie"},
		Kpital{City: "Alger", Country: "Algérie"},
		Kpital{City: "Alofi", Country: "Niue"},
		Kpital{City: "Amman", Country: "Jordanie"},
		Kpital{City: "Amsterdam", Country: "Pays-Bas"},
		Kpital{City: "Andorre-la-Vieille", Country: "Andorre"},
		Kpital{City: "Ankara", Country: "Turquie"},
		Kpital{City: "Antananarivo", Country: "Madagascar"},
		Kpital{City: "Apia", Country: "Samoa"},
		Kpital{City: "Asmara", Country: "Érythrée"},
		Kpital{City: "Asuncion", Country: "Paraguay"},
		Kpital{City: "Athènes", Country: "Grèce"},
		Kpital{City: "Avarua", Country: "Îles Cook"},
		Kpital{City: "Bagdad", Country: "Irak"},
		Kpital{City: "Bakou", Country: "Azerbaïdjan"},
		Kpital{City: "Bamako", Country: "Mali"},
		Kpital{City: "Bandar Seri Begawan", Country: "Brunei"},
		Kpital{City: "Bangkok", Country: "Thaïlande"},
		Kpital{City: "Bangui", Country: "République centrafricaine"},
		Kpital{City: "Banjul", Country: "Gambie"},
		Kpital{City: "Basseterre", Country: "Saint-Christophe-et-Niévès"},
		Kpital{City: "Belgrade", Country: "Serbie"},
		Kpital{City: "Belmopan", Country: "Belize"},
		Kpital{City: "Berlin", Country: "Allemagne"},
		Kpital{City: "Berne (capitale de facto)3", Country: "Suisse"},
		Kpital{City: "Beyrouth", Country: "Liban"},
		Kpital{City: "Bichkek", Country: "Kirghizistan"},
		Kpital{City: "Bissau", Country: "Guinée-Bissau"},
		Kpital{City: "Bloemfontein (capitale judiciaire)4", Country: "Afrique du Sud"},
		Kpital{City: "Bogota", Country: "Colombie"},
		Kpital{City: "Brasilia", Country: "Brésil"},
		Kpital{City: "Bratislava", Country: "Slovaquie"},
		Kpital{City: "Brazzaville", Country: "République du Congo"},
		Kpital{City: "Bridgetown", Country: "Barbade"},
		Kpital{City: "Bruxelles", Country: "Belgique"},
		Kpital{City: "Bucarest", Country: "Roumanie"},
		Kpital{City: "Budapest", Country: "Hongrie"},
		Kpital{City: "Buenos Aires", Country: "Argentine"},
		Kpital{City: "Bujumbura (capitale économique)", Country: "Burundi"},
		Kpital{City: "Le Caire", Country: "Égypte"},
		Kpital{City: "Canberra", Country: "Australie"},
		Kpital{City: "Le Cap (capitale législative)4", Country: "Afrique du Sud"},
		Kpital{City: "Caracas", Country: "Venezuela"},
		Kpital{City: "Castries", Country: "Sainte-Lucie"},
		Kpital{City: "Chisinau", Country: "Moldavie"},
		Kpital{City: "Sri Jayawardenapura Kotte (administrative et législative)5", Country: "Sri Lanka"},
		Kpital{City: "Conakry", Country: "Guinée"},
		Kpital{City: "Copenhague", Country: "Danemark"},
		Kpital{City: "Dakar", Country: "Sénégal"},
		Kpital{City: "Damas", Country: "Syrie"},
		Kpital{City: "Delap-Uliga-Darrit (Majuro)6", Country: "Îles Marshall"},
		Kpital{City: "Dacca", Country: "Bangladesh"},
		Kpital{City: "Dili", Country: "Timor oriental"},
		Kpital{City: "Djibouti", Country: "Djibouti"},
		Kpital{City: "Djoubanote", Country: "Soudan du Sud"},
		Kpital{City: "Dodoma", Country: "Tanzanie"},
		Kpital{City: "Doha", Country: "Qatar"},
		Kpital{City: "Douchanbé", Country: "Tadjikistan"},
		Kpital{City: "Dublin", Country: "Irlande"},
		Kpital{City: "Erevan", Country: "Arménie"},
		Kpital{City: "Freetown", Country: "Sierra Leone"},
		Kpital{City: "Funafuti (Vaiaku)7", Country: "Tuvalu"},
		Kpital{City: "Gaborone", Country: "Botswana"},
		Kpital{City: "Georgetown", Country: "Guyana"},
		Kpital{City: "Gitega (capitale politique)", Country: "Burundi"},
		Kpital{City: "Guatemala", Country: "Guatemala"},
		Kpital{City: "Hanoï", Country: "Viêt Nam"},
		Kpital{City: "Harare", Country: "Zimbabwe"},
		Kpital{City: "La Havane", Country: "Cuba"},
		Kpital{City: "Helsinki", Country: "Finlande"},
		Kpital{City: "Honiara", Country: "Salomon"},
		Kpital{City: "Islamabad", Country: "Pakistan"},
		Kpital{City: "Jakarta", Country: "Indonésie"},
		Kpital{City: "Jérusalem (contesté au profit de Tel-aviv)note 2", Country: "Israël"},
		Kpital{City: "Ramallah (capitale proclamée : Jérusalem-est)note 3", Country: "Palestine (État de)23"},
		Kpital{City: "Kaboul", Country: "Afghanistan"},
		Kpital{City: "Kampala", Country: "Ouganda"},
		Kpital{City: "Katmandou", Country: "Népal"},
		Kpital{City: "Khartoum", Country: "Soudan"},
		Kpital{City: "Kiev", Country: "Ukraine"},
		Kpital{City: "Kigali", Country: "Rwanda"},
		Kpital{City: "Kingston", Country: "Jamaïque"},
		Kpital{City: "Kingstown", Country: "Saint-Vincent-et-les-Grenadines"},
		Kpital{City: "Kinshasa", Country: "République démocratique du Congo"},
		Kpital{City: "Melekeok", Country: "Palaos"},
		Kpital{City: "Koweït", Country: "Koweït"},
		Kpital{City: "Kuala Lumpur", Country: "Malaisie"},
		Kpital{City: "Libreville", Country: "Gabon"},
		Kpital{City: "Lilongwe", Country: "Malawi"},
		Kpital{City: "Lima", Country: "Pérou"},
		Kpital{City: "Lisbonne", Country: "Portugal"},
		Kpital{City: "Ljubljana", Country: "Slovénie"},
		Kpital{City: "Lomé", Country: "Togo"},
		Kpital{City: "Londres", Country: "Royaume-Uni"},
		Kpital{City: "Luanda", Country: "Angola"},
		Kpital{City: "Lusaka", Country: "Zambie"},
		Kpital{City: "Luxembourg", Country: "Luxembourg"},
		Kpital{City: "Madrid", Country: "Espagne"},
		Kpital{City: "Malabo", Country: "Guinée équatoriale"},
		Kpital{City: "Malé", Country: "Maldives"},
		Kpital{City: "Managua", Country: "Nicaragua"},
		Kpital{City: "Manama", Country: "Bahreïn"},
		Kpital{City: "Manille", Country: "Philippines"},
		Kpital{City: "Maputo", Country: "Mozambique"},
		Kpital{City: "Mascate", Country: "Oman"},
		Kpital{City: "Maseru", Country: "Lesotho"},
		Kpital{City: "Mbabane", Country: "Eswatini"},
		Kpital{City: "Mexico", Country: "Mexique"},
		Kpital{City: "Minsk", Country: "Biélorussie"},
		Kpital{City: "Mogadiscio (Muqdisho26)", Country: "Somalie"},
		Kpital{City: "Monaco", Country: "Monaco"},
		Kpital{City: "Monrovia", Country: "Liberia"},
		Kpital{City: "Montevideo", Country: "Uruguay"},
		Kpital{City: "Moroni", Country: "Comores"},
		Kpital{City: "Moscou", Country: "Russie"},
		Kpital{City: "Nairobi", Country: "Kenya"},
		Kpital{City: "Nassau", Country: "Bahamas"},
		Kpital{City: "Naypyidaw", Country: "Birmanie"},
		Kpital{City: "N'Djaména", Country: "Tchad"},
		Kpital{City: "New Delhi", Country: "Inde"},
		Kpital{City: "Niamey", Country: "Niger"},
		Kpital{City: "Nicosie", Country: "Chypre"},
		Kpital{City: "Nouakchott", Country: "Mauritanie"},
		Kpital{City: "Noursoultan28", Country: "Kazakhstan"},
		Kpital{City: "Nuku'alofa", Country: "Tonga"},
		Kpital{City: "Oslo", Country: "Norvège"},
		Kpital{City: "Ottawa", Country: "Canada"},
		Kpital{City: "Ouagadougou", Country: "Burkina Faso"},
		Kpital{City: "Oulan-Bator", Country: "Mongolie"},
		Kpital{City: "Palikir", Country: "États fédérés de Micronésie"},
		Kpital{City: "Panama", Country: "Panama"},
		Kpital{City: "Paramaribo", Country: "Suriname"},
		Kpital{City: "Paris", Country: "France"},
		Kpital{City: "La Paz (gouvernement et ambassades)29", Country: "Bolivie"},
		Kpital{City: "Pékin", Country: "Chine"},
		Kpital{City: "Phnom Penh", Country: "Cambodge"},
		Kpital{City: "Podgorica", Country: "Monténégro"},
		Kpital{City: "Port Moresby", Country: "Papouasie-Nouvelle-Guinée"},
		Kpital{City: "Port-au-Prince", Country: "Haïti"},
		Kpital{City: "Port-d'Espagne", Country: "Trinité-et-Tobago"},
		Kpital{City: "Port-Louis", Country: "Maurice"},
		Kpital{City: "Porto-Novo (administrative)30", Country: "Bénin"},
		Kpital{City: "Port-Vila", Country: "Vanuatu"},
		Kpital{City: "Prague", Country: "République tchèque"},
		Kpital{City: "Praia", Country: "Cap-Vert"},
		Kpital{City: "Pretoria (capitale administrative)4", Country: "Afrique du Sud"},
		Kpital{City: "Putrajaya (capitale administrative)25", Country: "Malaisie"},
		Kpital{City: "Pyongyang", Country: "Corée du Nord"},
		Kpital{City: "Quito", Country: "Équateur"},
		Kpital{City: "Rabat", Country: "Maroc"},
		Kpital{City: "Ramallah (capitale de facto)note 3", Country: "Palestine (État de)23"},
		Kpital{City: "Reykjavik", Country: "Islande"},
		Kpital{City: "Riga", Country: "Lettonie"},
		Kpital{City: "Riyad", Country: "Arabie saoudite"},
		Kpital{City: "Rome", Country: "Italie"},
		Kpital{City: "Roseau", Country: "Dominique"},
		Kpital{City: "Saint John's", Country: "Antigua-et-Barbuda"},
		Kpital{City: "Saint-Domingue", Country: "République dominicaine"},
		Kpital{City: "Saint-Georges", Country: "Grenade"},
		Kpital{City: "Saint-Marin", Country: "Saint-Marin"},
		Kpital{City: "San José", Country: "Costa Rica"},
		Kpital{City: "San Salvador", Country: "Salvador"},
		Kpital{City: "Sanaa", Country: "Yémen"},
		Kpital{City: "Santiago", Country: "Chili"},
		Kpital{City: "São Tomé", Country: "Sao Tomé-et-Principe"},
		Kpital{City: "Sarajevo", Country: "Bosnie-Herzégovine"},
		Kpital{City: "Séoul", Country: "Corée du Sud"},
		Kpital{City: "Singapour", Country: "Singapour"},
		Kpital{City: "Skopje", Country: "Macédoine du Nord"},
		Kpital{City: "Sofia", Country: "Bulgarie"},
		Kpital{City: "Stockholm", Country: "Suède"},
		Kpital{City: "Sucre (constitutionnelle)29", Country: "Bolivie"},
		Kpital{City: "Suva", Country: "Fidji"},
		Kpital{City: "Tachkent", Country: "Ouzbékistan"},
		Kpital{City: "Tallinn", Country: "Estonie"},
		Kpital{City: "Tarawa-Sud31", Country: "Kiribati"},
		Kpital{City: "Tbilissi", Country: "Géorgie"},
		Kpital{City: "Tegucigalpa", Country: "Honduras"},
		Kpital{City: "Téhéran", Country: "Iran"},
		Kpital{City: "Tel Aviv-Jaffanote", Country: "Israël"},
		Kpital{City: "Thimphou", Country: "Bhoutan"},
		Kpital{City: "Tirana", Country: "Albanie"},
		Kpital{City: "Tokyo", Country: "Japon"},
		Kpital{City: "Tripoli", Country: "Libye"},
		Kpital{City: "Tunis", Country: "Tunisie"},
		Kpital{City: "Vaduz", Country: "Liechtenstein"},
		Kpital{City: "La Valette", Country: "Malte"},
		Kpital{City: "Varsovie", Country: "Pologne"},
		Kpital{City: "Cité du Vatican", Country: "Vatican"},
		Kpital{City: "Victoria", Country: "Seychelles"},
		Kpital{City: "Vienne", Country: "Autriche"},
		Kpital{City: "Vientiane", Country: "Laos"},
		Kpital{City: "Vilnius", Country: "Lituanie"},
		Kpital{City: "Washington D.C.", Country: "États-Unis"},
		Kpital{City: "Wellington", Country: "Nouvelle-Zélande"},
		Kpital{City: "Windhoek", Country: "Namibie"},
		Kpital{City: "Yamoussoukro", Country: "Côte d'Ivoire"},
		Kpital{City: "Yaoundé", Country: "Cameroun"},
		Kpital{City: "Yaren", Country: "Nauru"},
		Kpital{City: "Zagreb", Country: "Croatie"},
	}
}

