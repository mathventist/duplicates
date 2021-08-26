package duplicates

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Duplicates", func() {

	Describe("ScanSentences", func() {
		var data []byte
		var atEOF bool

		var adv int
		var t []byte
		var e error

		JustBeforeEach(func() {
			adv, t, e = ScanSentences(data, atEOF)
		})

		Context("when data is empty", func() {

			Context("when atEOF is true", func() {
				BeforeEach(func() {
					atEOF = true
				})

				It("returns 0, nil, nil", func() {
					Expect(adv).To(Equal(0))
					Expect(t).To(BeNil())
					Expect(e).To(BeNil())
				})
			})

			Context("when atEOF is false", func() {
				BeforeEach(func() {
					atEOF = false
				})

				It("returns 0, nil, nil", func() {
					Expect(adv).To(Equal(0))
					Expect(t).To(BeNil())
					Expect(e).To(BeNil())
				})
			})
		})

		Context("when the data does not contain any sentence terminators", func() {
			BeforeEach(func() {
				data = []byte("hello world  ")
			})

			Context("and atEOF is true", func() {
				BeforeEach(func() {
					atEOF = true
				})

				It("returns the length of the data, the data itself, and nil", func() {
					Expect(adv).To(Equal(13))
					Expect(t).To(Equal([]byte("hello world  ")))
					Expect(e).To(BeNil())
				})
			})

			Context("and atEOF is false", func() {
				BeforeEach(func() {
					atEOF = false
				})

				It("returns 0, nil, nil", func() {
					Expect(adv).To(Equal(0))
					Expect(t).To(BeNil())
					Expect(e).To(BeNil())
				})
			})
		})

		Context("when the data contains a sentence followed by no whitespace", func() {
			BeforeEach(func() {
				data = []byte("sentence one.")
			})

			Context("and atEOF is true", func() {
				BeforeEach(func() {
					atEOF = true
				})

				It("advances the position to the next spot after the end of the sentence", func() {
					Expect(adv).To(Equal(13))
				})

				It("returns the sentence", func() {
					Expect(t).To(Equal([]byte("sentence one.")))
				})

				It("return nil error", func() {
					Expect(e).To(BeNil())
				})
			})

			Context("and atEOF is false", func() {
				BeforeEach(func() {
					atEOF = false
				})

				It("advances the position to the next spot after the end of the sentence", func() {
					Expect(adv).To(Equal(13))
				})

				It("returns the sentence", func() {
					Expect(t).To(Equal([]byte("sentence one.")))
				})

				It("return nil error", func() {
					Expect(e).To(BeNil())
				})
			})
		})

		Context("when the data contains a sentence followed by whitespace", func() {
			BeforeEach(func() {
				data = []byte("sentence one.   ")
			})

			Context("and atEOF is true", func() {
				BeforeEach(func() {
					atEOF = true
				})

				It("advances the position to the next spot after the whitespace", func() {
					Expect(adv).To(Equal(16))
				})

				It("returns the sentence", func() {
					Expect(t).To(Equal([]byte("sentence one.")))
				})

				It("return nil error", func() {
					Expect(e).To(BeNil())
				})
			})

			Context("and atEOF is false", func() {
				BeforeEach(func() {
					atEOF = false
				})

				It("advances the position to the next spot after the whitespace", func() {
					Expect(adv).To(Equal(16))
				})

				It("returns the sentence", func() {
					Expect(t).To(Equal([]byte("sentence one.")))
				})

				It("return nil error", func() {
					Expect(e).To(BeNil())
				})
			})
		})

		Context("when the data contains a sentence ending in a quotation mark", func() {
			BeforeEach(func() {
				data = []byte("sentence one.\"")
			})

			Context("and atEOF is true", func() {
				BeforeEach(func() {
					atEOF = true
				})

				It("advances the position to the next spot after the quotation mark", func() {
					Expect(adv).To(Equal(14))
				})

				It("returns the sentence", func() {
					Expect(t).To(Equal([]byte("sentence one.\"")))
				})

				It("return nil error", func() {
					Expect(e).To(BeNil())
				})
			})

			Context("and atEOF is false", func() {
				BeforeEach(func() {
					atEOF = false
				})

				It("advances the position to the next spot after the quotation mark", func() {
					Expect(adv).To(Equal(14))
				})

				It("returns the sentence", func() {
					Expect(t).To(Equal([]byte("sentence one.\"")))
				})

				It("return nil error", func() {
					Expect(e).To(BeNil())
				})
			})

			Context("followed by whitespace", func() {
				BeforeEach(func() {
					data = []byte("sentence one.\"  ")
				})

				Context("and atEOF is true", func() {
					BeforeEach(func() {
						atEOF = true
					})

					It("advances the position to the next spot after the whitespace", func() {
						Expect(adv).To(Equal(16))
					})

					It("returns the sentence without the whitespace", func() {
						Expect(t).To(Equal([]byte("sentence one.\"")))
					})

					It("return nil error", func() {
						Expect(e).To(BeNil())
					})
				})

				Context("and atEOF is false", func() {
					BeforeEach(func() {
						atEOF = false
					})

					It("advances the position to the next spot after the whitespace", func() {
						Expect(adv).To(Equal(16))
					})

					It("returns the sentence without the whitespace", func() {
						Expect(t).To(Equal([]byte("sentence one.\"")))
					})

					It("return nil error", func() {
						Expect(e).To(BeNil())
					})
				})
			})
		})

		Context("the data starts with 'St.'", func() {
			BeforeEach(func() {
				data = []byte("St. Paul.")
			})

			It("advances past the abbreviation to the position after the actual sentence termintaor", func() {
				Expect(adv).To(Equal(9))
			})

			It("returns the sentence", func() {
				Expect(t).To(Equal([]byte("St. Paul.")))
			})

			It("return nil error", func() {
				Expect(e).To(BeNil())
			})
		})

		Context("the data contains 'St.'", func() {
			BeforeEach(func() {
				data = []byte("Hello St. Paul.")
			})

			It("advances past the abbreviation to the position after the actual sentence termintaor", func() {
				Expect(adv).To(Equal(15))
			})

			It("returns the sentence", func() {
				Expect(t).To(Equal([]byte("Hello St. Paul.")))
			})

			It("return nil error", func() {
				Expect(e).To(BeNil())
			})
		})
	})

	Describe("Preprocess", func() {
		var output string

		var input string
		var removeStops bool

		JustBeforeEach(func() {
			output = Preprocess(input, removeStops)
		})

		Context("when input contains 'St.'", func() {
			BeforeEach(func() {
				input = "St. Paul and St. Peter"
			})

			It("removes all occurrences from the input", func() {
				Expect(output).To(Equal("paul and peter"))
			})
		})

		Context("when the input contains the ligature æ", func() {
			BeforeEach(func() {
				input = "hail cæsar, cæsar is the best"
			})

			It("expands all occurrences to ae", func() {
				Expect(output).To(Equal("hail caesar caesar is the best"))
			})
		})

		Context("when the input contains hyphens", func() {
			BeforeEach(func() {
				input = "thirty-eight times forty-two"
			})

			It("replaces all occurrences with a blank space", func() {
				Expect(output).To(Equal("thirty eight times forty two"))
			})
		})

		Context("when the input contains commas", func() {
			BeforeEach(func() {
				input = "Paul, and, Peter"
			})

			It("removes all occurrences from the input", func() {
				Expect(output).To(Equal("paul and peter"))
			})
		})

		Context("when the input contains single quotes", func() {
			BeforeEach(func() {
				input = "Paul 'and' 'Peter'"
			})

			It("removes all occurrences from the input", func() {
				Expect(output).To(Equal("paul and peter"))
			})
		})

		Context("when the input contains double quotes", func() {
			BeforeEach(func() {
				input = "Paul \"and\" \"Peter\""
			})

			It("removes all occurrences from the input", func() {
				Expect(output).To(Equal("paul and peter"))
			})
		})

		Context("when the input contains upper case letters", func() {
			BeforeEach(func() {
				input = "This has UpperCase letters"
			})

			It("converts the entire input to lowercase", func() {
				Expect(output).To(Equal("this has uppercase letters"))
			})
		})

		Context("when removeStops is enabled", func() {
			BeforeEach(func() {
				removeStops = true
			})

			Context("and the first word in the string is a stop word", func() {
				BeforeEach(func() {
					input = "A one way street"
				})

				It("removes the leading stop word", func() {
					Expect(output).To(Equal("one way street"))
				})
			})

			Context("and an interior word in the string is a stop word", func() {
				BeforeEach(func() {
					input = "one same street"
				})

				It("removes the interior stop word", func() {
					Expect(output).To(Equal("one street"))
				})
			})
		})
	})

})
