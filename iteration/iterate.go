package iteration

func Iterate(val string, count int) string {
	var buildString string
	for i := 0; i < count; i++ {
		buildString += val
	}
	return buildString
}