# todo-app-go

## table
```
@startuml ERDiagram
title todo-app diagram
entity "users" {
    + id [PK]
    ==
    * uuid:varchar(255)
    name:varchar(255)
    email:varchar(255)
    password:varchar(255)
    created_at:timestamp
    updated_at:timestamp
    deleted_at:timestamp
    is_deleted:boolean
}
@enduml
```