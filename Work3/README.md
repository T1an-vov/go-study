​        自定义粘包协议协议解决粘包问题

​        客户端和服务端收发数据共计20次

​        客户端发送GetUserReq结构体，UserID为x(当前收发数据序号)。先将结构体进行json序列化后再调用proto包中的encode函数编码，发送至客户端

​         服务端收到数据后，先用decode函数解码，再json反序列化至GetUserResp结构体中，然后打印：

"收到第x次来自客户端的数据:"

"userID:UserID"

同时将GetUserResp的UserName设为"第x个用户的名字"，序列化并编码后发送给客户端

​         客户端收到数据后，再以同样的方式获取数据，并打印:

"收到第 x次来自服务端的数据:"  

"userID:UserID"

"username:UserName"

