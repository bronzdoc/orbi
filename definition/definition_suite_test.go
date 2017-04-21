package definition_test

import (
	"os"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestDefinition(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Definition Suite")
}

// Use ./tmp to put any kinf of files you need to test
// Use ./test-resource to test any resource related funcionality
var _ = BeforeSuite(func() {
	os.Mkdir("./test-resource", 0777)
	os.Mkdir("./tmp", 0777)
})

var _ = AfterSuite(func() {
	os.RemoveAll("./test-resource")
	os.RemoveAll("./tmp")
})
