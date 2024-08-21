# API Spec


## Get Value

endpoint: `/values/:key`\
method: GET

payload: `None`\
response: `{"value": ":value"}`

response codes:
- 200 (with response above)
- 404 (no data found -> with error response)
- 500 (server error -> with error response)

## Set Value

endpoint: `/values/:key`\
method: PUT

payload `{"value": ":value"}`\
response: `{"value": "added to store"}`

response codes:
- 200 (with above response)
- 500 (server error -> with error response)


## Error Response

```
{"error": ":error_message"}
```
