/*
 * MIT License
 *
 * Copyright (c)  2018 Kasun Vithanage
 *
 * Permission is hereby granted, free of charge, to any person obtaining a copy
 * of this software and associated documentation files (the "Software"), to deal
 * in the Software without restriction, including without limitation the rights
 * to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
 * copies of the Software, and to permit persons to whom the Software is
 * furnished to do so, subject to the following conditions:
 *
 * The above copyright notice and this permission notice shall be included in all
 * copies or substantial portions of the Software.
 *
 * THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
 * IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
 * FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
 * AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
 * LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
 * OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
 * SOFTWARE.
 */

package util

import (
	"testing"

	testifyAssert "github.com/stretchr/testify/assert"
)

func TestSplitSpacesWithQuotes(t *testing.T) {
	assert := testifyAssert.New(t)

	// test for balanced quotes
	s := `  lorem ipsum       "foo bar "   `
	res, err := SplitSpacesWithQuotes(s)
	assert.Nil(err)
	excpted := []string{"lorem", "ipsum", "foo bar "}
	assert.Equal(excpted, res)

	// test for unquoted balanced quotes
	s = `  lorem ipsum   "this is m\"e"    "foo bar "   `
	res, err = SplitSpacesWithQuotes(s)
	assert.Nil(err)
	assert.Equal([]string{"lorem", "ipsum", `this is m"e`, "foo bar "}, res)

	// test for characters quotes
	s = `  lorem ipsum       "foo bar &'' "   `
	res, err = SplitSpacesWithQuotes(s)
	assert.Nil(err)
	excpted = []string{"lorem", "ipsum", "foo bar &'' "}
	assert.Equal(excpted, res)

	// test for unbalanced quotes
	s = ` foo "bar""`
	res, err = SplitSpacesWithQuotes(s)
	assert.Equal(ErrUnbalancedQuotes, err)
	assert.Len(res, 0)
}

func BenchmarkSplitSpacesWithQuotes(b *testing.B) {
	testString := ` foo     bar "foo bar bar"    foo     bar "foo bar bar" foo     bar "foo bar bar" foo     bar "foo bar bar" \"`
	for i := 0; i < b.N; i++ {
		SplitSpacesWithQuotes(testString)
	}
}
