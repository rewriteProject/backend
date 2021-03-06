package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/Jeffail/gabs/v2"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"net/http"
	"os"
	//"strconv"
	"strings"
)

// Setting environment variables
var databaseUser string = os.Getenv("DATABASE_USER")
var databasePassword string = os.Getenv("DATABASE_PASSWORD")
var databasePort string = os.Getenv("DATABASE_PORT")
var databaseHost string = os.Getenv("DATABASE_HOST")

// Struct für Table Customer
type Customer struct {
	id_user       int
	gender        string
	city          string
	country       string
	amount_orders int
}

// Struct für Table Container
type Container struct {
	id_container   int
	container_name string
	amount_bills   int
	to_country     string
	create_date    string
	close_date     string
	status         bool
	curr_value     float64
	curr_weight    int
	max_weight     int
	curr_volume    int
	max_volume     int
}

// Struct für Table Bill
type Bill struct {
	id_bill      int
	id_container int
	id_user      int
	bill_date    string
	bill_city    string
	bill_country string
	total_value  float64
	total_weight float64
	total_volume float64
}

// Struct für Table Product
type Product struct {
	id_product    int
	id_bill       int
	product_name  string
	product_value float64
	amount        int
	color         string
	brand         string
	category      string
	weight        float64
	volume        float64
}

// Durchsucht einen Slice nach doppelten Einträgen
func checkForDoubleElement(list []string, s string) bool {
	for _, element := range list {
		if element == s {
			return true
		}
	}
	return false
}

/*
* - ResponseWriter Interface wird vom Router benutzt um eine Response zu erstellen
* - Request ist die eingehende Anfrage, die der Server empfängt
 */

// Selects für Informationen, die die Website beim Laden braucht
// ------------------------------------------------------------------------------------------------------------------------------------------------
// Gibt Zielländer der Container-Tabelle aus
func DestinationCountries(w http.ResponseWriter, r *http.Request) {

	//dbConn, error := sql.Open("mysql", databaseUser + ":" + databasePassword + "@tcp(" + databaseHost + ":" + databasePort + ")/logistics")
	//if error != nil {
	//	panic(error.Error())
	//}
	//defer dbConn.Close()
	//
	//results, error := dbConn.Query("Select to_country from Container")
	//if error != nil {
	//	panic(error.Error())
	//}
	//defer results.Close()

	//countries := make([]string, 0)
	countries := []string{"China", "Russia", "USA"}

	//for results.Next() {
	//	var container Container
	//	error = results.Scan(&container.to_country)
	//	if error != nil {
	//		panic(error.Error())
	//	}
	//
	//	if checkForDoubleElement(countries, container.to_country) == false {
	//		countries = append(countries, container.to_country)
	//	}
	//}
	jsonMap := map[string]interface{}{
		"countries": countries}
	json.NewEncoder(w).Encode(jsonMap)
}

// ------------------------------------------------------------------------------------------------------------------------------------------------
// Gibt Marke, Farbe und Kategory der Produkt-Tabelle aus
func ProductProperties(w http.ResponseWriter, r *http.Request) {
	//
	//dbConn, error := sql.Open("mysql", databaseUser + ":" + databasePassword + "@tcp(" + databaseHost + ":" + databasePort + ")/logistics")
	//if error != nil {
	//	panic(error.Error())
	//}
	//defer dbConn.Close()
	//
	//results, error := dbConn.Query("Select color, brand, category from Product")
	//if error != nil {
	//	panic(error.Error())
	//}
	//defer results.Close()

	// Slices, um Einträge der DB-Columns aufzuteilen
	//color := make([]string, 0)
	//brand := make([]string, 0)
	//cat := make([]string, 0)
	color := []string{"Green", "Red", "Blue"}
	brand := []string{"Nike", "Adidas", "Puma"}
	cat := []string{"Tshirt", "Pants", "Shoes"}

	//for results.Next() {
	//	var product Product
	//	error = results.Scan(&product.color, &product.brand, &product.category)
	//	if error != nil {
	//		panic(error.Error())
	//	}
	//
	//	// Doppelte Einträge in Slices filtern
	//	if checkForDoubleElement(color, product.color) == false {
	//		color = append(color, product.color)
	//	}
	//
	//	if checkForDoubleElement(brand, product.brand) == false {
	//		brand = append(brand, product.brand)
	//	}
	//
	//	if checkForDoubleElement(cat, product.category) == false {
	//		cat = append(cat, product.category)
	//	}
	//
	//}
	jsonMap := map[string]interface{}{
		"color":    color,
		"brand":    brand,
		"category": cat}
	json.NewEncoder(w).Encode(jsonMap)
}

// Selects für die Analyse Komponente; Dies geschieht mittels Prepared Statements
// ------------------------------------------------------------------------------------------------------------------------------------------------
// ------------------------------------------------------------------------------------------------------------------------------------------------
// Information
func InformationHandler(w http.ResponseWriter, r *http.Request) {
	//dbConn, error := sql.Open("mysql", databaseUser + ":" + databasePassword + "@tcp(" + databaseHost + ":" + databasePort + ")/logistics")
	//if error != nil {
	//	panic(error.Error())
	//}
	//defer dbConn.Close()

	// Abrufen der Pfad Parameter
	parameters := mux.Vars(r)
	Case := parameters["Case"]
	countries := parameters["countries"]
	//countries_list := strings.Split(countries, ",")

	// Fallunterscheidung; Jeweils passendes Prepared Statement wird mit Parameter aufgerufen
	// Fall Information I1
	if Case == "i1" {
		//results, error := dbConn.Query("select id_container, to_country, create_date from Container where status='1'")
		//if error != nil {
		//	panic(error.Error())
		//}
		//defer results.Close()

		//Gabs Package für erweitere json Strukturen
		nestedjson := gabs.New()
		_, _ = nestedjson.Set("OPEN", "container", "status")

		if countries == "all" {
			//for results.Next() {
			//	var container Container
			//	error = results.Scan(&container.id_container, &container.to_country, &container.create_date)
			//	if error != nil {
			//		panic(error.Error())
			//	}
			//jsonMap := map[string]interface{}{
			//	"container_id": 1,
			//	"open_date": "2021-01-30"}
			//_, _ = nestedjson.SetP(jsonMap, "container.country."+container.to_country)

			jsonMap1 := map[string]interface{}{
				"container_id": 1,
				"open_date":    "2021-01-30"}
			_, _ = nestedjson.SetP(jsonMap1, "container.country.china")

			jsonMap2 := map[string]interface{}{
				"container_id": 2,
				"open_date":    "2021-01-24"}
			_, _ = nestedjson.SetP(jsonMap2, "container.country.russia")

			jsonMap3 := map[string]interface{}{
				"container_id": 3,
				"open_date":    "2021-01-19"}
			_, _ = nestedjson.SetP(jsonMap3, "container.country.usa")
			//}

			json.NewEncoder(w).Encode(nestedjson)

		} else {
			//for results.Next() {
			//	var container Container
			//	error = results.Scan(&container.id_container, &container.to_country, &container.create_date)
			//	if error != nil {
			//		panic(error.Error())
			//	}

			// Nur falls das angegebene Land in der Selektion vorkommt, wird es hinzugefügt
			//for _, country := range countries_list {
			//	if country == container.to_country {
			//		jsonMap := map[string]interface{}{
			//			"container_id": container.id_container,
			//			"open_date":    container.create_date}
			//		_, _ = nestedjson.SetP(jsonMap, "container.country."+container.to_country)
			//	}
			//}
			//}
			jsonMap1 := map[string]interface{}{
				"container_id": 1,
				"open_date":    "2021-01-30"}
			_, _ = nestedjson.SetP(jsonMap1, "container.country."+countries)

			json.NewEncoder(w).Encode(nestedjson)
		}

		// Fall Information I2
	} else if Case == "i2" {
		//results, error := dbConn.Query("select id_container, to_country, curr_weight, max_weight from Container where status='1'")
		//if error != nil {
		//	panic(error.Error())
		//}
		//defer results.Close()

		//Gabs Package für erweitere json Strukturen
		nestedjson := gabs.New()
		_, _ = nestedjson.Set("OPEN", "container", "status")

		if countries == "all" {
			//for results.Next() {
			//	var container Container
			//	error = results.Scan(&container.id_container, &container.to_country, &container.curr_weight, &container.max_weight)
			//	if error != nil {
			//		panic(error.Error())
			//	}

			//jsonMap := map[string]interface{}{
			//	"container_id":   container.id_container,
			//	"curr_weight_kg": container.curr_weight,
			//	"max_weight_kg":  container.max_weight}
			//_, _ = nestedjson.SetP(jsonMap, "container.country."+container.to_country)
			//}
			jsonMap1 := map[string]interface{}{
				"container_id":   1,
				"curr_weight_kg": 22,
				"max_weight_kg":  25}
			_, _ = nestedjson.SetP(jsonMap1, "container.country.china")

			jsonMap2 := map[string]interface{}{
				"container_id":   2,
				"curr_weight_kg": 8,
				"max_weight_kg":  25}
			_, _ = nestedjson.SetP(jsonMap2, "container.country.russia")

			jsonMap3 := map[string]interface{}{
				"container_id":   3,
				"curr_weight_kg": 15,
				"max_weight_kg":  25}
			_, _ = nestedjson.SetP(jsonMap3, "container.country.usa")

			json.NewEncoder(w).Encode(nestedjson)

		} else {
			//for results.Next() {
			//	var container Container
			//	error = results.Scan(&container.id_container, &container.to_country, &container.curr_weight, &container.max_weight)
			//	if error != nil {
			//		panic(error.Error())
			//	}
			//
			//	// Nur falls das angegebene Land in der Selektion vorkommt, wird es hinzugefügt
			//	for _, country := range countries_list {
			//		if country == container.to_country {
			//			jsonMap := map[string]interface{}{
			//				"container_id":   container.id_container,
			//				"curr_weight_kg": container.curr_weight,
			//				"max_weight_kg":  container.max_weight}
			//			_, _ = nestedjson.SetP(jsonMap, "container.country."+container.to_country)
			//		}
			//	}
			//}
			jsonMap1 := map[string]interface{}{
				"container_id":   1,
				"curr_weight_kg": 22,
				"max_weight_kg":  25}
			_, _ = nestedjson.SetP(jsonMap1, "container.country."+countries)

			json.NewEncoder(w).Encode(nestedjson)
		}
	}
}

// ------------------------------------------------------------------------------------------------------------------------------------------------
// Statistik
func StatisticHandler(w http.ResponseWriter, r *http.Request) {
	//dbConn, error := sql.Open("mysql", databaseUser + ":" + databasePassword + "@tcp(" + databaseHost + ":" + databasePort + ")/logistics")
	//if error != nil {
	//	panic(error.Error())
	//}
	//defer dbConn.Close()

	// Abrufen der Pfad Parameter
	parameters := mux.Vars(r)
	country := parameters["country"]
	attributes := parameters["attributes"]
	attributes_list := strings.Split(attributes, ",")
	minDate, _ := r.URL.Query()["minDate"]
	maxDate, _ := r.URL.Query()["maxDate"]

	//Schickt die Query an die Datenbank
	//results, error := dbConn.Query("Select * from customer")
	//if error != nil {
	//	panic(error.Error())
	//}
	//defer results.Close()
	nestedjson := gabs.New()
	_, _ = nestedjson.Set("CLOSE", "container", "status")
	_, _ = nestedjson.Set(minDate[0], "container", "start_time")
	_, _ = nestedjson.Set(maxDate[0], "container", "end_time")
	//_, _ = nestedjson.SetP("", "container.country."+country)
	fmt.Println(attributes)
	fmt.Println(attributes_list)
	jsonMap := map[string]interface{}{}

	if checkForDoubleElement(attributes_list, "color") == true {
		jsonMap = map[string]interface{}{
			"red":   2,
			"blue":  5,
			"green": 10}
		_, _ = nestedjson.SetP(jsonMap, "container.country."+country+"."+"color")
		_, _ = nestedjson.Set("color", "container", "type")
	}
	if checkForDoubleElement(attributes_list, "brand") == true {
		jsonMap = map[string]interface{}{
			"nike":   7,
			"adidas": 3,
			"hm":     7}
		_, _ = nestedjson.SetP(jsonMap, "container.country."+country+"."+"brand")
		_, _ = nestedjson.Set("brand", "container", "type")
	}
	if checkForDoubleElement(attributes_list, "category") == true {
		jsonMap = map[string]interface{}{
			"t-shirt": 6,
			"shoes":   9,
			"pants":   2}
		_, _ = nestedjson.SetP(jsonMap, "container.country."+country+"."+"category")
		_, _ = nestedjson.Set("category", "container", "type")
	}
	if checkForDoubleElement(attributes_list, "all") == true {
		jsonMap = map[string]interface{}{
			"red":   2,
			"blue":  5,
			"green": 10}
		_, _ = nestedjson.SetP(jsonMap, "container.country."+country+"."+"color")
		jsonMap = map[string]interface{}{
			"nike":   7,
			"adidas": 3,
			"hm":     7}
		_, _ = nestedjson.SetP(jsonMap, "container.country."+country+"."+"brand")
		jsonMap = map[string]interface{}{
			"t-shirt": 6,
			"shoes":   9,
			"pants":   2}
		_, _ = nestedjson.SetP(jsonMap, "container.country."+country+"."+"category")
		_, _ = nestedjson.Set("", "container", "type")
	}
	//jsonMap = map[string]interface{}{
	//	"red": 2,
	//	"blue": 5,
	//	"green": 10}
	//_, _ = nestedjson.SetP(jsonMap, "container.country."+country+"."+"color")
	json.NewEncoder(w).Encode(nestedjson)

}

// ------------------------------------------------------------------------------------------------------------------------------------------------
type P2 struct {
	id_bill      int
	id_container int
	close_date   string
	color        string
	brand        string
	category     string
}

// Prognose
func ForecastHandler(w http.ResponseWriter, r *http.Request) {
	//dbConn, error := sql.Open("mysql", databaseUser + ":" + databasePassword + "@tcp(" + databaseHost + ":" + databasePort + ")/logistics")
	//if error != nil {
	//	panic(error.Error())
	//}
	//defer dbConn.Close()

	// Abrufen der Pfad Parameter
	parameters := mux.Vars(r)
	Case := parameters["Case"]
	country := parameters["country"]
	minDate, _ := r.URL.Query()["minDate"]
	//year, _ := r.URL.Query()["year"]
	typ, _ := r.URL.Query()["typ"]
	feature, _ := r.URL.Query()["feature"]

	// Fallunterscheidung; Jeweils passendes Prepared Statement wird mit Parameter aufgerufen
	// Fall Prognose P1-1
	if Case == "p1-1" {
		//Schickt die Query an die Datenbank
		//query := "select to_country, create_date, close_date from Container where status='0' and to_country='" + country + "' and create_date>='" + minDate[0] + "'"
		//results, error := dbConn.Query(query)
		//if error != nil {
		//	panic(error.Error())
		//}
		//defer results.Close()

		//Gabs Package für erweitere json Strukturen
		nestedjson := gabs.New()
		_, _ = nestedjson.Set("CLOSE", "container", "status")
		_, _ = nestedjson.Set(country, "container", "country")
		_, _ = nestedjson.Set(minDate[0], "container", "min_date")
		_, _ = nestedjson.Set("now", "container", "max_date")

		//counter := 1
		//for results.Next() {
		//	var container Container
		//	error = results.Scan(&container.to_country, &container.create_date, &container.close_date)
		//	if error != nil {
		//		panic(error.Error())
		//	}

		//jsonMap := map[string]interface{}{
		//	"open_date":  container.create_date,
		//	"close_date": container.close_date}
		//_, _ = nestedjson.SetP(jsonMap, "container.dates."+strconv.Itoa(counter))
		//counter += 1
		//}
		jsonMap1 := map[string]interface{}{
			"open_date":  "2019-01-01",
			"close_date": "2019-01-30"}
		_, _ = nestedjson.SetP(jsonMap1, "container.dates."+"1")

		jsonMap2 := map[string]interface{}{
			"open_date":  "2019-02-04",
			"close_date": "2019-03-02"}
		_, _ = nestedjson.SetP(jsonMap2, "container.dates."+"2")

		jsonMap3 := map[string]interface{}{
			"open_date":  "2019-05-10",
			"close_date": "2019-05-29"}
		_, _ = nestedjson.SetP(jsonMap3, "container.dates."+"3")

		json.NewEncoder(w).Encode(nestedjson)

	} else if Case == "p1-2" {
		//Schickt die Query an die Datenbank
		//query := "select to_country, create_date from Container where status='1' and to_country='" + country + "'"
		//fmt.Println(query)
		//results, error := dbConn.Query(query)
		//if error != nil {
		//	panic(error.Error())
		//}
		//defer results.Close()

		//Gabs Package für erweitere json Strukturen
		nestedjson := gabs.New()
		_, _ = nestedjson.Set("OPEN", "container", "status")
		_, _ = nestedjson.Set(country, "container", "country")

		//for results.Next() {
		//	var container Container
		//	error = results.Scan(&container.to_country, &container.create_date)
		//	if error != nil {
		//		panic(error.Error())
		//	}
		//
		//	_, _ = nestedjson.Set(container.create_date, "container", "create_date")
		//}
		_, _ = nestedjson.Set("2021-02-20", "container", "create_date")

		json.NewEncoder(w).Encode(nestedjson)

	} else if Case == "p2" {
		// Zu welchem Monat ein Container gezählt wird hängt vom Closedate ab
		//query := "select bill.id_bill, container.id_container, container.to_country, container.close_date, product.color, product.brand, product.category from container inner join bill on bill.id_container=container.id_container right join product on bill.id_bill=product.id_bill where status='0' and to_country='" + country + "' and close_date between '" + minDate[0] + "' and now()"
		//results, error := dbConn.Query(query)
		//if error != nil {
		//	panic(error.Error())
		//}
		//defer results.Close()

		//Gabs Package für erweitere json Strukturen
		nestedjson := gabs.New()
		_, _ = nestedjson.Set("CLOSE", "container", "status")
		_, _ = nestedjson.Set(country, "container", "country")
		_, _ = nestedjson.Set(typ[0], "container", "type")
		_, _ = nestedjson.Set(minDate[0], "container", "min_date")
		_, _ = nestedjson.Set("m", "container", "intervall")

		// Für jeden Monat die Anzahl an Einträgen für gegebenes Merkmal (color, brand, category)
		//_, _ = nestedjson.Set(feature[0], "container", "intervall")

		//list := make([]string, 0)
		//jsonList := make([]string, 0)
		//for results.Next() {
		//	var p2 P2
		//	error = results.Scan(&p2.close_date)
		//	if error != nil {
		//		panic(error.Error())
		//	}
		//	//list // append trimmed Date
		//
		//}

		/*
			Zusätzliche Schleife für Datumsliste
				Überprüfen ob gleiches Datum, wenn ja dann steigt counter
		*/

		jsonMap := map[string]interface{}{
			"01-2015": 56,
			"02-2015": 47,
			"03-2015": 87,
			"04-2015": 101,
			"05-2015": 168,
			"06-2015": 216,
			"07-2015": 231,
			"08-2015": 174,
			"09-2015": 168,
			"10-2015": 98,
			"11-2015": 65,
			"12-2015": 60,
			"01-2016": 41,
			"02-2016": 49,
			"03-2016": 71,
			"04-2016": 92,
			"05-2016": 156,
			"06-2016": 189,
			"07-2016": 198,
			"08-2016": 145,
			"09-2016": 101,
			"10-2016": 86,
			"11-2016": 57,
			"12-2016": 39,
			"01-2017": 42,
			"02-2017": 53,
			"03-2017": 65,
			"04-2017": 86,
			"05-2017": 91,
			"06-2017": 130,
			"07-2017": 156,
			"08-2017": 148,
			"09-2017": 132,
			"10-2017": 98,
			"11-2017": 85,
			"12-2017": 70,
			"01-2018": 73,
			"02-2018": 81,
			"03-2018": 88,
			"04-2018": 107,
			"05-2018": 136,
			"06-2018": 187,
			"07-2018": 212,
			"08-2018": 199,
			"09-2018": 156,
			"10-2018": 121,
			"11-2018": 99,
			"12-2018": 87,
			"01-2019": 71,
			"02-2019": 81,
			"03-2019": 109,
			"04-2019": 122,
			"05-2019": 158,
			"06-2019": 191,
			"07-2019": 220,
			"08-2019": 225,
			"09-2019": 198,
			"10-2019": 156,
			"11-2019": 98,
			"12-2019": 89,
			"01-2020": 92,
			"02-2020": 81,
			"03-2020": 78,
			"04-2020": 54,
			"05-2020": 61,
			"06-2020": 98,
			"07-2020": 121,
			"08-2020": 118,
			"09-2020": 101,
			"10-2020": 81,
			"11-2020": 47,
			"12-2020": 31,
			"01-2021": 33,
			"02-2021": 48,
			"03-2021": 56}
		_, _ = nestedjson.SetP(jsonMap, "container."+typ[0]+"."+feature[0])
		json.NewEncoder(w).Encode(nestedjson)
	}
}

// Selects für Alles aus einer Tabelle
// ------------------------------------------------------------------------------------------------------------------------------------------------
// ------------------------------------------------------------------------------------------------------------------------------------------------
// Fragt alle Informationen des Table Customer ab
func AllCustomers(w http.ResponseWriter, r *http.Request) {
	// Baut Verbindung zu Datenbank auf
	dbConn, error := sql.Open("mysql", databaseUser+":"+databasePassword+"@tcp("+databaseHost+":"+databasePort+")/logistics")
	if error != nil {
		panic(error.Error())
	}
	defer dbConn.Close()

	//Schickt die Query an die Datenbank
	results, error := dbConn.Query("Select * from Customer")
	if error != nil {
		panic(error.Error())
	}
	defer results.Close()

	// Geht die Rückgabe/Ergebnis Zeile für Zeile durch
	for results.Next() {
		var customer Customer
		error = results.Scan(&customer.id_user, &customer.gender, &customer.city, &customer.country, &customer.amount_orders)
		if error != nil {
			panic(error.Error())
		}

		// Erstellt eine Map mit den einzelnen Einträgen und encoded es zu Json
		jsonMap := map[string]interface{}{
			"id_customer":   customer.id_user,
			"gender":        customer.gender,
			"city":          customer.city,
			"country":       customer.country,
			"amount_orders": customer.amount_orders}
		json.NewEncoder(w).Encode(jsonMap)
	}
	fmt.Println("Table Customer ausgegeben")
}

// ------------------------------------------------------------------------------------------------------------------------------------------------

// Fragt alle Informationen des Table Container ab
func AllContainers(w http.ResponseWriter, r *http.Request) {
	dbConn, error := sql.Open("mysql", databaseUser+":"+databasePassword+"@tcp("+databaseHost+":"+databasePort+")/logistics")
	if error != nil {
		panic(error.Error())
	}
	defer dbConn.Close()

	results, error := dbConn.Query("Select to_country from Container")
	if error != nil {
		panic(error.Error())
	}
	defer results.Close()

	for results.Next() {
		var container Container
		//error = results.Scan(&container.id_container, &container.container_name, &container.amount_bills, &container.to_country, &container.create_date, &container.close_date, &container.status, &container.curr_value, &container.curr_weight, &container.max_weight, &container.curr_volume, &container.max_volume)
		error = results.Scan(&container.to_country)
		if error != nil {
			panic(error.Error())
		}

		jsonMap := map[string]interface{}{
			/*"id_container" : container.id_container,
			"container_name" : container.container_name,
			"amount_bills" : container.amount_bills,*/
			"to_country": container.to_country} /*,
		"create_date" : container.create_date,
		"close_date" : container.close_date,
		"status" : container.status,
		"curr_value" : container.curr_value,
		"curr_weight" : container.curr_weight,
		"max_weight" : container.max_weight,
		"curr_volume" : container.curr_volume,
		"max_volume": container. max_volume}*/
		json.NewEncoder(w).Encode(jsonMap)
	}
	fmt.Println("Table Container ausgegeben")
}

// ------------------------------------------------------------------------------------------------------------------------------------------------

// Fragt alle Informationen des Table Bill ab
func AllBills(w http.ResponseWriter, r *http.Request) {
	dbConn, error := sql.Open("mysql", databaseUser+":"+databasePassword+"@tcp("+databaseHost+":"+databasePort+")/logistics")
	if error != nil {
		panic(error.Error())
	}
	defer dbConn.Close()

	results, error := dbConn.Query("Select * from Bill")
	if error != nil {
		panic(error.Error())
	}
	defer results.Close()

	for results.Next() {
		var bill Bill
		error = results.Scan(&bill.id_bill, &bill.id_container, &bill.id_user, &bill.bill_date, &bill.bill_city, &bill.bill_country, &bill.total_value, &bill.total_weight, &bill.total_volume)
		if error != nil {
			panic(error.Error())
		}

		jsonMap := map[string]interface{}{
			"id_bill":      bill.id_bill,
			"id_container": bill.id_container,
			"id_user":      bill.id_user,
			"bill_date":    bill.bill_date,
			"bill_city":    bill.bill_city,
			"bill_country": bill.bill_country,
			"total_value":  bill.total_value,
			"total_weight": bill.total_weight,
			"total_volume": bill.total_volume}
		json.NewEncoder(w).Encode(jsonMap)
	}
	fmt.Println("Table Bill ausgegeben")
}

// ------------------------------------------------------------------------------------------------------------------------------------------------

// Fragt alle Informationen des Table Product ab
func AllProducts(w http.ResponseWriter, r *http.Request) {

	dbConn, error := sql.Open("mysql", databaseUser+":"+databasePassword+"@tcp("+databaseHost+":"+databasePort+")/logistics")
	if error != nil {
		panic(error.Error())
	}
	defer dbConn.Close()

	results, error := dbConn.Query("Select * from Product")
	if error != nil {
		panic(error.Error())
	}
	defer results.Close()

	for results.Next() {
		var product Product
		error = results.Scan(&product.id_product, &product.id_bill, &product.product_name, &product.product_value, &product.amount, &product.color, &product.brand, &product.category, &product.weight, &product.volume)
		if error != nil {
			panic(error.Error())
		}

		jsonMap := map[string]interface{}{
			"id_product":    product.id_product,
			"id_bill":       product.id_bill,
			"product_name":  product.product_name,
			"product_value": product.product_value,
			"amount":        product.amount,
			"color":         product.color,
			"brand":         product.brand,
			"category":      product.category,
			"weight":        product.weight,
			"volume":        product.volume}
		json.NewEncoder(w).Encode(jsonMap)
	}
	fmt.Println("Table Product ausgegeben")
}
