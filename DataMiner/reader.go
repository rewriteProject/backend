package main

// Value in Datenbank geupdatet zu DECIMAL

import (
	"database/sql"
	"encoding/csv"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

// Verzeichnis in dem sich CSV-Dateien befinden
var observedDir = "C:/Users/chris/OneDrive/Dokumente/Diplomarbeit/ReWriteGo/csv-files"

// Slices für INSERTs
var sliceCustomers []string
var sliceContainers []string
var sliceBills []string
var sliceProducts []string

func readFilesAndSaveSlices() {

	//------------------------------------------------------------------
	// Verzeichnis aus dem CSV-Files bezogen werden
	files, error := ioutil.ReadDir(observedDir)
	if error != nil {
		fmt.Println("Error reading directory  \n")
		log.Fatal(error)
	}
	//------------------------------------------------------------------

	// Files mit Schleife durchgehen und INSERTS erstellen
	fmt.Println("List of files: \n")
	for _, file := range files {

		// Aktuelle Datei öffnen
		openedFile, error := os.Open(observedDir + "/" + file.Name())
		if error != nil {
			fmt.Println("Error reading file " + openedFile.Name() + "\n")
			log.Fatal(error)
		}

		defer openedFile.Close()

		// CSV-Reader mit aktueller Datei
		csvReader := csv.NewReader(openedFile)

		// Erste Zeile überspringen
		if _, error := csvReader.Read(); error != nil {
			panic(error)
		}

		// Inhalt des CSV-Files lesen
		lines, error := csvReader.ReadAll()
		if error != nil {
			log.Fatal(error)
		}

		// Dateiname überprüfen und je nach Fall INSERTs erstellen
		switch file.Name() {

		// id_customer, first_name, last_name, email, gender, user_postalcode ,adress, city, country, successfull_order
		case "customer.csv":
			fmt.Println("--------------------------------------------------------------------------------------")
			fmt.Println("Case Customers\n")

			// INSERT-String Zeile für Zeile bauen
			insertStr := "INSERT INTO Customer(id_user, gender, city, country, amount_orders) VALUES ("
			for _, line := range lines {
				values := line[0] + ", '" + line[4] + "', '" + line[7] + "', '" + line[8] + "', " + line[9] + ");"
				a := insertStr + values
				sliceCustomers = append(sliceCustomers, a)
			}
			for _, a := range sliceCustomers {
				fmt.Println(a)
			}
			fmt.Println("--------------------------------------------------------------------------------------\n")

		//	id_container, container_name, bill_count, to_country, create_date, close_date, status, current_value, current_weight, max_weight, current_volume, max_volume
		case "containers.csv":
			fmt.Println("--------------------------------------------------------------------------------------")
			fmt.Println("Case Containers\n")

			// INSERT-String Zeile für Zeile bauen
			insertStr := "INSERT INTO Container(id_container, container_name, amount_bills, to_country, create_date, close_date, status, curr_value, curr_weight, max_weight, curr_volume, max_volume) VALUES ("
			for _, line := range lines {
				// Entscheidung ob Container OPEN/CLOSED ist und anpassen an TRUE/FALSE
				createDate := line[4][6:10] + "-" + line[4][3:5] + "-" + line[4][0:2]
				closeDate := line[5][6:10] + "-" + line[5][3:5] + "-" + line[5][0:2]

				values := ""
				if line[6] == "CLOSED" {
					values += line[0] + ", '" + line[1] + "', " + line[2] + ", '" + line[3] + "', '" + createDate + "', '" + closeDate + "', " + strings.Replace(line[6], "CLOSED", "FALSE", 1) + ", " + strings.Replace(line[7], ",", ".", 1) + ", " + line[8] + ", " + line[9] + ", " + line[10] + ", " + line[11] + ");"
				} else {
					values += line[0] + ", '" + line[1] + "', " + line[2] + ", '" + line[3] + "', '" + createDate + "', '" + closeDate + "', " + strings.Replace(line[6], "OPEN", "TRUE", 1) + ", " + strings.Replace(line[7], ",", ".", 1) + ", " + line[8] + ", " + line[9] + ", " + line[10] + ", " + line[11] + ");"
				}
				a := insertStr + values
				sliceContainers = append(sliceContainers, a)
			}
			for _, a := range sliceContainers {
				fmt.Println(a)
			}
			fmt.Println("--------------------------------------------------------------------------------------\n")

		// id_bill, id_container, id_customer, bill_date, bill_postalcode, bill_address, bill_city, bill_country, bill_identcode, total_value, weight, volume, , , , , , , Zusätzliche Daten,
		case "bills.csv":
			fmt.Println("--------------------------------------------------------------------------------------")
			fmt.Println("Case Bills\n")

			// INSERT-String Zeile für Zeile bauen
			insertStr := "INSERT INTO Bill(id_bill, id_container, id_user, bill_date, bill_city, bill_country, total_value, total_weight, total_volume) VALUES ("
			for _, line := range lines {
				billDate := line[3][6:10] + "-" + line[3][3:5] + "-" + line[3][0:2]

				values := line[0] + ", " + line[1] + ", " + line[2] + ", '" + billDate + "', '" + line[6] + "', '" + line[7] + "', " + strings.Replace(line[9], ",", ".", 1) + ", " + strings.Replace(line[10], ",", ".", 1) + ", " + strings.Replace(line[11], ",", ".", 1) + ");"
				a := insertStr + values
				sliceBills = append(sliceBills, a)
			}
			for _, a := range sliceBills {
				fmt.Println(a)
			}
			fmt.Println("--------------------------------------------------------------------------------------\n")

		// id_product, id_bill, product_name ,product_cost($) ,product_count, product_color, product_brand, product_category, product_weight(kg), product_volume(L)
		case "products.csv":
			fmt.Println("--------------------------------------------------------------------------------------")
			fmt.Println("Case Products\n")

			// INSERT-String Zeile für Zeile bauen
			insertStr := "INSERT INTO product(id_product, id_bill, product_name, product_value, amount, color, brand, category, weight, volume) VALUES ("
			for _, line := range lines {
				values := line[0] + ", " + line[1] + ", '" + line[2] + "', " + strings.Replace(line[3], ",", ".", 1) + ", " + line[4] + ", '" + line[5] + "', '" + line[6] + "', '" + line[7] + "', " + strings.Replace(line[8], ",", ".", 1) + ", " + strings.Replace(line[9], ",", ".", 1) + ");"
				a := insertStr + values
				sliceProducts = append(sliceProducts, a)
			}
			for _, a := range sliceProducts {
				fmt.Println(a)
			}
			fmt.Println("--------------------------------------------------------------------------------------\n")
		}
	}
}

// Löscht die Dateien nach Übernahme in Datenbank
func deleteFiles() {
	files, error := ioutil.ReadDir(observedDir)
	if error != nil {
		fmt.Println("Error reading directory  \n")
		log.Fatal(error)
	}

	for _, file := range files {
		error = os.Remove(observedDir + "/" + file.Name())
		if error != nil {
			fmt.Println("Error deleting file " + file.Name() + "\n")
			log.Fatal(error)
		}
	}
}

// Verbindet sich zur Datenbank und fügt die INSERTs der Slices ein
func saveToDatabase() {
	dbConn, error := sql.Open("mysql", "root:wasd123@tcp(127.0.0.1:3306)/logistics")

	if error != nil {
		panic(error.Error())
	}

	defer dbConn.Close()

	// INSERTs für Table Customer an DB senden
	for _, insert := range sliceCustomers {
		fmt.Println("Customer INSERT fertig")

		statement, error := dbConn.Query(insert)
		if error != nil {
			panic(error.Error())
		}

		defer statement.Close()
	}

	// INSERTs für Table Container an DB senden
	for _, insert := range sliceContainers {
		fmt.Println("Container INSERT fertig")

		statement, error := dbConn.Query(insert)
		if error != nil {
			panic(error.Error())
		}

		defer statement.Close()
	}

	// INSERTs für Table Bill an DB senden
	for _, insert := range sliceBills {
		fmt.Println("Bill INSERT fertig")

		statement, error := dbConn.Query(insert)
		if error != nil {
			panic(error.Error())
		}

		defer statement.Close()
	}

	// INSERTs für Table Product an DB senden
	for _, insert := range sliceProducts {
		fmt.Println("Product INSERT fertig")

		statement, error := dbConn.Query(insert)
		if error != nil {
			panic(error.Error())
		}

		defer statement.Close()
	}

	fmt.Println("Connection Successful")
}

func main() {
	readFilesAndSaveSlices()
	deleteFiles()
	saveToDatabase()
	fmt.Println("\nProcess completed!")
}
