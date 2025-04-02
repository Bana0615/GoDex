package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

//Run `go run .`

func main() {
	for {
		scanner := bufio.NewScanner(os.Stdin)
		fmt.Println("#################### Enter your favorite Pokemon! ####################")
		fmt.Print("> ")
		scanner.Scan()

		url := fmt.Sprintf("https://pokeapi.co/api/v2/pokemon/%s", scanner.Text())

		res, err := http.Get(url)
		if err != nil {
			panic(err)
		}

		defer res.Body.Close()

		body, err := io.ReadAll(res.Body)
		if err != nil {
			panic(err)
		}

		mon := Mon{}
		err = json.Unmarshal(body, &mon)
		if err != nil {
			panic(err)
		}

		fmt.Println("---")
		fmt.Println(mon.Name)

		for x := range mon.Stats {
			fmt.Println(mon.Stats[x].Stat.Name, mon.Stats[x].BaseStat)
		}
		fmt.Println()
	}
}
