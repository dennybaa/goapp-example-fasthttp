# Task

1. Написать web-сервер на Go/Python:

    1.1 При **GET** запросе на эндпойнт **/hello** сервер должен отдавать текст “Hello Page”

    1.2 При **POST** запросе на эндпойнт **/user** с параметром **name=\*имя пользователя\*** (пример: **.../user name=Joe**) сервер должен сохранять в текстовых лог-файл имя пользователя и время запроса (формат - **name: hh:mm:ss - dd.mm.yyyy**)

    1.3 Добавить возможность опционально использования SQLite вместо файла лога для сохранения имени пользователя и времени в таблицу **users** и возвращать эти данные по **GET** запросу с теми же параметрами (пример: **.../user?name=Joe**)

    1.4 Добавить эндпойнт **/metrics** который в prometheus-совместимом формате будет отдавать метрики по количеству обработанных **GET** и **POST** запросов

2. Написать helm-chart для вашего приложения.

    2.1 Чарт должен создавать service, deployment c томом для сохранения лог-файла или SQLite db-файла

    2.2 Возможность опционально указать селектор ноды на которую необходимо выкатить приложение

    2.3 Возможность опционально добавить serviceMonitor для сбора метрик prometheus-сервером

# Result

## Usage

### ;TL;DR; local

```bash
git clone github/dennybaa/goapp-example-fasthttp
cd goapp-example-fasthttp
go mod tidy
go build
./app &; sleep 3;

echo "============================="
curl -v http://localhost:8080/hello
echo -e "\n============================="
curl -v -XPOST "http://localhost:8080/user?name=John Doe"
```

### Environment

| Variable | Description | Default |
| - | - | - |
| ENVIRONMENT | Specifies environment  (production/development) | `production`|
| BACKEND | Specifies backend mode (logfile/sqlite) | `logfile` |
| FILEPATH | Data file store path, depends on the backend | For sqlite `data.db`, for logfile `app.log` |
| LOGLEVEL | Specifies the log level for application logger | `warn` |
| PORT | Specifies port to listen on | `8080` |

### Goals

Apart from the defined requirements this demo has an enducational focus. After implementing this code I got practiced with several popular goglang projects:

* [fasthttp](https://github.com/valyala/fasthttp) server
* [viper](https://github.com/spf13/viper) environment parser
* [zap](https://github.com/uber-go/zap) logger
* [beegoo/orm](https://github.com/beego/beego) a popular ORM with several backends


### Requirements clarifications

1. Go Web server
    
    1.2. **logfile mode (BACKEND=logfile)**

    Evidently GET /user has not been reuested thus returns `Not Found`.
    
    1.3. **sqlite mode (BACKEND=sqlite)**
    
    Evidently neither format has been requested for a GET response nor for the timestamp, thus the response is json, as bellow:
    ```json
    {"Id":107,"Name":"John Doe","AccessedAt":"2023-07-20T14:50:16.298093293Z"}
    ```
    
    1.4. */metrics endpoint* 
    
    Serves request counter and complementary request durarion:
    * http_requests_total (count)
    * http_requests_duration (histogram)
