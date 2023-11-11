# Project: RSSAGG

## End-point: Health Route

This HTTP GET request is used to check the readiness of the server. It sends a request to the endpoint at http://localhost:8080/v1/ready.

### Request

The request does not require any parameters.

### Response

The server responds with a status code of 200, indicating a successful request. The response body is empty.


### Method: GET
>```
>http://localhost:8080/v1/ready
>```

⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃

## End-point: Create User

This API endpoint allows you to create a new user by sending an HTTP POST request to the specified URL. The request should include the necessary parameters in the request body.

### Request

- Method: POST
- URL: `http://localhost:8080/v1/users`

### Response

The response to this request will have a status code of 201 if the user creation is successful. The response body will contain the following properties:

- `id`: The unique identifier of the created user.
- `created_at`: The timestamp when the user was created.
- `updated_at`: The timestamp when the user was last updated.
- `name`: The name of the user.
- `api_key`: The API key associated with the user.

Please note that the actual values for these properties will be provided in the response, but for security reasons, they are not included in this documentation.

Please ensure that you include all the required parameters in the request body according to the API documentation.

### Method: POST
>```
>http://localhost:8080/v1/users
>```
### Body (**raw**)

```json
{
    "name":"sssddd"
}
```


⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃

## End-point: Get User
This endpoint is used to retrieve a list of users. It sends an HTTP GET request to the URL `http://localhost:8080/v1/users`. The request does not require any additional parameters.

The last execution of this request returned a response with a status code of 200, indicating a successful request. The response body included the following properties:

- `id`: The unique identifier of the user.
- `created_at`: The date and time when the user was created.
- `updated_at`: The date and time when the user was last updated.
- `name`: The name of the user.
- `api_key`: The API key associated with the user.


Please note that the actual values for these properties were not provided in the response.

You can use this endpoint to retrieve a list of users and access their details. Make sure to handle the response appropriately based on the status code and the properties provided in the response body.
### Method: GET
>```
>http://localhost:8080/v1/users
>```
### Headers

|Content-Type|Value|
|---|---|
|Authorization|ApiKey c15dad31abe6dc4a342df56c49ea96e252128d6b8127cfe6aae757615c8a850c|



⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃

## End-point: Post Feeds
This API endpoint allows you to create a new feed. It is an HTTP POST request that should be sent to [http://localhost:8080/v1/feeds](http://localhost:8080/v1/feeds).

The request should include the following parameters in the request body:

- id (string): The unique identifier for the feed.
- created_at (string): The timestamp when the feed was created.
- updated_at (string): The timestamp when the feed was last updated.
- name (string): The name of the feed.
- url (string): The URL of the feed.
- user_id (string): The unique identifier of the user who owns the feed.


The response to this request will have a status code of 201, indicating that the feed was successfully created. The response body will contain the following properties:

- id (string): The unique identifier for the created feed.
- created_at (string): The timestamp when the feed was created.
- updated_at (string): The timestamp when the feed was last updated.
- name (string): The name of the feed.
- url (string): The URL of the feed.
- user_id (string): The unique identifier of the user who owns the feed.


Please note that the values for the properties in the response may be different from the empty values shown in the example response above.

To use this API, make sure to include all the required parameters in the request body and send an HTTP POST request to the specified URL. You will receive a response with the details of the created feed.
### Method: POST
>```
>http://localhost:8080/v1/feeds
>```
### Headers

|Content-Type|Value|
|---|---|
|Authorization|ApiKey c15dad31abe6dc4a342df56c49ea96e252128d6b8127cfe6aae757615c8a850c|


### Body (**raw**)

```json
{
    "name":"wagswlaneoqg",
    "url":"https://wawgslane.dev/index.xml"
}
```


⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃

## End-point: Get Feeds

This API endpoint is used to retrieve a list of feeds. It sends an HTTP GET request to the URL http://localhost:8080/v1/feeds.

### Request

The request does not require any parameters.

### Response

The response will be a JSON array containing multiple feed objects. Each feed object will have the following properties:

- `id`: The unique identifier of the feed.
- `created_at`: The timestamp when the feed was created.
- `updated_at`: The timestamp when the feed was last updated.
- `name`: The name of the feed.
- `url`: The URL of the feed.
- `user_id`: The ID of the user who owns the feed.

Example response:

```json
[
    {
        "id": "",
        "created_at": "",
        "updated_at": "",
        "name": "",
        "url": "",
        "user_id": ""
    }
]
```

Please note that the values for the properties `id`, `created_at`, `updated_at`, `name`, `url`, and `user_id` will be specific to each feed and may not be provided in this example response.


### Method: GET
>```
>http://localhost:8080/v1/feeds
>```

⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃

## End-point: Post Feed Follows

This API endpoint allows you to create a new feed follow. A feed follow represents a user following a specific feed.

To create a new feed follow, send a POST request to `http://localhost:8080/v1/feed_follows`.

### Request Parameters
The request should include the following parameters in the request body:
- `feed_id` (string): The ID of the feed that the user wants to follow.
- `user_id` (string): The ID of the user who wants to follow the feed.

### Response
If the request is successful, the API will return a response with a status code of 201 (Created) and the following properties:
- `id` (string): The ID of the newly created feed follow.
- `created_at` (string): The timestamp when the feed follow was created.
- `updated_at` (string): The timestamp when the feed follow was last updated.
- `feed_id` (string): The ID of the feed that the user is following.
- `user_id` (string): The ID of the user who is following the feed.

Please note that the values for `id`, `created_at`, `updated_at`, `feed_id`, and `user_id` will be populated by the API and will not be empty in the actual response.

Example Response:
```
{
    "id": "123456789",
    "created_at": "2021-01-01T12:00:00Z",
    "updated_at": "2021-01-01T12:00:00Z",
    "feed_id": "987654321",
    "user_id": "abcdef123"
}
```

Remember to include the required request parameters in the request body to create a new feed follow.

### Method: POST
>```
>http://localhost:8080/v1/feed_follows
>```
### Headers

|Content-Type|Value|
|---|---|
|Authorization|ApiKey c15dad31abe6dc4a342df56c49ea96e252128d6b8127cfe6aae757615c8a850c|


### Body (**raw**)

```json
{
    "feed_id":"efe7dc15-b177-41f2-b926-069155bb4513"
}
```


⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃

## End-point: Get Feed Follows
This HTTP GET request is used to retrieve a list of feed follows. The request is made to the endpoint `http://localhost:8080/v1/feed_follows`.

The response from the last execution of this request had a status code of 200, indicating a successful response. The response body was an array of feed follow objects, with each object containing the following properties:

- `id`: The unique identifier of the feed follow.
- `created_at`: The timestamp indicating when the feed follow was created.
- `updated_at`: The timestamp indicating when the feed follow was last updated.
- `feed_id`: The ID of the feed being followed.
- `user_id`: The ID of the user who is following the feed.


Please note that the actual values for these properties were not provided in the response, so they have been omitted in this description.

To use this endpoint, make an HTTP GET request to `http://localhost:8080/v1/feed_follows`. No request parameters are required.

Example Request:

```
GET http://localhost:8080/v1/feed_follows

 ```

Example Response:

```
HTTP/1.1 200 OK
Content-Type: application/json
[
    {
        "id": "123",
        "created_at": "2022-01-01T10:00:00Z",
        "updated_at": "2022-01-01T12:00:00Z",
        "feed_id": "456",
        "user_id": "789"
    },
    {
        "id": "234",
        "created_at": "2022-01-02T09:00:00Z",
        "updated_at": "2022-01-02T11:00:00Z",
        "feed_id": "567",
        "user_id": "890"
    }
]

 ```

Please note that the actual response may contain more feed follow objects depending on the number of follows in the system.
### Method: GET
>```
>http://localhost:8080/v1/feed_follows
>```
### Headers

|Content-Type|Value|
|---|---|
|Authorization|ApiKey c15dad31abe6dc4a342df56c49ea96e252128d6b8127cfe6aae757615c8a850c|



⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃

## End-point: Delete Feed Follow

This API endpoint is used to delete a feed follow record.

## Request

- Method: DELETE
- URL: `http://localhost:8080/v1/feed_follows/ef27a024-37c0-4c00-acd0-62f3fa4b108d`

## Response

The API returns a response with a status code of 200. The response body is a JSON object with a single property:

- `status`: The status of the deletion operation.

Please note that the response body may be empty or contain additional properties not mentioned here.


### Method: DELETE
>```
>http://localhost:8080/v1/feed_follows/ed615343-9bfc-4463-abca-b5c2b984069a
>```
### Headers

|Content-Type|Value|
|---|---|
|Authorization|ApiKey c15dad31abe6dc4a342df56c49ea96e252128d6b8127cfe6aae757615c8a850c|



⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃ ⁃
