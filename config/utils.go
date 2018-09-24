package configuration

// CheckErr function to check errors ...CheckErr
func CheckErr(err error) {
	if err != nil {
		panic(err)
	}
}
