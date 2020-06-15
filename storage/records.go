package storage

var stringRecords = map[string][]byte{
	"AX": []byte("ax_value"),
	"BX": []byte("bx_value"),
	"CX": []byte("cx_value"),
	"DX": []byte("dx_value"),
}

var boolRecords = map[string]byte{
	"IF": byte(0),
}

func FindAll() (map[string][]byte, map[string]byte) {
	return stringRecords, boolRecords
}

func Mov(a string, b string) {
	stringRecords[a] = stringRecords[b]
}

func Xchg(a string, b string) {
	ab := stringRecords[b]
	ba := stringRecords[a]

	stringRecords[a] = ab
	stringRecords[b] = ba
}

func Cli() {
	boolRecords["IF"] = 0
}

func Sti()  {
	boolRecords["IF"] = 1
}
