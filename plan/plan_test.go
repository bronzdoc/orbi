package plan_test

import (
	"os"
	"os/exec"

	"github.com/bronzdoc/orbi/definition"
	. "github.com/bronzdoc/orbi/plan"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Plan", func() {
	var plan *Plan

	BeforeEach(func() {
		resources := []map[interface{}]interface{}{
			{
				"dir": map[interface{}]interface{}{
					"name": "tmp-dir",
					"dir": map[interface{}]interface{}{
						"name": "tmp-dir-child",
					},
					"files": []interface{}{
						"tmp-file-child",
					},
				},
			},
		}

		map_definition := map[interface{}]interface{}{
			"context":   "./test-resource",
			"resources": resources,
		}

		options := map[string]interface{}{
			"templates_path": "./tmp/templates_path",
		}

		definition := definition.New(map_definition, options)
		plan = New(definition)
	})

	Describe("#Execute", func() {
		It("should execute the plan", func() {
			plan.Execute()
			resources := []string{
				"./test-resource/tmp-dir",
				"./test-resource/tmp-dir/tmp-dir-child",
				"./test-resource/tmp-dir/tmp-file-child",
			}

			for _, resource := range resources {
				resource_exists, _ := exists(resource)
				Expect(resource_exists).To(Equal(true))
			}
		})
	})

	Describe("List", func() {
		It("should get a list of all the plans", func() {
			MockPlanPath()
			plan_list := List()

			Expect(plan_list).To(Equal([]string{
				"plan_a",
				"plan_b",
				"plan_c",
			}))
		})
	})

	Describe("Edit", func() {
		MockPlanPath()
		Context("When plan doesn't exists", func() {
			It("should return the correct error message", func() {
				err := Edit("plan_z")
				Expect(err.Error()).To(Equal(`plan "plan_z" doesn't exists.`))
			})
		})

		Context("When $EDITOR is not set", func() {
			It("should return the correct error message", func() {
				os.Setenv("EDITOR", "")
				err := Edit("plan_a")
				Expect(err.Error()).To(Equal(`$EDITOR is empty, could not edit "plan_a" plan.`))
			})
		})

		Context("When $EDITOR is set correctly", func() {
			It("should return no error", func() {
				defer func() { ExecCommand = exec.Command }()
				ExecCommand = fakeExecCommand

				os.Setenv("EDITOR", "vim")
				err := Edit("plan_a")
				Expect(err).To(BeNil())
			})
		})

	})

	Describe("PlanDefinition", func() {
		It("should return a definition for a new plan", func() {
			MockPlanPath()
			var options map[string]interface{}
			pd := PlanDefinition("plan_x", options)
			resource := pd.Search("plan_x/definition.yml")
			content := resource.(*definition.File).Content()

			Expect(content).To(Equal([]byte(`---
context: .
resources:
  - dir:
     name: dir_name_sample
     files:
      - file_name_sample`),
			))

		})
	})
})

func MockPlanPath() {
	os.MkdirAll("./tmp/.orbi/plans", 0777)
	os.Create("./tmp/.orbi/plans/plan_a")
	os.Create("./tmp/.orbi/plans/plan_a/definition.yml")
	os.Create("./tmp/.orbi/plans/plan_b")
	os.Create("./tmp/.orbi/plans/plan_c")
	os.Setenv("HOME", "./tmp")
}

// See https://npf.io/2015/06/testing-exec-command/
func fakeExecCommand(command string, args ...string) *exec.Cmd {
	cs := []string{"-test.run=TestHelperProcess", "--", command}
	cs = append(cs, args...)
	cmd := exec.Command(os.Args[0], cs...)
	cmd.Env = []string{"GO_WANT_HELPER_PROCESS=1"}
	return cmd
}