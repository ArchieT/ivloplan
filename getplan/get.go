package getplan

import (
	"bytes"
	"fmt"
	//	"io"
	"io/ioutil"
	//"log"
	"net/http"
	//	"net/url"
	"os"
	"regexp"
	//	"strings"
)

func mergebytesandgivestring(b [][]byte) string {
	var c bytes.Buffer
	for _, j := range b {
		c.Write(j)
	}
	return c.String()
}

func GetInfo() (string, []string) {
	we, err := http.Get(`http://loiv.torun.pl/index.php/pl/dla-uczniow/organizacja-zajec/plan-lekcji`)
	if err != nil {
		//log.Fatal(err)
	}
	replan := regexp.MustCompile(`attachments.{6,26}plan_KLAS.pdf`)
	rezast := regexp.MustCompile(`attachments.{6,26}ast.{0,6}pstwa.{0,15}pdf`)
	//ourw := make([]byte, 0, 100)
	ourw, werr := ioutil.ReadAll(we.Body)
	//wn, werr := we.Body.Read(ourp)
	if werr != nil {
		//log.Fatal(string(werr.Error()))
	}
	wplan := replan.FindAll(ourw, -1)
	wzast := rezast.FindAll(ourw, -1)
	fmt.Println(mergebytesandgivestring(wplan))
	fmt.Println(mergebytesandgivestring(wzast))
	if len(wplan) > 1 {
		//log.Fatal(wplan)
		fmt.Println("Too long!")
	}
	oz := make([]string, 0, 10)
	for i := range wzast {
		oz = append(oz, string(wzast[i]))
	}
	return string(wplan[0]), oz
}

func GetPlan(u string) []byte {
	we, err := http.Get(`http://loiv.torun.pl/` + u)
	if err != nil {
		//log.Fatal(string(err.Error()))
	}
	naszplik, err := ioutil.ReadAll(we.Body)
	if err != nil {
		//log.Fatal(string(err.Error()))
	}
	return naszplik
}

func SaveZasts(ul []string) {
	//out := make([][]byte, 0, len(ul))
	for i := range ul {
		//out = append(out, GetPlan(ul[i]))
		ZapiszPlik(GetPlan(ul[i]), ul[i])
	}
	//return out
}

//func SavePlan

func findByte(s string, char byte) int {
	bs := []byte(s)
	last := 0
	for i := range bs {
		if bs[i] == char {
			last = i
			fmt.Println(string(bs[i+1:]))
		}
	}
	return last + 1
}

func DajNazwePliku(s string) string {
	last := findByte(s, []byte(`/`)[0])
	a := string([]byte(s)[last:])
	fmt.Println(a)
	return a
}

func ZapiszPlik(cozap []byte, url string) {
	nazwa := DajNazwePliku(url)
	_, err := os.Stat(`./` + nazwa)
	if err == nil {
		fileopen, openerr := os.Open(`./` + nazwa)
		if openerr != nil {
			//log.Fatal(string(openerr.Error()))
		}
		fileread, readallerr := ioutil.ReadAll(fileopen)
		if readallerr != nil {
			//log.Fatal(string(readallerr.Error))
		}
		if bytes.Equal(cozap, fileread) {
			return
		}
	} else {
		os.Create(`./` + nazwa)
	}
	fmt.Println(err)
	err = ioutil.WriteFile(`./`+nazwa, cozap, 0644)
	fmt.Println(err)
	return
}

func FullService() {
	pi, zi := GetInfo()
	fmt.Println(1)
	p := GetPlan(pi)
	fmt.Println(2)
	ZapiszPlik(p, pi)
	fmt.Println(3)
	SaveZasts(zi)
	fmt.Println(4)
}
