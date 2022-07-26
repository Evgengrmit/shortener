# API приложения для укорочения URL
## Для запуска
1. Поднятие базы данных:
````
make base
````
Перед первым запуском приложения:
````
make migrate
````
2. Docker:
````
docker build .
````
3. Запуск приложения:
   (В данном пунке происходит выбор репозитория для хранения данных: оперативная память или postgresql база данных)
````
docker run --rm --network="host" -e WORKMODE=<memory или db> <app>
````

## Задача
Наобходимо реализовать сервис, который должен предоставлять API по созданию сокращённых ссылок следующего формата:
1. Ссылка должна быть уникальной и на один оригинальный URL должна ссылаться только одна сокращенная ссылка
2. Ссылка должна быть длинной 10 символов
3. Ссылка должна состоять из символов латинского алфавита в нижнем и верхнем регистре, цифр и символа «_»

## Условие
Сервис должен быть написан на Go и принимать следующие запросы по http:
1. Метод Post, который будет сохранять оригинальный URL в базе и возвращать сокращённый
2. Метод Get, который будет принимать сокращённый URL и возвращать оригинальный URL

## Результат
Решение должно быть предоставлено в «конечном виде», а именно:
1. Сервис должен быть распространён в виде Docker-образа
2. В качестве хранилища ожидается использовать две реализации. Какое хранилище использовать, указывается параметром при запуске сервиса: 1) PostgreSQL, 2) Самостоятельно написать пакет для хранения ссылок в памяти приложения
3. Покрыть реализованный функционал Unit тестами