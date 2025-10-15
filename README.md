# Wall Info
The project for to create a wall, where your friends could sign/comment. 

You can create account and wall, write comments, edit comment and wall, see all comments user and comments wall. 

# UPDATES
- Create new column "Id" in JWT Token 
- Migration created with the help Goose
- New URL

## Tools
**Programming language:** *Golang* 

**Backend:** *Gin* 

**DB:** *PostgresSQL, Gorm* 

**Hach:** *Redis* 

**Else:** *Docker, Godotenv* 

## Install

```bash
git clone https://github.com/TheDEKKs/Wall.git
cd Wall
```

Now create new file ".env", and paste:
```env
secretKey = "you secter key"
DATABASE_URL = "postgres://postgres:2242@postgres/mydb"
```

After which, start project command 
```bash
sudo dcoker-compose up --build
```

## Requests
| Type | URL | Desription |
| --- | --- | --- |
| GET |  /wall/:id | Open wall |
| GET | /searchallcomment | Search all comment |
| PUT | /comment/editcomment | Update comment |
| PUT | /wall/editwall | Update wall |
| POST | /comment/newcomment  | New comment |
| POST | /au/registration | Registration in app |


- /comment/newcomment 
  
  JSON Request:
 ```JSON
    "token": "null",
    "comment": "Text Comment",
    "id_wall": 0
  ```

- /au/registration

 JSON Request:
  ```JSON
    "password": "Password",
    "User": "User Name",
    "ID_Telegram": 0 
```

- /comment/editcomment 

 JSON Request:
  ```JSON
    "token": "None",
    "id_comment": 0,
    "id_commentor": 0,
    "new_comment": "New text"
```

- /searchallcomment
  ```
  query parameters - id (int), hach (int)
  ```
  
  **Example**
  
  *GET*
  
  ```URL
  localhost:8080/searchallcomment?id=1&hach=0 


>[!IMPORTANT]
> Gets user comments, if hach != 1 search data in hach
   
  
- /wall/:id
  ```
  gets wall data
  ```

  **Example**
  
   *GET*
  ```URL
  localhost:8080/wall/1
  ```

- /wall/editwall
  ```query parameters - mat (bool)
  update user wall, string "Mat"
  ```

  **Example**

   *PUT*
  ```URL
    localhost:8080/wall/editwall?mat=true
  ```

>[!NOTE]
>I would be grateful if you leave your recommendation in "GitHub Issues"

