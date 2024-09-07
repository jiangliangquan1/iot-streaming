# api-list

[toc]

___

## 1 用户相关api

### 1.1 用户注册

- Method: POST
- Url: /iot-streaming/api/v1/users/sign-up
- Content-Type: application/json
- 请求参数

    | 参数名  | 类型      | 必填 | 示例 | 说明    |
    |------|---------|----|---|-------|
    | name |  string | 是  |  xiaoming |  用户名，唯一 |
    |nick  | string  | 否 | liuxiaoming | 昵称|
    | email| string  | 是  | xiaoming@gmail.com| 邮箱，唯一|
    | password| string| 是 | 123456| 密码|
- 响应参数

  | 参数名          | 类型     | 必填 | 示例                 | 说明         |
  |--------------|--------|----|--------------------|------------|
  | code         | int    | 是  | 200                | 响应码，200为成功 |
  | msg          | string | 否  | success            | 响应描述       |
  | result       | object | 否  |                    | 结果实体       | 
  | result.id    | long   | 是  | 1                  | 用户id       |
  | result.name  | string | 是  | xiaoming           | 用户名，唯一     |
  | result.nick  | string | 否  | liuxiaoming        | 昵称         |
  | result.email | string | 是  | xiaoming@gmail.com | 邮箱，唯一      |

- 请求示例
  - url:127.0.0.1:9966/iot-streaming/api/v1/users/sign-up
  - 请求参数
    ```json
    {
      "name":"xiaoming",
      "nick":"liuxiaoming",
      "email":"xiaoming@gmail.com",
      "password":"HwNFBXIVUSJLAMDljbusDg=="
    }
    ```
  - 响应参数
    ```json
    {
      "code": 200,
      "msg": "success",
      "result": {
         "id": 1,
         "name":"xiaoming",
         "nick":"liuxiaoming",
         "email":"xiaoming@gmail.com"
      }
    }
    ```







