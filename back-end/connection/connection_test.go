package main

import "testing"

func Test_sign_in1(t *testing.T) {
	var (
		expected = "Success"
	)
	actual := sign_in("Roman", "11")
	if actual != expected {
		t.Errorf("actual=%s; expected %s", actual, expected)
	}
}

func Test_sign_in2(t *testing.T) {
	var (
		expected = "Wrong password"
	)
	actual := sign_in("Roman", "1")
	if actual != expected {
		t.Errorf("actual=%s; expected %s", actual, expected)
	}
}

func Test_sign_in3(t *testing.T) {
	var (
		expected = "Wrong password"
	)
	actual := sign_in("roman", "1")
	if actual != expected {
		t.Errorf("actual=%s; expected %s", actual, expected)
	}
}

func Test_sign_up1(t *testing.T) {
	var (
		expected = "username already exists"
	)
	actual := sign_up("Roman", "fadsfa")
	if actual != expected {
		t.Errorf("actual=%s; expected %s", actual, expected)
	}
}

func Test_sign_up2(t *testing.T) {
	var (
		expected = "username already exists"
	)
	actual := sign_up("roman", "fadsfa")
	if actual != expected {
		t.Errorf("actual=%s; expected %s", actual, expected)
	}
}

func Test_order_food1(t *testing.T) {
	var (
		expected = "Success"
	)
	actual := order_food("Roman")
	if actual != expected {
		t.Errorf("actual=%s; expected %s", actual, expected)
	}
}

func Test_get_info(t *testing.T) {
	var (
		expected = "Roman"
	)
	sign_in("Roman", "11")
	actual := user_info()
	if actual != expected {
		t.Errorf("actual=%s; expected %s", actual, expected)
	}
}
