package singleton_test

import (
	"fmt"
	. "github.com/hiromaily/go-design-pattern/creatinal/singleton"
	tu "github.com/hiromaily/golibs/testutil"
	"testing"
	"time"
)

//-----------------------------------------------------------------------------
// Test Framework
//-----------------------------------------------------------------------------
// Initialize
func init() {
	tu.InitializeTest("[Singleton]")
}

//-----------------------------------------------------------------------------
// Test
//-----------------------------------------------------------------------------
func TestSingleton(t *testing.T) {
	ins := Get()

	defer func() {
		//stop
		fmt.Println("[Debug]Defer: `ins.Controller <- Stop`")
		ins.Controller <- Stop //stop timer

		fmt.Println("[Debug]Defer: `close(ins.Controller)`")
		close(ins.Controller)

		fmt.Println("[Debug]Defer: `close(ins.Receiver)`")
		close(ins.Receiver)
	}()

	ins.AddNumber()

	ins.AddNumber()

	ins.ReduceNumber()

	//timer 3s
	fmt.Println(" -> sleep 3s")
	time.Sleep(3 * time.Second)

	//receive
	result := ins.GetNumber()
	if result != 2 {
		t.Errorf("ins.GetNumber() return %d, expected value: %d", result, 2)
	}
}
