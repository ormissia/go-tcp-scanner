// @File: stringRegexp_test
// @Date: 2020/12/26 20:02
// @Author: 安红豆
// @Description: 正则检测工具方法的测试方法
package util

import (
	"testing"
)

func TestStringRegexp(t *testing.T) {
	//测试用例
	tests := []struct {
		regexpString   string
		toBeTestString string
		matched        bool
	}{
		//IP
		{RegexpExpressionIP, "127.0.0.1", true},
		{RegexpExpressionIP, "1.0.0.1", true},
		{RegexpExpressionIP, "255.255.255.255", true},
		{RegexpExpressionIP, "0.0.0.0", false},
		{RegexpExpressionIP, "0.0.0.1", false},
		{RegexpExpressionIP, "256.255.255.255", false},
		{RegexpExpressionIP, "255.255.255.256", false},
		{RegexpExpressionIP, "255.255.255", false},
		{RegexpExpressionIP, "a.b.c.d", false},
		{RegexpExpressionIP, "^&*^*.&(*,(*&.^)*.@#$.0", false},
		//域名
		{RegexpExpressionDomainName, "www.baidu.com", true},
		{RegexpExpressionDomainName, "baidu.com", true},
		{RegexpExpressionDomainName, "bai-du.com", true},
		{RegexpExpressionDomainName, "baidu.123", true},
		{RegexpExpressionDomainName, "baidu", false},
	}

	for _, tt := range tests {
		if matched := StringRegexp(tt.regexpString, tt.toBeTestString); matched != tt.matched {
			t.Errorf("TestStringRegexp(,%s),got %t, but %t", tt.toBeTestString, matched, tt.matched)
		}
	}
}
