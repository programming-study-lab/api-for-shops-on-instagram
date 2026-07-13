# api-for-shops-on-instagram
## ข้อมูล Instagram ของตนเอง
*** Request Body: ***
```
curl -X GET http://localhost:5000/api/v1/instagram/
```
*** Response: ***
```
{
    "status": true,
    "message": "success",
    "data": [
        {
            "id":"<instagram_id>"
        }
    ]
}
```
## โพสต์ content
***Request Body:***
```
curl -X POST http://localhost:5000/api/v1/instagram/media \
  -H "Content-Type: application/json" \
  -d '{
        "image_url":"<image_url>",
        "caption": "<ข้อความ>",
    }'
```
***Response:***
```
{
    "data": {
        "id": "<media_id>"
    },
    "message": "success",
    "status": true
}
```
## การสนทนา
***Request Body:***
```
curl -X GET http://localhost:5000/api/v1/instagram/conversations
```
***Response:***
```
{
    "status": true,
    "message": "succsss",
    "data": {
        "data": [
            {
                "id": "<conversation_id>",
                "updated_time": "<time>"
            }
        ]
    }
}
```
## รายการ การสนทนา
***Request Body:***
```
curl -X GET http://localhost:5000/api/v1/instagram/<conversation_id>/messages
```
***Response:***
```
{
    "status": true,
    "message": "succsss",
    "data": {
        "data": [
            {
                "id": "<message_id>"
            }
        ],
        "paging": {
            "cursors": {
                "after": "<api_pagination>"
            },
            "next": "<api_pagination>"
        }
    }
}
```
## รายการ ข้อความแชท
***Request Body:***
```
curl -X GET http://localhost:5000/api/v1/instagram/messages/<message_id>
```
***Response:***
```
{
    "status": true,
    "message": "succsss",
    "data": {
        "created_time": "<time>",
        "from": {
            "id": "<instagram_id ผู้ส่ง>",
            "username": "<instagram_username ผู้ส่ง>"
        },
        "id": <message_id>",
        "message": "<ข้อความแชท>",
        "to": {
            "data": [
                {
                    "id": "<instagram_id ผู้รับ>",
                    "username": "<instagram_username ผู้รับ>"
                }
            ]
        }
    }
}
```
## ส่งข้อความ
***Request Body:***
```
curl -X POST http://localhost:5000/api/v1/instagram/messages \
  -H "Content-Type: application/json" \
  -d '{
    "recipient": { 
        "id": "<instagram_id ผู้รับ>" 
    },
    "message": { 
        "text": "<ข้อความ>" 
    }
}'
```
***Response:***
```
{
    "status": true,
    "message": "succsss",
    "data": {
        "message_id": "<message_id>",
        "recipient_id": "<instagram_id ผู้รับ>"
    }
}
```
## แสดงรายการโพสต์
***Request Body:***
```
curl -X GET http://localhost:5000/api/v1/instagram/media
```
***Response:***
```
{
    "data": {
        "id": "<media_id>",
    },
    "message": "message",
    "status": true
}
```
## แสดงรายละเอียดโพสต์
***Request Body:***
```
curl -X GET http://localhost:5000/api/v1/instagram/media/<media_id>
```
***Response:***
```
{
    "status": true,
    "message": "succsss",
    "data": {
        "caption": "<caption>",
        "id": "<media_id>",
        "media_type": "<media_type>",
        "media_url": "<media_url>",
        "permalink": "<link ไปยังโพสต์ใน instagram>",
        "timestamp": "<time>"
    }
}
```