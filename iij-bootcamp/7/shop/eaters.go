package shop

import (
	"fmt"
	"net/http"
	"time"
)

type Gyudon struct {
	Menu string
}

func NewGyudon() Gyudon {
	return Gyudon{
		Menu: "NegitamaGyudon",
	}
}

func (self *Gyudon) Eat(w http.ResponseWriter, r *http.Request) {
	if self.Menu == "" {
		return
	}

	time.Sleep(time.Second * 10) //擬似食べてる時間
	fmt.Fprintf(w, "'%s'\n", self.Menu)
	return
}
