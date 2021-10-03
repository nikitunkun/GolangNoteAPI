# GoLang Note API
This is my first **CRUD** API using GoLang

___

### Technologies
*All used packages in the file **go.mod***

- Mux
- gORM

---

### Functions
- Create a Note
- Get a list of all Notes
- Update the Note
- Delete Note

---

### API Endpoints

- Get a list of all Notes

Method: **GET**
```bash
http://localhost:8000/posts
```
- Create a new Note

Method: **POST**
```bash
http://localhost:8000//post/{title}/{author}
```
- Delete Note

Method: **DELETE**
```bash
http://localhost:8000/post/{title}
```
- Update the Note

Method: **PUT**
```bash
http://localhost:8000/post/{title}/{id}
```