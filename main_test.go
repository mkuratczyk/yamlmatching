package main

import (
	"fmt"

	. "github.com/Benjamintf1/Expanded-Unmarshalled-Matchers"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("YAML generator", func() {
	It("generates the expected YAML", func() {
		var expectedYAML = `
foo:
- name: item1
`

		generatedYAML := generate()
		fmt.Printf("\n--- GENERATED ---\n%v---------\n", generatedYAML)
		fmt.Printf("\n--- EXPECTED ---\n%v---------\n", expectedYAML)

		Expect(MatchUnorderedYAML(expectedYAML).Match(generatedYAML))
		//Expect(generatedYAML).To(Equal(expectedYAML))
	},
	)
})
