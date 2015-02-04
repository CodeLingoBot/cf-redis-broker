package api_test

import (
	. "github.com/onsi/ginkgo"
	"github.com/onsi/ginkgo/reporters"
	. "github.com/onsi/gomega"

	"testing"
)

func TestApi(t *testing.T) {
	RegisterFailHandler(Fail)
	junitReporter := reporters.NewJUnitReporter("junit_api.xml")
	RunSpecsWithDefaultAndCustomReporters(t, "Api Suite", []Reporter{junitReporter})
}
