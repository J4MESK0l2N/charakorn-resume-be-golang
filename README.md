## API with Golang

### API เส้นที่ Test แล้ว
- api/v1/profile [GET]
- api/v1/profile [POST]
- api/v1/profile [DELETE]
- api/v1/organizations [GET]
- api/v1/organizations [POST]
- api/v1/organizations [GET LIST]
- api/v1/organizations [DELETE]
- api/v1/project [POST]

### จะมีบางเส้นที่ยังไม่ได้เทส เนื่องจากทำ API เส้นที่ใช้งานก่อน
- Project ถูก Deploy อยู่บน Lambda > API Gateway โดยทำการเก็บ Image Docker ไว้ที่ ECS เพราะ Golang 1.x เลิก Support บน aws deploy จึงต้องใช้ Docker ขึ้นมาในการ Build Image

## Environments
- Golang (gin)
- DynamoDB
- Docker
