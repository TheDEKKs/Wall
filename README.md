# Wall Info
The project for to create a wall, where your friends could sign/comment. 

You can create account and wall, write comments, edit comment and wall, see all comments user and comments wall. 


## Tools
**Programming language:** *Golang* 

**Backend:** *Gin* 

**DB:** *PostgresSQL, Gorm* 

**Hach:** *Redis* 

**Else:** *Docker, Godotenv* 

## Install

```
git clone https://github.com/TheDEKKs/Wall.git
cd Wall
```

Now create new file ".env", and paste:
```
secretKey = "you secter key"
DATABASE_URL = "postgres://postgres:2242@postgres/mydb"
```

After which, start project command 
```
sudo dcoker-compose up --build
```

## Requests
| Type | URL | Desription |
| --- | --- | --- |
| GET |  /wall/:id | Open wall |
| GET | /searchallcomment | Search all comment |
| PUT | /wall/editcomment | Update comment |
| PUT | /wall/editwall | Update wall |
| POST | /wall/newcomment | New comment |
| POST | /login | Login in app |


- /wall/newcomment
  
  JSON Request:
 ```
    "token": "null" (string),
    "comment": "Text Comment" (string, not null),
    "id_wall": Id Wall (int, not null)
  ```

- /login

 JSON Request:
  ```
    "password": "Password" (string),
    "User": "User Name" (string, unique;not null),
    "ID_Telegram": Id Telegram (int, unique;not null) 
```

- /wall/editcomment

 JSON Request:
  ```
    "token": "None" (string),
    "id_comment": Id comment (int, not null),
    "id_commentor": Id creator (int, not null),
    "new_comment": "New text" (string, not null)
```

- /searchallcomment
  ```
  query parameters - id (int), hach (int)
  ```

>[!IMPORTANT]
> Gets user comments, if hach != 1 search data in hach
   

  **Example**
  > *GET* localhost:8080/searchallcomment?id=1&hach=0
  
- /wall/:id
  ```
  gets wall data
  ```
  **Example**
  > *GET* localhost:8080/wall/1
  

- /wall/editwall
  ```query parameters - mat (bool)
  update user wall, string "Mat"
  ```

  **Example**
  > *PUT* localhost:8080/wall/editwall?mat=true

>[!NOTE]
>I would be grateful if you leave your recommendation in "GitHub Issues"

