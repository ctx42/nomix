//go:build ignore

package main

import (
	"github.com/ctx42/testing/pkg/mocker"
)

func main() {
	opts := []mocker.Option{
		mocker.WithTgtOnHelpers,
	}
	mocks := []func(...mocker.Option) error{
		GenTagMock,
	}
	for _, mock := range mocks {
		if err := mock(opts...); err != nil {
			panic(err)
		}
	}
}

func GenTagMock(opts ...mocker.Option) error {
	opts = append(opts, mocker.WithTgtFilename("tag_mock_test.go"))
	return mocker.Generate("Tag", opts...)
}
