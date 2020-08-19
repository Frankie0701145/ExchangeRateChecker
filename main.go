package main

import (
	"fmt"
	"flag"
	"net/http" 
	"log"
	"encoding/csv"
	"io"
)

func main(){
	currencyCode := flag.String("currencyCode", "", "Please pass the ISO 4217 currency code.")
	flag.Parse()
	fmt.Println(*currencyCode)
	answer := CheckCurrencySupport(*currencyCode)
	fmt.Println(answer)
}



func CheckCurrencySupport ( currencyCode string ) string{

	//fetch the csv file
	res,err := http.Get("https://focusmobile-interview-materials.s3.eu-west-3.amazonaws.com/Cheap.Stocks.Internationalization.Currencies.csv")

	//log any error if any
	if err !=nil{
		log.Fatal(err)
	}
	
	//defer closing the res.Body after aat the end of the function
	defer res.Body.Close()

	//Parse the file 
	r := csv.NewReader(res.Body)
	//variable to hold answer to whether the currency code is supported
	var answer string

	//Iterate through the record
	for index:=0; true; index++ {
		//skip the first record which contains the field names
		if index==0 {
			continue;
		}
		// Read each record from csv
		record, err := r.Read()
		//check if there is a record 
		if err == io.EOF  {
			//if there is no more record break the loop and return a message stating the currency is not supported
			answer = fmt.Sprintf("This %s currency code is not supported.", currencyCode)
			break;
		}
		//check if the current record  matches the passed currency code
		if (currencyCode == record[2]) {
			//if they match return a message stating the currency are matching
			answer = fmt.Sprintf("This %s currency code is supported.", currencyCode)	
			break;
		}	
	}
	return answer		
}



