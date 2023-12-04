**Сервис, предоставляющий API по созданию сокращённых ссылок.**

## Оглавление

**[1 Описание сервиса](#1)**.   
**[2 Запуск](#1)**.     
**[3 API](#2)**.    
&emsp;**[3.1: Сохранить оригинальный URL](#2.1)**.    
&emsp;**[3.2: Получить оригинальный URL](#2.2)**. 

## <a name="1">1 Описание сервиса</a>

Cервис предоставляет API по созданию сокращённых ссылок. Сервис возвращает уникальные ссылки для каждого полученного URL. 
Все ссылки имеют длину в 10 символов и состоят из символов латинского алфавита в нижнем и верхнем регистре, цифр и символа _ (подчеркивание).
Сервис написан на языке Go и работает через GRPC. В качестве хранилища сервис использует in-memory решение с map или postgres на выбор.
Хранилище задается при запуске сервиса.

## <a name="2">2 Запуск</a>

Для запуска используется команда.
```
go run cmd/main.go -s postgres/in-memory
```
postgres/in-memory - выбор хранилища. Также этот параметр можно передавать через переменные окружения с ключом STORAGE. 
Если параметр передается через переменные окружения, то флаги не считываются. Поддерживается .env файл. При выборе postgres 
хранилища необходимо также передать POSTGRES_PASSWORD.


## <a name="3">ЧАСТЬ 3:API</a>
Для сервиса реализованы 2 метода. Метод SaveLink сохраняет оригинальную ссылку и возвращает сокращенную. GetLink соответственно принимает сокращенную ссылку 
и возвращает оригинальную.
```protobuf
syntax = "proto3";

option go_package = "api/shortened_links";

service shortenedLinks {
  rpc SaveLink(OriginalLink) returns (ShortenedLink) {}
  rpc GetLink(ShortenedLink) returns (OriginalLink) {}
}

message OriginalLink {
  string message = 1;
}

message ShortenedLink {
  string message = 1;
}
```

## DockerHub

https://hub.docker.com/repository/docker/ivankonoplich/links-service/general

## Contact

Ivan Konoplich - konoplich_i@mail.ru

Project Link: [https://github.com/IvanKonoplich/shortened_links_service.git](https://github.com/IvanKonoplich/shortened_links_service.git)
