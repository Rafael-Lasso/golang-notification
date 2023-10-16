
 <div align="center" class="imgs">
<img height="120em" src="https://cdn.worldvectorlogo.com/logos/twilio.svg">
</div>

## API

A Api escolhida foi Api <a style="text-decoration:none" href="https://www.twilio.com/"><b style="color:#FB8D12;cursor:pointer;"><i>Twilio</i></b></a> pois possui facil integração com mais diversos sistemas, suporte e uma documentação completa, alem de ser ultilisada por grandes empresas como
<br>

-   Nubank
-   Mercado Livre
-   Dasa
-   Inter
-   Rappi

Não é um serviço qualquer sem confiança, por isso sua escolha

### Variaveis de Ambiente

Para poder rodar a aplicação é de estrema importancia

-   Criar um arquivo <b><i>.env</i></b>
    (Windows)

            type nul > ".env"

-   Criar um arquivo <b><i>.env</i></b>
    (Linux)

            cat > .env

-   Criar um arquivo <b><i>.env</i></b>
    (Mac OS)

            touch .env

-   Escrever as Variaveis da seguinte forma ( SMS )

          ACCOUNT_SID=SeuSidGeradoDaApi
          AUTH_TOKEN=SeuTockenGeradoDaApi
          TWILIO_TO=+NumRemetente (Codigo Nacional + Numero)
          TWILIO_FROM=+SeuNumeroGeradoDaApi

-   Escrever as Variaveis da seguinte forma ( WHATSAPP )

          WHATSAPP_FROM=:+WhatsappDoRemetente
          WHATSAPP_TO=+WhatsappDoDestinatario

-   Escrever as Variaveis da seguinte forma ( EMAIL )

          SEND_FROM_NAME=NomeDoRemetente
          SEND_FROM_ADDRESS=<EmailDoRemetente>
          SEND_TO_NAME=NomeDoDestinatario
          SEND_TO_ADDRESS=<EmailDoDestinatario>