# latihan GRPC

## structure table user

- Id uint64 Priamry key
- Name string
- Email string
- Age uint64
- Jenkel string

## services User

- rpc ShowAll (User) returns (DataUsers);
- rpc CreateUser (User) returns (User);
- rpc FindById (User) returns (User);
- rpc UpdateUser(User) returns (User);
- rpc DeleteUser (User) returns (User);
