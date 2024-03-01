
## Run app
```shell
go run main.go
```

## Testing
```shell
go test  -coverprofile cover.out ./...
```

## API endpoints
#### GET /block
Returns last parsed block

**Response**
```json
{
  "current_block":19338329
}
```
#### GET /address/{address}
Returns all transactions to the requested address

**Response**
```json
[
  {
    "blockHash": "0xcf6c01ec452f73a1c3dcdfe68ba5d70da6ca0b01d2dd95afe1485bbdb11ea646",
    "blockNumber": "0x12714f4",
    "from": "0x95222290dd7278aa3ddd389cc1e1d165cc4bafe5",
    "gas": "0x6ac1",
    "gasPrice": "0xaffdc638b",
    "hash": "0xb37936d9e59eeed69592ed984d46413cb5a2330f51b1b2737c0620429b69029f",
    "input": "0x",
    "nonce": "0xb1ebf",
    "to": "0x8306300ffd616049fd7e4b0354a64da835c1a81c",
    "transactionIndex": "0x6e",
    "value": "0x4c2aa381b1f343",
    "v": "0x1",
    "r": "0xcb4e1f179b96cb55a220c6391b9b60d42e664bc254d1c620e70e1591251a37c1",
    "s": "0x6c8e3e778557114ad5881ceb3b5d6d48a6ab70594cf81c0b330e8ec128f61d77"
  }
]
```

#### POST /address
adds address to observer

**Request body**
```json
{
  "address": "0x95222290dd7278aa3ddd389cc1e1d165cc4bafe5"
}
```
**Response**
```json
{
  "address": "0x95222290dd7278aa3ddd389cc1e1d165cc4bafe5",
  "status": true
}
```