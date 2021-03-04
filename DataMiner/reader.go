package main

import "fmt"
import "io/ioutil"
import "os"

var observed_dir = "C:/Users/chris/OneDrive/Dokumente/Diplomarbeit/ReWriteGo/csv-files"

func transferFiles() {

	//------------------------------------------------------------------
	// Dir aus dem CSV-Files bezogen werden
	files, error := ioutil.ReadDir(observed_dir)
	if error != nil {
		fmt.Println("Error reading directory  \n")
		fmt.Println(error)
	}
	//------------------------------------------------------------------

	// Jedes File lesen und in neues Dir kopieren

	for _, file := range files {
		fmt.Println(file.Name() + "\n")
		data, error := ioutil.ReadFile(observed_dir + "/" + file.Name())
		if error != nil {
			fmt.Println("Error reading file " + file.Name() + "\n")
			fmt.Println(error)
		}

		error = os.Remove(observed_dir + "/" + file.Name())
		if error != nil {
			//log.Fatal(error)
			fmt.Println("Error deleting file " + file.Name() + "\n")
			fmt.Println(error)
		}

		error = ioutil.WriteFile("DataMiner/input/"+file.Name(), data, 0644)
		if error != nil {
			fmt.Println("Error writing file " + file.Name() + "\n")
			fmt.Println(error)
		}

	}

}

func main() {
	transferFiles()
	fmt.Println("Completed Filetransfer!")
}
