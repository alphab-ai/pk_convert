# pk_convert
convert secp256k1 private key to armored privkey


1) modify `main.go` and replace template string with your PRIVATE KEy  
2) Run the script with `go run main.go`


Example of output:

```
$ go run main.go

Armored Private Key:

-----BEGIN TENDERMINT PRIVATE KEY-----
type: secp256k1
kdf: bcrypt
salt: B2DE1066B0BD653430CC538AE46D02

y+9zUKdLac9M0WHzxfhdfhdfhkcS48Gsr5f4BPZJ+HB2X85U9ctYJw
JlJzvLLhF2K1DEfdfhdfhlbHdR9Ql8XOuM=
=FhMu
-----END TENDERMINT PRIVATE KEY-----
```