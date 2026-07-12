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
            "id":"instagram_id"
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
        "uploadImageType": "file|url <(อัพโหลดโดยใช้ไฟล์ หรือ ใช้ url)>"
    }'
```
***Response:***
