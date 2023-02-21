package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

var ALLOCATION = [5][4]int{{0, 0, 0, 0}, {0, 0, 0, 0}, {0, 0, 0, 0}, {0, 0, 0, 0}, {0, 0, 0, 0}}
var NEEDS = [5][4]int{{0, 0, 0, 0}, {0, 0, 0, 0}, {0, 0, 0, 0}, {0, 0, 0, 0}, {0, 0, 0, 0}}
var AVALIABLE [][]int
var MAX [5][4]int

// function that parses a given string array to int array
func to_int_arr(arr []string) []int {
	var int_arr []int
	for i := 0; i < len(arr); i++ {
		ie, err := strconv.Atoi(arr[i])
		int_arr = append(int_arr, ie)
		if err != nil {
			log.Fatal(err)
		}
	}
	return int_arr
}

func update_need() {
	for i := 0; i < 5; i++ {
		for j := 0; j < 4; j++ {
			NEEDS[i][j] = MAX[i][j] - ALLOCATION[i][j]
		}
	}
}

func main() {
	// reading max resources used by each process from txt file
	f, err := os.Open("max.txt")
	if err != nil {
		log.Fatal(err)
	}
	scanner := bufio.NewScanner(f)
	row := 0
	for scanner.Scan() {
		res := strings.Split(scanner.Text(), ",")
		for i := 0; i < len(res); i++ {
			temp, err := strconv.Atoi(res[i])
			MAX[row][i] = temp
			if err != nil {
				log.Fatal(err)
			}
		}
		row++
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	// accepting avaliable amount of resources from user
	println("Enter the number of avaliable resources : ")
	input := bufio.NewScanner(os.Stdin)
	input.Scan()
	sinput := input.Text()
	sin := strings.Split(sinput, " ")
	AVALIABLE = append(AVALIABLE, to_int_arr(sin))

	// accepting resource requests
	for {
		print("->> ")
		input := bufio.NewScanner(os.Stdin)
		input.Scan()
		command := input.Text()
		switch command {
		case "*":
			// printing MAX and ALLOCATION
			print("MAX = ")
			for i := 0; i < len(MAX); i++ {
				for j := 0; j < len(MAX[0]); j++ {
					print(MAX[i][j], ",")
				}
				print(" ")
			}
			println()
			print("ALLOCATION = ")
			for i := 0; i < len(ALLOCATION); i++ {
				for j := 0; j < len(ALLOCATION[0]); j++ {
					print(ALLOCATION[i][j], ",")
				}
				print(" ")
			}
			println()
			print("NEEDS = ")
			for i := 0; i < len(NEEDS); i++ {
				for j := 0; j < len(NEEDS[0]); j++ {
					print(NEEDS[i][j], ",")
				}
				print(" ")
			}
			println()
			print("AVALIABLE = ")
			for i := 0; i < len(AVALIABLE); i++ {
				for j := 0; j < 4; j++ {
					print(AVALIABLE[i][j], ",")
				}
			}
			println()
		case "help":
			println("* - to print all the data structures \n RQ P A B C D - to request resources for process P \n RL P A B C D - to release resources held by process P \n help - to get the commands \n exit - to exit the program")
		case "exit":
			os.Exit(0)
		default:
			cmd := strings.Split(command, " ")
			if cmd[0] == "RQ" {
				println("requesting resource..")
				r, err := strconv.Atoi(cmd[1])
				rs := to_int_arr(cmd[2:])

				// checking if a process is asking more than permitted, if true notify user if not continue
				if rs[0] <= MAX[r][0] && rs[1] <= MAX[r][1] && rs[2] <= MAX[r][2] && rs[3] <= MAX[r][3] {
					// checking if a process is asking more than avaliable resource, if true notify user if not update data structures
					if rs[0] <= AVALIABLE[0][0] && rs[1] <= AVALIABLE[0][1] && rs[2] <= AVALIABLE[0][2] && rs[3] <= AVALIABLE[0][3] {
						println("Updating data structures..")
						for k := 0; k < 4; k++ {
							ALLOCATION[r][k] = ALLOCATION[r][k] + rs[k]
							AVALIABLE[0][k] = AVALIABLE[0][k] - rs[k]
						}
						update_need()
						println("Done!!")
					} else {
						println("Process asking more than the avaliable resources!!")
					}
				} else {
					println("Process asking more than listed in the MAX data structure!!")
				}
				if err != nil {
					log.Fatal(err)
				}
			} else if cmd[0] == "RL" {
				// handling relase resource
				println("releasing resource..")
				r, err := strconv.Atoi(cmd[1])
				rs := to_int_arr(cmd[2:])

				// check if the process has allocated the requestsed amount of resource or not, if true release resources else notify user
				if rs[0] <= ALLOCATION[r][0] && rs[1] <= ALLOCATION[r][1] && rs[2] <= ALLOCATION[r][2] && rs[3] <= ALLOCATION[r][3] {
					println("Updating data structures..")
					for k := 0; k < 4; k++ {
						ALLOCATION[r][k] = ALLOCATION[r][k] - rs[k]
						AVALIABLE[0][k] = AVALIABLE[0][k] + rs[0]
					}
					update_need()
					println("Done!!")
				} else {
					println("Process didn't allocate this much resources, please try again!!")
				}
				// handling handling errors
				if err != nil {
					log.Fatal(err)
				}
			} else {
				println("Invalid command, please refer by using help command!!")
			}
		}
	}
}
