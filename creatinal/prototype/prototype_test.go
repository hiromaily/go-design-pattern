package prototype_test

import (
	"fmt"
	. "github.com/hiromaily/go-design-pattern/creatinal/prototype"
	tu "github.com/hiromaily/golibs/testutil"
	"testing"
)

//-----------------------------------------------------------------------------
// Test Framework
//-----------------------------------------------------------------------------
// Initialize
func init() {
	tu.InitializeTest("[Prototype]")
}

//-----------------------------------------------------------------------------
// Test
//-----------------------------------------------------------------------------
func TestPrototype(t *testing.T) {
	product := Product{Name: "BaseObject"}
	product.SetUp()

	manager := Manager{}
	manager.Register(&product)

	cloned := manager.Create("ClonedObject")

	fmt.Println("product.GetName():", product.GetName())
	fmt.Println("cloned.GetName():", cloned.GetName())

	if cloned.GetName() != product.GetName() {
		t.Errorf("Expect instance to equal, but not equal.")
	}
}
