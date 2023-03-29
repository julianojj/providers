# providers
Um cliente de e-mail que suporte vários provedores

# Introdução
Nesse desafio você tem a responsabilidade de criar uma biblioteca que contem um cliente de e-mail que suporta diferentes provedores/providers. O desenvolver que utilizar seu cliente pode definir qualquer ordem de providers e na hora de enviar, se o primeiro provider da lista falhar, então o segundo será usado e assim sucessivamente até que um provider consiga disparar o email.

# Requisitos
- Crie um cliente de email em Go que receberá os seguintes provedores:
    - Sendgrid
    - AWS SES
- Provider "Null" que retorna sempre verdadeiro simulando um envio
- O cliente aceita os providers em qualquer ordem
- O cliente utiliza cada provider em ordem até que o e-mail seja disparado por um deles.
- Tenha em mente que no futuro novos providers podem ser adicionados ao cliente.

# Design patterns
Foi aplicando 2 padrões de projetos, adapters e factory method
