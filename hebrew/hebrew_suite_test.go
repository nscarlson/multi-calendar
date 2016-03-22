package hebrewtime_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestHebrew(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Hebrew Suite")
}
