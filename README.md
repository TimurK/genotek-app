## Приложение конвертер Unix Timestamp Epoch использующее Linux утилиту `date`.

Принимает через Query string параметр date в формате (прим. 1984-04-11) по пути /epoch.

Пример:

```
curl -v -X GET --location http://localhost/epoch?date=1984-04-11
```

Для запуска необходим Docker с плагином Compose.

Запуск:

```
docker compose -p genotek up -d
```
