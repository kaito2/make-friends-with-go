package fuzzing

func target(a int, b string) {
	if 1024 < a && a < 4096 {
		panic("1024!!!")
	}
}
