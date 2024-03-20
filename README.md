# Movie Service

This service provides endpoints to manage movies.

## Endpoints

- **Add Movie**: `POST /movie`
## Add Movie

### Request Body

```json
{
    "title": "Avatar",
    "description": "Avatar Of Ang",
    "rating": 7.6,
    "image": ""
}
```

### Response Example

```json
{
    "id": 4,
    "title": "Avatar",
    "description": "Avatar Of Ang",
    "rating": 7.6,
    "image": "",
    "created_at": "2024-03-21 02:37:22",
    "updated_at": "2024-03-21 02:37:22"
}
```

- **List Movies**: `GET /movie`
## List Movies

### Response Example

```json
[
    {
        "id": 1,
        "title": "Pengabdi Setan Comunion",
        "description": "dalah sebuah film horor Indonesia tahun 2022 yang disutradarai dan ditulis oleh Joko Anwar sebagai sekuel dari film tahun 2017, Pengabdi Setan.",
        "rating": 7,
        "image": "",
        "created_at": "2024-03-21 00:38:05",
        "updated_at": "2024-03-21 00:40:32"
    },
    {
        "id": 3,
        "title": "Test List",
        "description": "dalah sebuah film horor Indonesia tahun 2022 yang disutradarai dan ditulis oleh Joko Anwar sebagai sekuel dari film tahun 2017, Pengabdi Setan.",
        "rating": 7,
        "image": "",
        "created_at": "2024-03-21 00:42:59",
        "updated_at": "2024-03-21 00:42:59"
    },
    {
        "id": 4,
        "title": "Avatar",
        "description": "Avatar Of Ang",
        "rating": 7.6,
        "image": "",
        "created_at": "2024-03-21 02:37:22",
        "updated_at": "2024-03-21 02:37:22"
    }
]
```

- **Detail Movie**: `GET /movie/:id`
## Detail Movies

### Response Example

```json
{
    "id": 4,
    "title": "Avatar",
    "description": "Avatar Of Ang",
    "rating": 7.6,
    "image": "",
    "created_at": "2024-03-21 02:37:22",
    "updated_at": "2024-03-21 02:37:22"
}
```

- **Delete Movie**: `DELETE /movie/:id`
## Delete Movie

## Response Example

```json
{
    "message": "delete movie success"
}
```

- **Update Movie**: `PATCH /movie/:id`
## Update Movie

### Request Body

```json
{
    "title": "Avatar",
    "description": "Avatar Of Ang",
    "rating": 7.6,
    "image": ""
}
```
### Response Example

```json
{
    "id": 4,
    "title": "Avatar",
    "description": "Avatar Of Ang",
    "rating": 7.6,
    "image": "",
    "created_at": "2024-03-21 02:37:22",
    "updated_at": "2024-03-21 02:37:22"
}
```

# Menjalankan Aplikasi dengan Docker Compose.

Jalankan perintah berikut untuk memulai kontainer Docker sesuai dengan konfigurasi yang ada:
   
   ```bash
   docker-compose up
   ```
atau

running background

```bash
  docker-compose up -d
```
