### About
Репозиторий с решением тестового задания в школу Backend от Xsolla (https://github.com/xsolla/xsolla-school-backend-2021)

### Развертывание
Разворачивать решение рекомендуется с помощью Docker. Для этого нужно:
1. Установить [Docker](https://docs.docker.com/get-docker/), если еще не установлен
1. Выкачать репозиторий
1. Открыть терминал в корне репозитория и выполнить команду `docker-compose up -d`

В результате будут запущены 3 Docker контейнера:
- Контейнер с БД PostgreSQL (по умолчанию логин `postgres` и пароль `postgres`)
- Контейнер с серверным приложением (которое взаимодействует с БД в вышеупомянутом контейнере)
- Контейнер с pgAdmin 4 (логин `user@domain.com`, пароль `1q2w3e`)

_Если нужно, чтобы приложение подключалось к другой БД, можно создать в корне файл `.env` и указать в нем `POSTGRES_HOST`, `POSTGRES_DB`, `POSTGRES_USERNAME` и `POSTGRES_PASSWORD`_

### Описание решения
Что сделано:
- [x] Обязательная часть.
- [ ] Дополнительная часть:
   - [ ] Фильтрация товаров по их типу и/или стоимости в методе получения каталога товаров.
   - [x] Спецификация OpenAPI (https://app.swaggerhub.com/apis/egorkartashov/Products_API/1.0.0#/).
   - [x] Поддержка Docker (Dockerfile и docker-compose файл).
   - [ ] Модульные и функциональные тесты.
   - [ ] Развертывание приложения на любом публичном хостинге, например, на heroku.

