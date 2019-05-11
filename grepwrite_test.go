package grepwrite_test

import (

	// . "github.com/a-b/grepwrite"
	"fmt"
	"io/ioutil"
	"log"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Grepwrite", func() {
	Describe("Replace corresponding lines", func() {
		var (
			tmpDir string
			err    error
		)

		BeforeEach(func() {
			tmpDir, err = ioutil.TempDir("tmp", "test")
			if err != nil {
				log.Fatal(err)
			}
		})

		Describe("test", func() {
			It("does something", func() {
				inputFile, err := ioutil.TempFile(tmpDir, "input")
				if err != nil {
					log.Fatal(err)
				}
				// defer os.Remove(inputFile.Name())

				content := `
fileA:2:  lINE2
folderA/fileB:3:   newLine3
`

				n, err := inputFile.WriteString("")
				fmt.Printf(content, n)
				inputFile.Sync()

				Expect("test").To(Equal("test"), "annotation %s", tmpDir)

			})
		})

		AfterEach(func() {
			// os.RemoveAll(tmpDir)
		})

	})
})
