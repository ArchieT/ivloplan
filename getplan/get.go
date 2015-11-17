package getplan

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"regexp"
	"strings"
)

func GetInfo() (string, []string) {
	we, err := http.Get(`http://loiv.torun.pl/index.php/pl/dla-uczniow/organizacja-zajec/plan-lekcji`)
	if err != nil {
		log.Fatal(err)
	}
	replan := regexp.MustCompile(`attachments.*plan_KLAS.pdf`)
	rezast := regexp.MustCompile(`attachments.*ast.*pstwa.*pdf`)
	ourw := make([]byte, 0, 100)
	wn, werr := we.Body.Read(ourp)
	if werr != nil {
		log.Fatal(perr)
	}
	wplan := replan.FindAll(ourw, -1)
	wzast := rezast.FindAll(ourw, -1)
	if len(wplan) > 0 {
		log.Fatal(wplan)
		fmt.Println("Too long!")
	}
	oz := make([]string, 0, 10)
	for i := range wzast {
		oz = append(oz, wzast[i])
	}
	return string(wplan[0]), oz
}
