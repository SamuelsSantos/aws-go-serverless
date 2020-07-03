# SERVERLESS GO AWS

Informações do desafio

    Esse desafio vai te introduzir ao mundo serverless! 
    E por conta disso você terá que gerar um endpoint no seguinte formato: /soma?a={numero}&b={numero}.

    Quando alguém acessar através do método get um json deve ser retornado no formato:
    {"resultado":valor} 

Para realizar o desafio, fique na liberdade para escolher entre as linguagens: javascript, golang ou python. Também utilize o framework Serverless para realizar essa tarefa e utilize o cloud provider que achar mais conveniente. 


## RUN

### Setup your keys


```bash
# AWS provider sample
serverless config credentials --provider aws --key {YOUR_KEY} --secret {YOUR_SECRET}
```

### Deploy

Makefile provides some supported commands.

```bash
# Build
make build

# Clean
make clean

# Deploy
make deploy
```


