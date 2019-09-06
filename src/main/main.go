package main


import (
	"fmt"
	"net/http"
	"io"
	"time"
)


var m = map[string]string {
	"users": `{
"id": 123456,
"nickname": "RADIUS",
"registration_date": "1999-12-21T00:00:00.000-04:00",
"country_id": "BR",
"address": {
"city": "Bauru",
"state": "BR-SP"
},
"user_type": "normal",
"tags": [
"normal"
],
"logo": null,
"points": 0,
"site_id": "MLB",
"permalink": "http://perfil.mercadolivre.com.br/RADIUS",
"seller_reputation": {
"level_id": null,
"power_seller_status": null,
"transactions": {
"canceled": 0,
"completed": 0,
"period": "historic",
"ratings": {
"negative": 0,
"neutral": 0,
"positive": 0
},
"total": 0
}
},
"buyer_reputation": {
"tags": []
},
"status": {
"site_status": "active"
}
}`,
	"sites": `{
  "id": "MLB",
  "name": "Brasil",
  "country_id": "BR",
  "sale_fees_mode": "not_free",
  "mercadopago_version": 3,
  "default_currency_id": "BRL",
  "immediate_payment": "optional",
  "payment_method_ids": [
    "MLBMP",
    "MLBWC",
    "MLBMO",
    "MLBBC",
    "MLBCC",
    "MLBDE",
    "MLBCD",
    "MLBOU"
  ],
  "settings": {
    "identification_types": [
      "CPF",
      "CNPJ",
      "CRECI"
    ],
    "taxpayer_types": [
    ],
    "identification_types_rules": null
  },
  "currencies": [
    {
      "id": "BRL",
      "symbol": "R$"
    }
  ],
  "categories": [
    {
      "id": "MLB5672",
      "name": "Acessórios para Veículos"
    },
    {
      "id": "MLB1499",
      "name": "Agro, Indústria e Comércio"
    },
    {
      "id": "MLB1403",
      "name": "Alimentos e Bebidas"
    },
    {
      "id": "MLB1071",
      "name": "Animais"
    },
    {
      "id": "MLB1367",
      "name": "Antiguidades"
    },
    {
      "id": "MLB1368",
      "name": "Arte e Artesanato"
    },
    {
      "id": "MLB1384",
      "name": "Bebês"
    },
    {
      "id": "MLB1246",
      "name": "Beleza e Cuidado Pessoal"
    },
    {
      "id": "MLB1132",
      "name": "Brinquedos e Hobbies"
    },
    {
      "id": "MLB1430",
      "name": "Calçados, Roupas e Bolsas"
    },
    {
      "id": "MLB1039",
      "name": "Câmeras e Acessórios"
    },
    {
      "id": "MLB1743",
      "name": "Carros, Motos e Outros"
    },
    {
      "id": "MLB1574",
      "name": "Casa, Móveis e Decoração"
    },
    {
      "id": "MLB1051",
      "name": "Celulares e Telefones"
    },
    {
      "id": "MLB1798",
      "name": "Coleções e Comics"
    },
    {
      "id": "MLB5726",
      "name": "Eletrodomésticos"
    },
    {
      "id": "MLB1000",
      "name": "Eletrônicos, Áudio e Vídeo"
    },
    {
      "id": "MLB1276",
      "name": "Esportes e Fitness"
    },
    {
      "id": "MLB263532",
      "name": "Ferramentas e Construção"
    },
    {
      "id": "MLB3281",
      "name": "Filmes e Seriados"
    },
    {
      "id": "MLB1144",
      "name": "Games"
    },
    {
      "id": "MLB1459",
      "name": "Imóveis"
    },
    {
      "id": "MLB1648",
      "name": "Informática"
    },
    {
      "id": "MLB218519",
      "name": "Ingressos"
    },
    {
      "id": "MLB1182",
      "name": "Instrumentos Musicais"
    },
    {
      "id": "MLB3937",
      "name": "Joias e Relógios"
    },
    {
      "id": "MLB1196",
      "name": "Livros, Revistas e Comics"
    },
    {
      "id": "MLB1168",
      "name": "Música"
    },
    {
      "id": "MLB264586",
      "name": "Saúde"
    },
    {
      "id": "MLB1540",
      "name": "Serviços"
    },
    {
      "id": "MLB1953",
      "name": "Mais Categorias"
    }
  ]
}`,

	"countries": `{
  "id": "BR",
  "name": "Brasil",
  "locale": "pt_BR",
  "currency_id": "BRL",
  "decimal_separator": ",",
  "thousands_separator": ".",
  "time_zone": "GMT-03:00",
  "geo_information": {
    "location": {
      "latitude": -23.6821604,
      "longitude": -46.875494
    }
  },
  "states": [
    {
      "id": "BR-AC",
      "name": "Acre"
    },
    {
      "id": "BR-AL",
      "name": "Alagoas"
    },
    {
      "id": "BR-AP",
      "name": "Amapá"
    },
    {
      "id": "BR-AM",
      "name": "Amazonas"
    },
    {
      "id": "BR-BA",
      "name": "Bahia"
    },
    {
      "id": "BR-CE",
      "name": "Ceará"
    },
    {
      "id": "BR-DF",
      "name": "Distrito Federal"
    },
    {
      "id": "BR-ES",
      "name": "Espírito Santo"
    },
    {
      "id": "BR-GO",
      "name": "Goiás"
    },
    {
      "id": "BR-MA",
      "name": "Maranhão"
    },
    {
      "id": "BR-MT",
      "name": "Mato Grosso"
    },
    {
      "id": "BR-MS",
      "name": "Mato Grosso do Sul"
    },
    {
      "id": "BR-MG",
      "name": "Minas Gerais"
    },
    {
      "id": "BR-PR",
      "name": "Paraná"
    },
    {
      "id": "BR-PB",
      "name": "Paraíba"
    },
    {
      "id": "BR-PA",
      "name": "Pará"
    },
    {
      "id": "BR-PE",
      "name": "Pernambuco"
    },
    {
      "id": "BR-PI",
      "name": "Piauí"
    },
    {
      "id": "BR-RN",
      "name": "Rio Grande do Norte"
    },
    {
      "id": "BR-RS",
      "name": "Rio Grande do Sul"
    },
    {
      "id": "BR-RJ",
      "name": "Rio de Janeiro"
    },
    {
      "id": "BR-RO",
      "name": "Rondônia"
    },
    {
      "id": "BR-RR",
      "name": "Roraima"
    },
    {
      "id": "BR-SC",
      "name": "Santa Catarina"
    },
    {
      "id": "BR-SE",
      "name": "Sergipe"
    },
    {
      "id": "BR-SP",
      "name": "São Paulo"
    },
    {
      "id": "BR-TO",
      "name": "Tocantins"
    }
  ]
}`,

}

func main() {
	moc()
}



func homePage(w http.ResponseWriter, r *http.Request) {
	time.Sleep(time.Millisecond*500)
	fmt.Println(w, "Respuesta recibida")
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	io.WriteString(w, "HOMEPAGE")
}


func homePageUsers(w http.ResponseWriter, r *http.Request) {
	time.Sleep(time.Millisecond*100)
	fmt.Println(w, "Respuesta recibida")
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	io.WriteString(w, m["users"])
}


func homePageCountries(w http.ResponseWriter, r *http.Request) {
	time.Sleep(time.Millisecond*11000)
	fmt.Println(w, "Respuesta recibida")
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	io.WriteString(w, m["countries"])
}


func homePageSites(w http.ResponseWriter, r *http.Request) {
	time.Sleep(time.Millisecond*1500)
	fmt.Println(w, "Respuesta recibida")
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	io.WriteString(w, m["sites"])
}

func pong(w http.ResponseWriter, r *http.Request) {
	time.Sleep(time.Millisecond * 200)
	w.WriteHeader(http.StatusOK)
	io.WriteString(w, "Pong")
}
func moc() { // para hcacer un get necesitamos un browser y para un browser necesitamos un servidor
	http.HandleFunc("/", homePage)
	http.HandleFunc("/users/", homePageUsers)
	http.HandleFunc("/countries/BR", homePageCountries)
	http.HandleFunc("/sites/MLB", homePageSites)
	http.HandleFunc("/users/ping", pong)
	http.HandleFunc("/countries/ping", pong)
	http.HandleFunc("/sites/ping", pong)
	http.ListenAndServe(":8090", nil)
}