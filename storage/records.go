package storage

var stringRecords = map[string]uint16{
	"AX": uint16(100),
	"BX": uint16(500),
	"CX": uint16(2000),
	"DX": uint16(30000),
}

var boolRecords = map[string]byte{
	"IF": byte(0),
}

func FindAll() (map[string]uint16, map[string]byte) {
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
