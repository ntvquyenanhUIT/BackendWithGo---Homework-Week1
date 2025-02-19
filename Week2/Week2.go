package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"reflect"
	"strconv"
	"strings"
	"time"
	"unicode/utf8"
)

type Human struct {
	fullName    string
	occupation  string
	yearOfBirth int
}

func (v Human) getAge() int {
	return time.Now().Year() - v.yearOfBirth
}

func (v Human) isSuitableForCurrentJob() bool {
	numOfCharacter := utf8.RuneCountInString(v.fullName)
	if v.yearOfBirth%numOfCharacter == 0 {
		return true
	}
	return false
}

func (v Human) createHuman(data string) Human {
	dataArray := [3]string{}
	dataIndex := 0
	lastIndex := 0
	for i := 0; i < len(data); i++ {
		// Storing 3 different datas :Name, Job, and Year Of Birth into an array of 3 strings
		// the times that '|' has appeared

		if data[i] == '|' {
			dataArray[dataIndex] = data[lastIndex:i]
			dataIndex += 1
			lastIndex = i + 1
		}
		if dataIndex == 2 {
			dataArray[dataIndex] = data[lastIndex:]
			break
		}
	}
	v.fullName = dataArray[0]
	v.occupation = dataArray[1]
	yearOfBirthStr := strings.TrimSpace(dataArray[2])
	v.yearOfBirth, _ = strconv.Atoi(yearOfBirthStr)
	return v
}

func Exercise1(reader *bufio.Reader) {
	human := Human{}
	fmt.Println("Nhập tên: ")
	human.fullName, _ = reader.ReadString('\n')
	human.fullName = strings.TrimSpace(human.fullName)

	fmt.Println("Nhập năm sinh: ")
	yearOfBirthStr, _ := reader.ReadString('\n')
	yearOfBirthStr = strings.TrimSpace(yearOfBirthStr)
	human.yearOfBirth, _ = strconv.Atoi(yearOfBirthStr)

	fmt.Println("Nhập nghề nghiệp: ")
	human.occupation, _ = reader.ReadString('\n')
	human.occupation = strings.TrimSpace(human.occupation)

	if !human.isSuitableForCurrentJob() {
		fmt.Println("Bạn không phù hợp với nghề nghiệp hiện tại")
	} else {
		fmt.Println("Bạn phù hợp với nghề nghiệp của mình")

	}
}

func Exercise2(str string) {
	stringMap := make(map[string]int)
	fmt.Println(reflect.TypeOf(str))
	for i := 0; i < len(str); i++ {
		fmt.Println(reflect.TypeOf(str[i]))
		stringMap[string(str[i])]++
	}
	fmt.Println(stringMap)
}

func Exercise3(nums []int, target int) []int {

	numMap := make(map[int]int)
	for i, num := range nums {
		complement := target - num
		if index, found := numMap[complement]; found {
			return []int{index, i}
		} else {
			numMap[complement] = i
		}
	}
	return nil
}

func Exercise4() {
	// Process the file
	// Create a newreader object
	// Process as you want

	sliceOfHumans := []Human{}

	file, err := os.Open("text.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	reader := bufio.NewReader(file)
	for {
		line, err := reader.ReadString('\n')
		if err != nil {
			break
		}
		newHuman := Human{}
		newHuman = newHuman.createHuman(line)
		sliceOfHumans = append(sliceOfHumans, newHuman)

	}
	fmt.Println(sliceOfHumans)
}

func main() {
	//reader := bufio.NewReader(os.Stdin)
	//Exercise1(reader)
	//Exercise2("Hello")
	Exercise4()
}
