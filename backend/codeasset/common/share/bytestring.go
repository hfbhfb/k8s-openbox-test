package share

func ByteAndErrToString(a []byte, err error) string {
	if len(a) == 0 || err != nil {
		return ""
	}
	return string(a)
}
