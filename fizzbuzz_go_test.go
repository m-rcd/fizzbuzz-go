package fizzbuzz_go_test

import (
	"github.com/m-rcd/fizzbuzz-go"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Fizzbuzz", func() {
	It("returns true if divisible by 3", func() {
		Expect(fizzbuzz_go.IsDivisibleBy(3, 3)).To(Equal(true))
	})

	It("returns false if not divisble by 3", func() {
		Expect(fizzbuzz_go.IsDivisibleBy(5, 3)).To(Equal(false))
	})

	It("returns true if divisble by 5", func() {
		Expect(fizzbuzz_go.IsDivisibleBy(5, 5)).To(Equal(true))
	})

	It("returns false if not divisible by 5", func() {
		Expect(fizzbuzz_go.IsDivisibleBy(1, 5)).To(Equal(false))
	})

	It("returns true if divisible by 15", func() {
		Expect(fizzbuzz_go.IsDivisibleBy(15, 15)).To(Equal(true))
	})

	It("returns false if not divisible by 15", func() {
		Expect(fizzbuzz_go.IsDivisibleBy(20, 15)).To(Equal(false))
	})

	It("says fizz if divisible by 3", func() {
		Expect(fizzbuzz_go.FizzbuzzSays(3)).To(Equal("fizz"))
	})

	It("says buzz if divisible by 5", func() {
		Expect(fizzbuzz_go.FizzbuzzSays(5)).To(Equal("buzz"))
	})

	It("says fizzbuzz if divisible by 15", func() {
		Expect(fizzbuzz_go.FizzbuzzSays(15)).To(Equal("fizzbuzz"))
	})

	It("returns number if not divisible by 3, 5 or 15", func() {
		Expect(fizzbuzz_go.FizzbuzzSays(4)).To(Equal("4"))
	})
})
