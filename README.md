# inDriwer test task
Original googledock: https://docs.google.com/document/d/1QNWc1Ijc-Xr0bi7Ngcg0CLp-D3_V9kLbj9piuibUhIA/edit
## Actual task
Необходимо реализовать простенькое CRUD API
на любую сущность (допустим на добавление моделей автомобиля, поля фирма, марка, литраж).
Для реализации необходимо:
- Покрыть код Unit тестами
- Использовать golang 1.13
- Использовать реляционную СУБД
Допы (это - необязательная часть задания но будет плюсом):
- Контейнеризировать приложение используя оркестраторы (docker-compose, kuber на выбор)
- Создать коллекцию к апи, доступ через гуи (Postman, Swagger)
- Написать интеграционные тесты
- Кэшировать данные (memcached, redis, aerospike и тд)
Допы реализовывать не обязательно, они являются небольшим усложнением поставленной задачи и не влияют на общее представление о качестве кода, полученного в результате выполнения обязательной части задачи.

Working plan
- create README
- write working plan
- make docker-compose (go, postgres, redis)
- make application blueprint (gin, swaggo, gin-swagger, gorm)
- define gorm models
- make migrations
- define gin routers and controllers
- write swaggo comments
- add swagger route
- add redis for cache
- mb some other staff
- write integrational tests