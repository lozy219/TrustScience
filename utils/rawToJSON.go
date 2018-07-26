package main

import (
	"bufio"
	"encoding/json"
	"os"
	"strconv"
	"strings"
)

func main() {
	data := make(map[string]map[string]int)

	file, err := os.Open("../data/raw")
	checkErr(err)
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		text := scanner.Text()
		fields := strings.Fields(text)
		w, _ := strconv.Atoi(fields[1])
		l, _ := strconv.Atoi(fields[2])
		data[fields[0]] = map[string]int{"winning": w, "losing": l}
	}

	dataJ, err := json.Marshal(data)
	checkErr(err)

	fout, err := os.Create("../frontend/data/data.json")
	checkErr(err)
	defer fout.Close()

	_, ferr := fout.WriteString(string(dataJ))
	checkErr(ferr)
	fout.Sync()
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
