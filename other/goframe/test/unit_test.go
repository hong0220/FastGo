package test

import (
	"github.com/gogf/gf/v2/test/gtest"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestUnit(t *testing.T) {
	//ctx := gctx.New()

	gtest.C(t, func(t *gtest.T) {
		expected := []int{0, 1}
		actual := []int{0, 1}
		assert.Equal(t, expected, actual)

		t.AssertEQ(len(expected), 2)
		t.AssertGT(len(expected), 0)
		t.AssertGT(len(expected), 1)

		//t.AssertNil(err)
	})
}
