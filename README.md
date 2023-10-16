<br>
<div align="center" class="imgs"> 
<img style="margin: 0 20px;" height="100em" src="https://upload.wikimedia.org/wikipedia/commons/thumb/0/05/Go_Logo_Blue.svg/800px-Go_Logo_Blue.svg.png">
</div>
<br>

<br>
<hr>

## Golang Notification

O Serviço de Notificação é escrito em codigo Go, localizado na package
<b><i style="color:#FB8D12;cursor:pointer">cmd</i></b>

- Para rodar a aplicação certifiquice de ter o Go instalado em sua maquina, ou , execute um codigo Docker para poder compilar o codigo ( Lembrando, a verção ultilizada foi a <b style="color:#FB8D12;cursor:pointer">Go 1.19</b> )

- Acesse a pasta <b><i>services</i></b> com o comando no terminal

        cd cmd

- Apos acessar o diretorio do arquivo, rode o arquivo no seu terminal com o comando

        go run main.go

- Para se realizar o <b style="color:#FB8D12;cursor:pointer"><i>Build</i></b> use o comando

        go build main.go

<br>


## Arquitetura

<div align="center" class="imgs"> 
   <img width="510rem" src="https://miro.medium.com/v2/resize:fit:800/1*0R0r00uF1RyRFxkxo3HVDg.png">
</div>

### Os Principios

<p>A arquitetura do Software é com base nos principios da Cleand Architeture, por objetivo padronizar e organizar o código desenvolvido, favorecer a sua reusabilidade, assim como independência de tecnologia.</p>


### Build
<p>
Contem o arquivo</p>

        .dockerfile

### Cmd
<p>
Contem o arquivo main do projeto</p>

        main.go

### Config
        .env
<p>
Contem o arquivo de configurações do projeto</p>

### Docs
<p>
Contem o arquivo de Documentações do projeto</p>

        *.md

### Internal
<p>
Contem os arquivo de funcionalidade interna do projeto</p>

### Pkg
<p>
Contem os arquivo de funcionalidade externa do projeto</p>

### Scripts
<p>
Contem os Scripts do projeto </p>

### Test
<p>
Contem os Testes do projeto</p>

        e2e
        env
        integration
        web