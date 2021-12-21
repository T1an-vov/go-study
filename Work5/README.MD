使用了gin+gorm框架，在本地8080端口测试

//
各包说明：
    dao：连接数据库并用自动迁移创建了users、courses、records三个表
    service:对应三个的服务，实现对users、courses、records的增删查改。
    controller:对应三个服务的控制器，并对service返回的结果进行处理，返回json
    module:定义了user、course、record三个模板
    router:定义了分组路由
//

各个接口说明：
  user:
    localhost/user/
    GET:传入参数为username
        若用户不存在，则查询失败
        查询成功则输出用户的当前学分和最大学分
    POST:传入参数为username、password、maxcredit
         其中username为唯一值，若已存在则创建失败；password不能为空；maxcredit必须为整数
         若参数无误则创建对应的用户
    PUT:传入参数为username、newpassword
        若用户不存在则修改密码失败
        成功则修改用户密码
    DELETE:传入参数为username，删除对应的用户
           若用户不存在，则删除失败
           成功则删除用户
           //
 course:
    localhost/course/
    GET:传入参数为coursename
        若课程不存在则查询失败
        查询成功输出课程信息
    POST:传入参数为coursename、credit、maxnumber
        若课程已存在则创建失败
        成功则创建课程
    PUT:传入参数为coursename、credit/maxnumber
        若课程不存在则修改失败
        成功则修改课程信息
    DELETE:传入参数为coursename
           若课程不存在 ，则删除失败
           存在则删除课程
           //
 record:
    localhost/record/
    GET:传入参数为username
        若该用户不存在选课记录则查询失败
        查询成功则输出选课信息
    POST:传入参数为username、coursename，表示用户进行选课；
         若用户或课程不存在，或选课人数已满或用户学分不够则创建记录失败
         成功则创建选课记录，同时修改用户的学分和课程的选课人数
    DELETE:传入参数为username和coursename，表示删除用户的选课记录。
           若该记录不存在，则删除失败
           成功则删除记录，同时修改用户学分和课程选课人数
           //

//

