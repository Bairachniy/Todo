#Todo server request description
##Todo
### Create todo:

>Method: POST
>
>Path: **/todos**
>
>Body

```json5
{
  // id generating automaticaly uuid.v4 
  "name": "Your's todo" 
}
```

### Delete todo:

>Method DELETE
>
>Path **todos/{id}**
>
### Uptade todo:

>Method POST
>
>Path **todos/{id}**
>
>Body

```json5
{
  "name": "New todo"
}
```

### Get All:

>Method GET
>
>Path **/todos**
