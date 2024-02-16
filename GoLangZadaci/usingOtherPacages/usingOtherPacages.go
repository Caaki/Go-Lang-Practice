package usingOtherPacages

// we called go get github.com/wagslane/go-tinytime to add the tinytime pacage
import (
	"fmt"
	tTime "github.com/wagslane/go-tinytime" // i gave the tTime name we can put it as we please
	"time"
)

func main() {
	tt := tTime.New(1585750374)

	tt = tt.Add(time.Hour * 48)
	fmt.Println(tt)
}
