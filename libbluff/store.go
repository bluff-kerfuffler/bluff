package libbluff

var m  map[string]verify

func storeVerifAdd(v verify) {
	if m == nil {
		m = make(map[string]verify)
	}
	m[v.token] = v
}

func getVerifAdd(token string) verify {
	return m[token]
}