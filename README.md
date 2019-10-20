# github.com/greyireland/kratos-demo

## 项目简介
1.


## 测试

```markdown
# 添加
curl -H 'Content-Type:application/json' -XPOST '127.1:8000/v1/message/add' -d '{"session_id":"1:2","uid":1,"peer_uid":2,"message":"hi"}'

# 获取
curl '127.1:8000/v1/message/list?uid=1&peer_uid=2'

```