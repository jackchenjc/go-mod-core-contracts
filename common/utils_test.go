//
// Copyright (C) 2023 IOTech Ltd
//
// SPDX-License-Identifier: Apache-2.0

package common

import (
	"net/url"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestURLEncode(t *testing.T) {
	tests := []struct {
		name   string
		input  string
		output string
	}{
		{"valid", "^[this]+{is}?test:string*#", "%5E%5Bthis%5D%2B%7Bis%7D%3Ftest:string%2A%23"},
		{"valid - special character", "this-is_test.string~哈囉世界< >/!#%^*()+,`@$&", "this%2Dis%5Ftest%2Estring%7E%E5%93%88%E5%9B%89%E4%B8%96%E7%95%8C%3C%20%3E%2F%21%23%25%5E%2A%28%29%2B%2C%60@$&"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res := URLEncode(tt.input)
			assert.Equal(t, tt.output, res)

			unescaped, err := url.PathUnescape(tt.output)
			require.NoError(t, err)
			assert.Equal(t, tt.input, unescaped)
		})
	}
}
