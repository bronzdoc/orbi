package definition_test

import (
	"io/ioutil"
	"log"

	. "github.com/bronzdoc/orbi/definition"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Resource", func() {
	Describe("Directory", func() {
		var directory *Directory

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

	Describe("File", func() {
		var file *File

		BeforeEach(func() {
			file = NewFile(
				"tmp-file",
				"./test-resource/tmp-file",
				[]byte("Oooh, that's a bingo!"),
			)
		})

		Describe("#Name", func() {
			It("should return the correct name", func() {
				Expect(file.Name()).To(Equal("tmp-file"))
			})
		})

		Describe("#Id", func() {
			It("should return the correct id", func() {
				Expect(file.Id()).To(Equal("./test-resource/tmp-file"))
			})
		})

		Describe("#Create", func() {
			It("should create a file", func() {
				file.Create(map[string]interface{}{})

				file_exists, err := exists(file.Id())
				if err != nil {
					// Todo  change to Fail(err)
					log.Fatal(err)
				}

				Expect(file_exists).To(Equal(true))

				content, err := ioutil.ReadFile("./test-resource/tmp-file")
				if err != nil {
					log.Fatal(err)
				}

				Expect(content).To(Equal([]byte("Oooh, that's a bingo!")))
			})
		})

		Describe("#Children", func() {
			It("should return the correct children", func() {
				var empty_resource []Resource
				children := file.Children()

				Expect(len(children)).To(Equal(0))
				Expect(children).To(Equal(empty_resource))
			})
		})
	})
})
