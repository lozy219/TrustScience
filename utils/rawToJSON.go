package main

import (
	"bufio"
	"encoding/json"
	"os"
	"strconv"
	"strings"
)

func main() {
	result := make(map[string]map[string]map[string]int)

	targetFolder := "../data/raw/"
	dir, err := os.Open(targetFolder)
	checkErr(err)
	defer dir.Close()

	list, _ := dir.Readdirnames(0)
	for _, name := range list {
		data := make(map[string]map[string]int)
		file, err := os.Open(targetFolder + name)
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

		result[name] = data
	}

	dataJ, err := json.Marshal(result)
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
