package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strconv"
)

const comicUrl = "https://xkcd.com/%sinfo.0.json"

func main() {
	dict := make(map[int]*Comics)
	input := bufio.NewScanner(os.Stdin)

	for input.Scan() {
		num, err := strconv.Atoi(input.Text())

		if err != nil {
			fmt.Printf("This num is valid\n")
			continue
		}

		if num < 0 {
			break
		}

		if comic, ok := dict[num]; ok {
			fmt.Printf("You search Comic is(from dict):\n%v\n", comic)
		} else {
			comic, err = getCurrentComic(num)
			if err != nil {
				fmt.Printf("Request No.%d Comic Failed: %s\n", num, err.Error())
			} else {
				fmt.Printf("You search Comic is(from site):\n%#v\n", comic)
				dict[num] = comic
			}
		}
	}

}

func getCurrentComic(num int) (*Comics, error)  {
	var requestUrl string
	if num > 0 {
		strNum := strconv.Itoa(num)
		requestUrl = fmt.Sprintf(comicUrl,  strNum + "/")
	} else {
		requestUrl = fmt.Sprintf(comicUrl,  "")
	}

	resp, err := http.Get(requestUrl)

	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		resp.Body.Close()
		return nil, fmt.Errorf("Request comic failed: %s", resp.Status)
	}

	var comicInfo Comics
	if err = json.NewDecoder(resp.Body).Decode(&comicInfo); err != nil {
		resp.Body.Close()
		return nil, err
	}
	resp.Body.Close()
	return &comicInfo, nil
}
