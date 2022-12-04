# dbexport
Pacote para exportar e sincronizar de forma fácil os objetos do seu banco de dados.

## Como usar

* Criar o arquivo `.env.local` com os mesmos dados do `.env.local.example`
* Caso queira usar em ambiente diferentes você pode criar arquivos para cada ambiente ex `.env.local`, `.env.prod`, `.env.test` com os mesmos dados do `.env.local.example`

#### rodando

`./dbexport all`

`./dbexport procedures`

`./dbexport triggers tg_pessoa_ins_after`

`./dbexport observe -t 5`

Rodando para um ambiente específico

`DB_ENV=prod ./dbexport all`
