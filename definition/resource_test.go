package definition_test

import (
	"log"
	"os"

	. "github.com/bronzdoc/orbi/definition"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Resource", func() {
	var directory *Directory

	Describe("Directory", func() {
		BeforeSuite(func() {
			os.Mkdir("./test-resource", 0777)
		})

		BeforeEach(func() {
			directory = NewDirectory(
				"tmp",
				"./test-resource/tmp",
				[]Resource{
					NewDirectory(
						"tmp-child",
						"./test-resource/tmp-child",
						nil,
					),
				},
			)
		})

		AfterSuite(func() {
			os.Remove("./test-resource")
		})

		Describe("#Name", func() {
			It("should return the correct name", func() {
				Expect(directory.Name()).To(Equal("tmp"))
			})
		})

		Describe("#Id", func() {
			It("should return the correct id", func() {
				Expect(directory.Id()).To(Equal("./test-resource/tmp"))
			})
		})

		Describe("#Create", func() {
			It("should create a directory", func() {
				directory.Create(map[string]interface{}{})
				dir_exists, err := exists(directory.Id())
				if err != nil {
					log.Fatal(err)
				}
				Expect(dir_exists).To(Equal(true))
			})
		})

		Describe("#Children", func() {
			It("should return the correct children", func() {
				children := directory.Children()
				Expect(len(children)).To(Equal(1))
				Expect(children[0].Name()).To(Equal("tmp-child"))
				Expect(children[0].Id()).To(Equal("./test-resource/tmp-child"))
			})
		})
	})

})

func exists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return true, err
}
