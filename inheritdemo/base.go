/*
  Base is a Base interface and its implementation for simple methods
*/

package inheritdemo

import (
	"log"
)

type Base interface {
	DoIt(str string) (string, error)
}

type Baser struct {
}

func (b *Baser) DoIt(str string) (string, error) {
	log.Println("Base.DoIt()")
	return str, nil
}
