// @File: stringRegexp_test
// @Date: 2020/12/26 20:02
// @Author: 安红豆
// @Description: 正则检测工具方法的测试方法
package util

import (
	"testing"
)

type testDate struct {
	regexpString   string
	toBeTestString string
	matched        bool
}

//测试用例
var (
	//IP
	testDateIP = []testDate{
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
	}
	//域名
	testDateDomainName = []testDate{
		{RegexpExpressionDomainName, "www.baidu.com", true},
		{RegexpExpressionDomainName, "baidu.com", true},
		{RegexpExpressionDomainName, "bai-du.com", true},
		{RegexpExpressionDomainName, "baidu.123", true},
		{RegexpExpressionDomainName, "baidu", false},
	}
)

//普通正则匹配方法测试
func TestStringRegexp(t *testing.T) {
	//测试IP正则
	for _, tt := range testDateIP {
		if matched := StringRegexp(tt.regexpString, tt.toBeTestString); matched != tt.matched {
			t.Errorf("TestStringRegexp(RegexpExpressionIP,%s),got %t, but %t", tt.toBeTestString, matched, tt.matched)
		}
	}

	//测试域名正则
	for _, tt := range testDateDomainName {
		if matched := StringRegexp(tt.regexpString, tt.toBeTestString); matched != tt.matched {
			t.Errorf("TestStringRegexp(RegexpExpressionDomainName,%s),got %t, but %t", tt.toBeTestString, matched, tt.matched)
		}
	}
}

//闭包正则匹配方法测试
func TestRegexpBase(t *testing.T) {
	//测试IP正则
	for _, tt := range testDateIP {
		regexpIP := RegexpBase(RegexpExpressionIP)
		if matched := regexpIP(tt.toBeTestString); matched != tt.matched {
			t.Errorf("TestStringRegexp(RegexpExpressionIP,%s),got %t, but %t", tt.toBeTestString, matched, tt.matched)
		}
	}

	//测试域名正则
	for _, tt := range testDateDomainName {
		regexpDomainName := RegexpBase(RegexpExpressionDomainName)
		if matched := regexpDomainName(tt.toBeTestString); matched != tt.matched {
			t.Errorf("TestStringRegexp(RegexpExpressionDomainName,%s),got %t, but %t", tt.toBeTestString, matched, tt.matched)
		}
	}
}

//普通正则匹配方法性能测试
func BenchmarkStringRegexp(b *testing.B) {
	for i := 0; i < b.N; i++ {
		StringRegexp(RegexpExpressionIP, "127.0.0.1")
	}
}

//闭包正则匹配方法性能测试
func BenchmarkRegexpBase(b *testing.B) {
	regexpIP := RegexpBase(RegexpExpressionIP)
	for i := 0; i < b.N; i++ {
		regexpIP("127.0.0.1")
	}
}
