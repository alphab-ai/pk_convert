# pk_convert
convert secp256k1 private key to armored privkey


1) build `go build -o pk_convert ./main.go`  
```
./pk_convert -h                                                                   
Usage of ./pk_convert:
  -password string
    	 (default "12345678")
  -pk string
    	hex private key 64 chars (default " ")
```

2) Run pk_convert example:  
`./pk_convert --pk=a0da0296d11eb47bcbee2254369b80d7400862f5ec44eb1cd03c820b92856a4c`  
`./pk_convert --pk=a0da0296d11eb47bcbee2254369b80d7400862f5ec44eb1cd03c820b92856a4c --password=12345678`

Example of output:

```
./pk_convert --pk=a0da0296d11eb47bcbee2254369b80d7400862f5ec44eb1cd03c820b92856a4c --password=12345678
-----BEGIN TENDERMINT PRIVATE KEY-----
salt: 97FB5D3F89AFECE096F7C770DA35F6F4
type: secp256k1
kdf: bcrypt

Z1uXTf0sCFP15ZPjtixn9xSKU0E5XsF6j5u6vJXJq1ukQaIyWUwJ4HNo+Fki/4Fy
ROKiVxbIOu+h4UuOEVAPp7oVX5LfxYiQMZF3LWk=
=7HWh
-----END TENDERMINT PRIVATE KEY-----
➜  pk_convert git:(main) ✗ ./pk_convert --pk=a0da0296d11eb47bcbee2254369b80d7400862f5ec44eb1cd03c820b92856a4c                    
-----BEGIN TENDERMINT PRIVATE KEY-----
kdf: bcrypt
salt: 8FC5C45401E12A8D9EA6B8D8BC0F4EB5
type: secp256k1

RiB1QAVdyNE++OxaS4wUUVTfxCvkP97V9/1fXTc6O55Kd07PtCIB4tRGwvdN6nZY
oHs98X7HACAeucHMFi+D3Kni4fiiqy24C+Pj9hE=
=+uAM
-----END TENDERMINT PRIVATE KEY-----

```