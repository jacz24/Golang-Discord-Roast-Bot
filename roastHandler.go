package main

import (
	"bufio"
	"bytes"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strings"
	"time"
)

var splitRoastArray []string

func createRoast() string{ // Starts the creation of the roast
	file, err := os.Open("roast.txt")
	if err != nil {
		fmt.Println(err)}
	defer file.Close()

	var strArray string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() { // interates through text file creating an array line by line
		buf := bytes.Buffer{}
		buf.WriteString(scanner.Text())
		strArray += "," + buf.String()}
	if err := scanner.Err(); err != nil {
		fmt.Println("Scanner Error Reading Text File!", err)}
	splitRoastArray := strings.Split(strArray, ",")
	roast := splitRoastArray[random(1, len(splitRoastArray))] // Starts at 1 because 0 is a "" entry

	roast = combineRoast(createAdjective(), roast)

	readRoast(roast)
	return roast
}

func random(min, max int) int {
	rand.Seed(time.Now().Unix())
	return rand.Intn(max - min) + min}

func readRoast(roast string) { // Prints the roast to console mainly for debugging
	log.Println("The roast created was!",roast)}

func readRoastArray(roastArray []string){
	log.Println("Output of roast.txt")
	for i:=0; i <= len(roastArray)-1; i++{
		log.Println(roastArray[i])}}

func combineRoast(setup string, punchline string) string {
	return setup + " " + punchline
}

func createAdjective() string {
	file, err := os.Open("adjective.txt")
	if err != nil {
		fmt.Println(err)}
	defer file.Close()

	var strArray string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() { // interates through text file creating an array line by line
		buf := bytes.Buffer{}
		buf.WriteString(scanner.Text())
		strArray += "," + buf.String()}
	if err := scanner.Err(); err != nil {
		fmt.Println("Scanner Error Reading Text File!", err)}
	splitRoastArray := strings.Split(strArray, ",")

	adjective := splitRoastArray[random(1, len(splitRoastArray))] // Starts at 1 because 0 is a "" entry

	return adjective
}