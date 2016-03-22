package hebrewtime_test

import (
	. "github.com/nscarlson/multi-calendar/hebrew"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Hebrew", func() {
	Describe("Leap Years", func() {
		Context("Identify leap years in Hillelian cycle", func() {
			It("is 3rd year in Hillelian cycle and leap year in 5760", func() {
				Expect(IsLeap(5760)).To(BeTrue())
			})
			It("is 6th year in Hillelian cycle and leap year in 5763", func() {
				Expect(IsLeap(5763)).To(BeTrue())
			})
			It("is 8th year in Hillelian cycle and leap year in 5765", func() {
				Expect(IsLeap(5765)).To(BeTrue())
			})
			It("is 11th year in Hillelian cycle and leap year in 5768", func() {
				Expect(IsLeap(5768)).To(BeTrue())
			})
			It("is 14th year in Hillelian cycle and leap year in 5771", func() {
				Expect(IsLeap(5771)).To(BeTrue())
			})
			It("is 17th year in Hillelian cycle and leap year in 5774", func() {
				Expect(IsLeap(5774)).To(BeTrue())
			})
			It("is 19th year in Hillelian cycle and leap year in 5776", func() {
				Expect(IsLeap(5776)).To(BeTrue())
			})
		})

		Context("Identify non-leap years in Hillelian cycle", func() {
			It("is 1st year in Hillelian cycle and is not a leap year in 5758", func() {
				Expect(IsLeap(5758)).To(BeFalse())
			})
			It("is 2nd year in Hillelian cycle and is not a leap year in 5759", func() {
				Expect(IsLeap(5759)).To(BeFalse())
			})
			It("is 4th year in Hillelian cycle and is not a leap year in 5761", func() {
				Expect(IsLeap(5761)).To(BeFalse())
			})
			It("is 5th year in Hillelian cycle and is not a leap year in 5762", func() {
				Expect(IsLeap(5762)).To(BeFalse())
			})
			It("is 7th year in Hillelian cycle and is not a leap year in 5764", func() {
				Expect(IsLeap(5764)).To(BeFalse())
			})
			It("is 9th year in Hillelian cycle and is not a leap year in 5766", func() {
				Expect(IsLeap(5766)).To(BeFalse())
			})
			It("is 10th year in Hillelian cycle and is not a leap year in 5767", func() {
				Expect(IsLeap(5767)).To(BeFalse())
			})
			It("is 12th year in Hillelian cycle and is not a leap year in 5769", func() {
				Expect(IsLeap(5769)).To(BeFalse())
			})
			It("is 13th year in Hillelian cycle and is not a leap year in 5770", func() {
				Expect(IsLeap(5770)).To(BeFalse())
			})
			It("is 15th year in Hillelian cycle and is not a leap year in 5772", func() {
				Expect(IsLeap(5772)).To(BeFalse())
			})
			It("is 16th year in Hillelian cycle and is not a leap year in 5773", func() {
				Expect(IsLeap(5773)).To(BeFalse())
			})
			It("is 18th year in Hillelian cycle and is not a leap year in 5775", func() {
				Expect(IsLeap(5775)).To(BeFalse())
			})
		})
	})
})
