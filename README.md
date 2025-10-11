# Wall
> [!NOTE] 
> Hello World! This project is Pet-Project, I am would be grateful if you leave on "GitHub Issues" with advice. I hope what with understanding to my small project 

This is project backend part. Sory I can't write Fronternd:)
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
sudo dcoker-compos up --build
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
