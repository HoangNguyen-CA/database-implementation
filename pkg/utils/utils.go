package utils

func Assert(c bool) {
	if !c {
		panic("assertion failed")
	}
}
