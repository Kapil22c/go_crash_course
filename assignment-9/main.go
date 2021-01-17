package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	"io/ioutil"
)

func main() {

	var option int
	fmt.Println("1.command line interface\n2.Json file\n3.Gin-Gonic")
	fmt.Printf("enter your option :")
	fmt.Scanf("%d", &option)
	fmt.Println("You have entered option", option)
	options := option
	switch options {
	case 1:
		fmt.Println("welcome to command line interface")
		cmd()
	case 2:
		fmt.Println("Upload json file")
		jsonupload()
	case 3:
		fmt.Println("Gin gonic Framework")
		gingonic()
	default:
		fmt.Println("Invalid option")
		break

	}

}

type Cycleslist struct {
	Id   int64  `json:"id"`
	Date string `json:"date"`

	Frame string `json:"frame"`

	Handlebar string `json:"handlebar"`
	Gear      int64  `json:"gear"`
	Geargrip  int64  `json:"geargrip"`

	Seating       int64 `json:"seating"`
	Seatingbottle int64 `json:"seatingbottle"`

	Wheels string `json:"wheels"`
	Spokes int64  `json:"spokes"`
	Rim    int64  `json:"rim"`
	Tube   int64  `json:"tube"`
	Tyre   string `json:"tyre"`

	Chain string `json:"chain"`
}

func cmd() {
	fmt.Println("Command line interface option not availabale")
	fmt.Println("try json upload file")
}

func jsonupload() {
	var obj []Cycleslist
	fmt.Println("json interface")
	data, err := ioutil.ReadFile("example.json")
	if err != nil {
		fmt.Print(err)
	}
	err = json.Unmarshal(data, &obj)
	if err != nil {
		fmt.Println("error:", err)
	}
	for _, x := range obj {
		frameprice, handlebarprice, wheelprice, chainprice, tseatingprice, totalcyclengineprice := Calculateginprice(x.Date, x.Frame, x.Handlebar, x.Gear, x.Geargrip, x.Seating, x.Seatingbottle, x.Wheels, x.Spokes, x.Rim, x.Tube, x.Tyre, x.Chain)
		fmt.Println("the frame price", frameprice)
		fmt.Println("the handlebar price", handlebarprice)
		fmt.Println("the wheel price", wheelprice)
		fmt.Println("the chain price", chainprice)
		fmt.Println("the tseating price", tseatingprice)
		fmt.Println("the price of cycle is", totalcyclengineprice)
	}

}

func gingonic() {
	fmt.Println("Gin-Gonic Starting")

	r := gin.Default()
	r.POST("/books", CalculatePrice)
	r.Run(":8054")
}

func CalculatePrice(c *gin.Context) {

	var cycle Cycleslist
	var user []Cycleslist

	if err := c.ShouldBindJSON(&cycle); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user = append(user, cycle)
	frameprice, handlebarprice, wheelprice, chainprice, tseatingprice, totalcyclengineprice := Calculateginprice(cycle.Date, cycle.Frame, cycle.Handlebar, cycle.Gear, cycle.Geargrip, cycle.Seating, cycle.Seatingbottle, cycle.Wheels, cycle.Spokes, cycle.Rim, cycle.Tube, cycle.Tyre, cycle.Chain)
	c.JSON(http.StatusOK, gin.H{"frameprice": frameprice, "handlebarprice": handlebarprice, "wheelprice": wheelprice, "chainprice": chainprice, "tseatingprice": tseatingprice, "totalprice": totalcyclengineprice})
	fmt.Println(user)
}

func Calculateginprice(date string, frame string, handlebar string, gear int64, geargrip int64, seating int64, seatingbottle int64, wheels string, spokes int64, rim int64, tube int64, tyre string, chain string) (int64, int64, int64, int64, int64, int64) {

	var tyred, frameprice, chainprice, tseatingprice, wheelprice, handlebarprice, totalcyclengineprice int64
	Date, _ := strconv.Atoi(date[0:2])
	Year, _ := strconv.Atoi(date[3:7])
	if Date >= 1 && Date <= 11 && Year == 2016 {
		tyred = 200
	} else {
		tyred = 230
	}

	if frame == "steel" {
		frameprice = 150
	} else if frame == "Aluminium" {
		frameprice = 100
	}

	if handlebar == "shockabsorb" {
		handlebarprice = 200 + (gear * 100) + geargrip
	} else {
		handlebarprice = 100 + (gear * 100) + geargrip
	}
	if seating == 1 {
		tseatingprice = 100 + seatingbottle
	} else {
		tseatingprice = 200 + seatingbottle
	}
	if wheels == "steel" {
		wheelprice = 150 + spokes + rim + tyred + 200 + tube
	} else if wheels == "Aluminium" {
		wheelprice = 200 + spokes + rim + tyred + 100 + tube
	}
	if chain == "onespeed" {
		chainprice = 150
	} else if chain == "twospeed" {
		chainprice = 200
	}

	totalcyclengineprice = frameprice + handlebarprice + wheelprice + chainprice + tseatingprice
	return frameprice, handlebarprice, wheelprice, chainprice, tseatingprice, totalcyclengineprice

}
