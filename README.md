# Rinha de Backend 2023

This repository contains the solution to the challenge <a herf="https://github.com/zanfranceschi/rinha-de-backend-2023-q3/blob/main/INSTRUCOES.md">rinha-de-backend-2023</a>.

## Routes

### Route to register a new person
```bash
POST {host}/pessoas
```

#### Request
```json
{
    "id": "075c196d-c38d-4715-8a16-03e2efdbe988",
    "nome": "José Roberto",
    "apelido": "josé",
    "nascimento": "2000-10-01T00:00:00-03:00",
    "stack": [
        "C#",
        "Node",
        "Oracle"
    ]
}
```

### Route to search for a person by id
```bash
GET {host}/pessoas/{id}
```

#### Request
```bash
http://localhost:3000/pessoas/075c196d-c38d-4715-8a16-03e2efdbe988
```

#### Response
```json
{
    "id": "075c196d-c38d-4715-8a16-03e2efdbe988",
    "nome": "José Roberto",
    "apelido": "josé",
    "nascimento": "2000-10-01T00:00:00-03:00",
    "stack": [
        "C#",
        "Node",
        "Oracle"
    ]
}
```

### Route to search for people by terms
```bash
GET {host}/pessoas
```

#### Request
```bash
http://localhost:3000/pessoas?t=node
```

#### Response
```json
[
    {
        "id": "075c196d-c38d-4715-8a16-03e2efdbe988",
        "nome": "José Roberto",
        "apelido": "josé",
        "nascimento": "2000-10-01T00:00:00-03:00",
        "stack": [
            "C#",
            "Node",
            "Oracle"
        ]
    },
    {
        "id": "83de8978-d543-440e-a3e3-d6c1d5a925d5",
        "nome": "Lucas Emanuel",
        "apelido": "luquinhas",
        "nascimento": "2005-10-21T00:00:00-03:00",
        "stack": [
            "PHP",
            "Node",
            "MySQL",
            "MongoDB"
        ]
    },
    {
        "id": "864897f7-037d-4f4c-9be7-a4db20446bbc",
        "nome": "José Mateus",
        "apelido": "ponto e vírgula",
        "nascimento": "2005-10-01T00:00:00-03:00",
        "stack": [
            "JS",
            "Node",
            "Oracle"
        ]
    },
    {
        "id": "d58edc37-751d-4786-936a-ffacf417099c",
        "nome": "Zé Roberto",
        "apelido": "zé",
        "nascimento": "2001-10-04T00:00:00-03:00",
        "stack": [
            "C#",
            "Node",
            "Oracle",
            "Docker",
            "Linux"
        ]
    }
]
```

### Route to count the number of people stored
```bash
GET {host}/contagem-pessoas
```

#### Request
```bash
http://localhost:3000/contagem-pessoas
```

#### Response
```json
"quantidade": 4
```