package config

import "testing"

func TestIsIPv4Net(t *testing.T) {
	type testCase struct {
		value string
		wait  bool
	}

	var tests = []testCase{
		{"127.0.0.1", true},
		{"0.0.0.0", true},
		{"255.255.255.255", true},
		{"0.0.0.0/0", false},
		{"0.0.0.0/32", false},
		{"256.0.0.0", false},
		{"0.256.0.0", false},
		{"0.0.256.0", false},
		{"0.0.0.256", false},
		{"abcd", false},
		{"a.b.c.d", false},
		{"", false},
		{"-1.0.0.0", false},
	}

	for _, test := range tests {
		if IsIPv4Net(test.value) != test.wait {
			t.Errorf("For %s expected %t", test.value, test.wait)
		}
	}
}

func TestIsURL(t *testing.T) {
	type testCase struct {
		value string
		wait  bool
	}

	var tests = []testCase{
		{"http://127.0.0.1", true},
		{"https://127.0.0.1", true},
		{"https://127.0.0.1/", true},
		{"https://127.0.0.1/test.html", true},
		{"https://127.0.0.1/sub/test.html", true},
		{"https://127.0.0.1/sub/test.html?param1=value", true},
		{"https://127.0.0.1/sub/test.html?param1=value#anker", true},
		{"http://google.com", true},
		{"https://google.com/ru/?go=test#ok", true},
		{"test://google.com/ru/?go=test#ok", true},
		{"file:///home/?go=test#ok", true},
		{"///home/?go=test#ok", true},
		{"http://:8080", true},
		{"#http://:8080", false},
		{"?http://:8080", false},
		{":///home/?go=test#ok", false},
		{"home/?go=test#ok", false},
		{"./home/?go=test#ok", false},
		{"./home.html", false},
	}

	for _, test := range tests {
		if IsURL(test.value) != test.wait {
			t.Errorf("For %s expected %t", test.value, test.wait)
		}
	}
}

// в Go нет тернарного оператора..
func ternary(t string, f string, condition bool) string {
	if condition {
		return t
	}
	return f
}

func TestTestConfig(t *testing.T) {
	type testCase struct {
		value  Config
		isNull bool
	}

	var tests = []testCase{
		{value: Config{
			Debug: false, MyUrl: "https://google.com", Database: confDatabase{Host: "127.0.0.1", Port: 54321, User: "test", Pass: "pwd123", Ssl: true}, Server: confServer{Bind: []string{"10.0.0.1", "127.0.0.1"}, Port: 8888, LogLevel: 2},
		}, isNull: true},
		{value: Config{
			Debug: false, MyUrl: "https://google.com", Database: confDatabase{Host: "127.0.0.1", Port: 54321, User: "test", Pass: "pwd123", Ssl: true}, Server: confServer{Bind: []string{"10.0.0.1", "127.0.0.1"}, Port: 8888, LogLevel: 999},
		}, isNull: true},
		{value: Config{
			Debug: false, MyUrl: "https://google.com", Database: confDatabase{Host: "127.0.0.1", Port: 54321, User: "test", Pass: "pwd123", Ssl: true}, Server: confServer{Bind: []string{"10.0.0.1", "127.0.0.1"}, Port: 8888, LogLevel: 2},
		}, isNull: true},
		{value: Config{
			Debug: false, MyUrl: "FAIL!!", Database: confDatabase{Host: "127.0.0.1", Port: 54321, User: "test", Pass: "pwd123", Ssl: true}, Server: confServer{Bind: []string{"10.0.0.1", "127.0.0.1"}, Port: 8888, LogLevel: 2},
		}, isNull: false},
		{value: Config{
			Debug: false, MyUrl: "https://google.com", Database: confDatabase{Host: "FAIL!!", Port: 54321, User: "test", Pass: "pwd123", Ssl: true}, Server: confServer{Bind: []string{"10.0.0.1", "127.0.0.1"}, Port: 8888, LogLevel: 2},
		}, isNull: false},
		{value: Config{
			Debug: false, MyUrl: "https://google.com", Database: confDatabase{Host: "127.0.0.1", Port: 54321, User: "test", Pass: "pwd123", Ssl: true}, Server: confServer{Bind: []string{"FAIL!!", "10.0.0.1", "127.0.0.1"}, Port: 8888, LogLevel: 2},
		}, isNull: false},
	}

	for _, test := range tests {
		// в Go нет XOR.. (X || Y) && !(X && Y)
		if ok := testConfig(&test.value); (ok == nil || test.isNull) && !(ok == nil && test.isNull) {
			t.Errorf("For %+v %s nill\n", test.value, ternary("expected", "unexpected", test.isNull))
		}
	}
}
