# Weather API
Этот проект представляет собой простой API для получения данных о погоде с помощью OpenWeatherMap API на основе указанного названия города. Он использует язык программирования Go и пакет net/http для обработки HTTP-запросов и ответов.

# Функциональность
API предоставляет следующую функциональность:

Получение данных о погоде: API позволяет клиентам отправлять GET-запросы к конечной точке /weather с параметром запроса city для получения текущих данных о погоде для указанного города.
JSON-ответ: API возвращает данные о погоде в формате JSON, включая информацию о температуре, влажности, скорости ветра и облачности.
Установка и использование
Клонируйте репозиторий.
Замените переменную apiKey в функции weatherHandler на ваш собственный ключ OpenWeatherMap API.
Соберите и запустите проект с помощью команды go run main-api.go.
Отправьте GET-запрос на http://localhost:5555/weather?city={название_города} для получения данных о погоде для указанного города.
# Зависимости
Этот проект имеет следующие зависимости:

Язык программирования Go (версия 1.16 или выше)
Пакет net/http
Пакет encoding/json
# Внесение вклада
Вклад в этот проект приветствуется. Пожалуйста, следуйте стандартным рекомендациям для внесения вклада, включая отправку сообщений об ошибках, запросов новых функций и запросов на объединение изменений.

# Лицензия
Этот проект лицензирован в соответствии с лицензией MIT.

# Авторы
DrBuzzer

# Контактная информация
По всем вопросам и запросам обращайтесь по адресу your-email@example.com.
# Weather-API
Weather API
This project is a simple weather API that retrieves weather data from the OpenWeatherMap API based on the provided city name. It uses the Go programming language and the net/http package to handle HTTP requests and responses.

# Functionality
The API provides the following functionality:

Weather Retrieval: The API allows clients to send a GET request to the /weather endpoint with a city query parameter to retrieve the current weather data for the specified city.
JSON Response: The API returns the weather data in JSON format, including information such as temperature, humidity, wind speed, and cloud cover.
# Installation and Usage
Clone the repository.
Replace the apiKey variable in the weatherHandler function with your own OpenWeatherMap API key.
Build and run the project using the go run main-api.go command.
Send a GET request to http://localhost:5555/weather?city={city_name} to retrieve the weather data for a specific city.
# Dependencies
This project has the following dependencies:

Go programming language (version 1.16 or higher)
net/http package
encoding/json package
# Contributing
Contributions to this project are welcome. Please follow the standard guidelines for contributing, including submitting bug reports, feature requests, and pull requests.

# License
This project is licensed under the MIT License.

# Authors
DrBuzzer

# Contact
For any questions or inquiries, please contact kuzminskk@yandex.ru.
