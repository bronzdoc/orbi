package definition_test

import (
	"io/ioutil"
	"os"

	. "github.com/bronzdoc/orbi/definition"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Template", func() {
	var template *Template

	BeforeEach(func() {
		content := []byte("Antonio Marga-reeeeeiiiiiiiiti.")
		var vars map[string]string
		template = NewTemplate("template_0", content, vars)
	})

	Describe("#Execute", func() {
		It("Should generate a template with the correct content", func() {
			file, _ := os.Create("./tmp/some-file")
			template.Execute(file)

			content, _ := ioutil.ReadFile("./tmp/some-file")
			Expect(string(content)).To(Equal("Antonio Marga-reeeeeiiiiiiiiti."))
		})
	})
})
