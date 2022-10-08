# Kuliner API

Dokumen ini berisi penjelasan mengenai Kuliner API.

Kuliner API adalah API sederhana yang digunakan untuk melakukan CRUDS (Create, Read, Update, Delete, & Search) pada katalog kuliner Indonesia.

Endpoint-endpoint yang terdapat pada API ini:

- [Indeks Kuliner](#indeks-kuliner)
- [Hapus Kuliner](#hapus-kuliner)
- [Cari Kuliner](#cari-kuliner)

---

## Indeks Kuliner

POST: `/foods`

Endpoint ini digunakan untuk mengindeks data kuliner ke dalam database.


**Request Body:**

- `name`, String => nama dari kuliner yang akan diindeks, jika nama kuliner sudah ada di database, maka data dari kuliner yang sudah ada tersebut akan ditimpa dengan data yang baru.
- `description`, String => deskripsi dari kuliner yang akan diindeks.

**Example Request:**

```json
POST /foods
Content-Type: application/json

{
"name": "Bakso Istighfar",
"description": "Bakso sebesar bola voli yang banyak tersedia di daerah Bandung. Saking besarnya bakso ini, orang yang melihat sampai harus ber-istighfar, karena itulah bakso ini dinamakan sebagai bakso istighfar."
}
```
**Success Response:**

```json
HTTP/1.1 200 OK
Content-Type: application/json

{
    "ok": true,
    "data": {
        "id": "bakso-istighfar",
        "name": "Bakso Istighfar",
        "description": "Bakso sebesar bola voli yang banyak tersedia di daerah Bandung. Saking besarnya bakso ini, orang yang melihat sampai harus ber-istighfar, karena itulah bakso ini dinamakan sebagai bakso istighfar."
    }
}
```

**Error Responses:**

- Bad Request (`400`)

    ```json
    HTTP/1.1 400 Bad Request
    Content-Type: application/json

    {
        "ok": false,
        "err": "ERR_BAD_REQUEST",
        "msg": "missing `name`"
    }
    ```

    Client akan menerima error ini jika ada parameter yang tidak terisi ketika melakukan request.

[Back to Top](#kuliner-api)

---

## Hapus Kuliner

DELETE: `/foods/{food_id}`

Endpoint ini digunakan untuk menghapus data suatu kuliner dari database.

**Example Request:**

```json
DELETE /foods/bakso-istighfar
```

**Success Response:**

```json
HTTP/1.1 200
Content-Type: application/json

{
    "ok": true
}
```

**Error Responses:**

- Not Found (`404`)

    ```json
    HTTP/1.1 404 Not Found
    Content-Type: application/json

    {
        "ok": false,
        "err": "ERR_NOT_FOUND",
        "msg": "data is not found"
    }
    ```

[Back to Top](#kuliner-api)

---

## Cari Kuliner

GET: `/foods?q={query}`

Endpoint ini akan mengembalikan maksimum 10 kuliner yang paling relevan dengan query yang diberikan.

```json
GET /foods?q=Bakso
```

**Success Response:**

```json
HTTP/1.1 200 OK
Content-Type: application/json

{
    "ok": true,
    "data": {
        "query": "Bakso",
        "foods": [
            {
                "id": "bakso-istighfar",
                "name": "Bakso Istighfar",
                "description": "Bakso sebesar bola voli yang banyak tersedia di daerah Bandung. Saking besarnya bakso ini, orang yang melihat sampai harus ber-istighfar, karena itulah bakso ini dinamakan sebagai bakso istighfar."
            }
        ]
    }
}
```

**Error Responses:**

Tidak ada response error spesifik

[Back to Top](#kuliner-api)

---

## Update Kuliner

PUT: `/foods/{food_id}`

Endpoint ini digunakan untuk mengubah data suatu kuliner dari database.

**Request Body:**

- `name`, String => nama dari kuliner yang akan diindeks, jika nama kuliner sudah ada di database, maka data dari kuliner yang sudah ada tersebut akan ditimpa dengan data yang baru.
- `description`, String => deskripsi dari kuliner yang akan diindeks.

**Example Request:**

```json
PUT /foods/bakso-istighfar

Content-Type: application/json

{
  "name": "Bakso Istighfar",
  "description": "Lorem ipsum dolor sit amet"
}
```

**Success Response:**

```json
HTTP/1.1 200
Content-Type: application/json

{
    "ok": true,
    "data": {
      "id": "bakso-istighfar",
      "name": "Bakso istigfar",
      "description": "Lorem ipsum dolor sit amet"
    }
}
```

**Error Responses:**

- Not Found (`404`)

    ```json
    HTTP/1.1 404 Not Found
    Content-Type: application/json

    {
        "ok": false,
        "err": "ERR_NOT_FOUND",
        "msg": "data is not found"
    }
    ```

[Back to Top](#kuliner-api)

---