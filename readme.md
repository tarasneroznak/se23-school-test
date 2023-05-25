# Software engineering school 23

## Decisions

- `subscription.route.go:[subscribeController]` - не явно оброблені помилки, виклик методу `subscribe` може повертати не тільки конфлікти а й інші типи виключень, наприклад `permission denied` при спробі читання файлу. Не став додавати типи для `err`, так як не знаю як краще реалізувати це в golang, тому всі виключення з цього методу це - конфлікт.

## API docs

```
http://localhost:8080/api/swagger/index.html
```
