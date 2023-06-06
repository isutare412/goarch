package pkgerr_test

import (
	"errors"
	"fmt"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"github.com/isutare412/goarch/http-base/pkg/pkgerr"
)

var _ = Describe("Simple", func() {
	It("unwraps inner origin error", func() {
		var (
			givenOriginErrs = [...]error{
				errors.New("origin-error-one"),
				errors.New("origin-error-two"),
			}
			givenPkgErr = pkgerr.Simple{
				Origin: givenOriginErrs[0],
				Simple: fmt.Errorf("somethings got wrong"),
			}
			givenJoinedErr = fmt.Errorf("complex error occurred: %w: %w", givenPkgErr, givenOriginErrs[1])
		)

		Expect(errors.Is(givenJoinedErr, givenOriginErrs[0])).To(BeTrue())
		Expect(errors.Is(givenJoinedErr, givenOriginErrs[1])).To(BeTrue())
		Expect(errors.Is(givenJoinedErr, givenPkgErr)).To(BeTrue())
	})

	It("finds simple error from error chain", func() {
		var (
			givenSimpleErrorMsg = "simple-error"
			givenPkgErr         = pkgerr.Simple{Simple: fmt.Errorf(givenSimpleErrorMsg)}
			givenWrappedErr     = fmt.Errorf("wrapped: %w", givenPkgErr)
		)

		var serr pkgerr.Simple
		Expect(errors.As(givenWrappedErr, &serr)).To(BeTrue())
		Expect(serr.SimpleError()).To(Equal(givenSimpleErrorMsg))
	})
})
