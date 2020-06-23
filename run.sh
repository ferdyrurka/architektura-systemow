echo "############ MOV"
go run main.go mov -a BX -b DX
go run main.go mov -a AX -b BX

echo "############ XCHG"
go run main.go xchg -a BX -b AX
go run main.go xchg -a DX -b CX

echo "############ STI"
go run main.go sti

echo "############ CLI"
go run main.go cli