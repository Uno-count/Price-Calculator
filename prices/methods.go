package prices

import (
	"bufio"
	"fmt"
	"os"
)

func (job TaxIncludedPriceJob) Process() {
	result := make(map[string]float64)

	for _, price := range job.InputPrices {
		result[fmt.Sprintf("%.2f", price)] = price * (1 + job.TaxRate)
	}

	fmt.Println(result)
}

func (job TaxIncludedPriceJob) LoadData() {
	file, err := os.Open("data.txt")

	if err != nil {
		fmt.Println("An error occured")
		fmt.Println(err)

		return
	}

	scanner := bufio.NewScanner(file)

	var lines []string

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	err = scanner.Err()
	if err != nil {
		fmt.Println("Could not open file!")
		fmt.Println(err)
		file.Close()
		return
	}
}
