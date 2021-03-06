package fate_test

import (
	"github.com/godcong/chronos"
	"github.com/godcong/fate"

	"log"
	"testing"
)

func TestPoint(t *testing.T) {
	t1 := chronos.New("2017/11/11 12:17")
	log.Println(t1.Lunar().EightCharacter())
	log.Printf("%+v", fate.NewBazi(t1).XiYong())
}
