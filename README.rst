==========================
URL Shortener Microservice
==========================

.. image:: https://img.shields.io/badge/Language-Go-blue.svg
   :target: https://golang.org/

.. image:: https://godoc.org/github.com/siongui/go-kit-url-shortener-micro-service?status.svg
   :target: https://godoc.org/github.com/siongui/go-kit-url-shortener-micro-service

.. image:: https://github.com/siongui/go-kit-url-shortener-micro-service/workflows/ci/badge.svg
    :target: https://github.com/siongui/go-kit-url-shortener-micro-service/blob/master/.github/workflows/ci.yml

.. image:: https://goreportcard.com/badge/github.com/siongui/go-kit-url-shortener-micro-service
   :target: https://goreportcard.com/report/github.com/siongui/go-kit-url-shortener-micro-service

.. image:: https://img.shields.io/badge/license-Unlicense-blue.svg
   :target: https://github.com/siongui/go-kit-url-shortener-micro-service/blob/master/UNLICENSE


URL shortener microservice using `Go kit`_, Docker_, and `Amazon EKS`_.

Development Environment:

  - `Ubuntu 20.04`_
  - `Go 1.17`_


Description
+++++++++++

See `Assignment <doc/Assignment.rst>`_.


Usage
+++++

First install ``Docker Engine`` and `Docker Compose`_.
To run the micro-service in Docker_ environment:

.. code-block:: bash

  $ sudo docker-compose up --build

To create short URL:

.. code-block:: bash

  $ curl -XPOST -d'{"url":"https://github.com/siongui/go-kit-url-shortener-micro-service"}' localhost:8080/create
  {"short_url":"20wogJaCuRtH0U5p60LvjZoVwsz"}

To recover the original URL:

.. code-block:: bash

  $ curl -XGET localhost:8080/20wogJaCuRtH0U5p60LvjZoVwsz
  {"url":"https://github.com/siongui/go-kit-url-shortener-micro-service"}

If you HTTP GET some URL that does not exist, you will get:

.. code-block:: bash

  $ curl -XGET localhost:8080/20wogJaCuRtH0U5p
  {"url":"","err":"sql: no rows in result set"}

To see metrics of the the micro-service:

.. code-block:: bash

  $ curl -XGET localhost:8080/metrics


Service Configuration
+++++++++++++++++++++

Possible service configuration in local and deployment environment.

Local Development Environment
-----------------------------

- URL shortener + SQLite: 1 container, including both.
- URL shortener + PostgresSQL: 2 container, one for URL shortener, the other for
  PostgresSQL

Deployment Environment
----------------------

- URL shortener + SQLite: Deployed on Amazon EKS.
- URL shortener + PostgresSQL: To be studied.
- URL shortener + `Amazon RDS`_: To be studied.


Deployment
++++++++++

`Deploy to Amazon EKS <doc/deployEKS.rst>`_


UNLICENSE
+++++++++

Released in public domain. See UNLICENSE_.


References
++++++++++

.. [1] | `golang microservice framework - Google search <https://www.google.com/search?q=golang+microservice+framework>`_
       | `golang microservice framework - DuckDuckGo search <https://duckduckgo.com/?q=golang+microservice+framework>`_
       | `golang microservice framework - Ecosia search <https://www.ecosia.org/search?q=golang+microservice+framework>`_
       | `golang microservice framework - Qwant search <https://www.qwant.com/?q=golang+microservice+framework>`_
       | `golang microservice framework - Bing search <https://www.bing.com/search?q=golang+microservice+framework>`_
       | `golang microservice framework - Yahoo search <https://search.yahoo.com/search?p=golang+microservice+framework>`_
       | `golang microservice framework - Baidu search <https://www.baidu.com/s?wd=golang+microservice+framework>`_
       | `golang microservice framework - Yandex search <https://www.yandex.com/search/?text=golang+microservice+framework>`_

.. [2] | `golang 微服务 - Google search <https://www.google.com/search?q=golang+%E5%BE%AE%E6%9C%8D%E5%8A%A1>`_
       | `golang 微服务 - DuckDuckGo search <https://duckduckgo.com/?q=golang+%E5%BE%AE%E6%9C%8D%E5%8A%A1>`_
       | `golang 微服务 - Ecosia search <https://www.ecosia.org/search?q=golang+%E5%BE%AE%E6%9C%8D%E5%8A%A1>`_
       | `golang 微服务 - Qwant search <https://www.qwant.com/?q=golang+%E5%BE%AE%E6%9C%8D%E5%8A%A1>`_
       | `golang 微服务 - Bing search <https://www.bing.com/search?q=golang+%E5%BE%AE%E6%9C%8D%E5%8A%A1>`_
       | `golang 微服务 - Yahoo search <https://search.yahoo.com/search?p=golang+%E5%BE%AE%E6%9C%8D%E5%8A%A1>`_
       | `golang 微服务 - Baidu search <https://www.baidu.com/s?wd=golang+%E5%BE%AE%E6%9C%8D%E5%8A%A1>`_
       | `golang 微服务 - Yandex search <https://www.yandex.com/search/?text=golang+%E5%BE%AE%E6%9C%8D%E5%8A%A1>`_
       | `如何使用 Go kit 工具包编写微服务  - Go语言中文网 - Golang中文社区 <https://studygolang.com/articles/21378>`_

.. [3] `Microservices in Go <https://medium.com/seek-blog/microservices-in-go-2fc1570f6800>`_
       (`Chinese translation <https://learnku.com/go/t/36973>`__)

.. [4] | `GO-TAMBOON ไปทำบุญ <https://github.com/siongui/tamboongo>`_
       | `Go URL shortener <https://github.com/siongui/goshorturl>`_
       | `Go Employee Open API <https://github.com/siongui/go-employee-api>`_

.. [5] | `golang generate unique id - Google search <https://www.google.com/search?q=golang+generate+unique+id>`_
       | `golang generate unique id - DuckDuckGo search <https://duckduckgo.com/?q=golang+generate+unique+id>`_
       | `golang generate unique id - Ecosia search <https://www.ecosia.org/search?q=golang+generate+unique+id>`_
       | `golang generate unique id - Qwant search <https://www.qwant.com/?q=golang+generate+unique+id>`_
       | `golang generate unique id - Bing search <https://www.bing.com/search?q=golang+generate+unique+id>`_
       | `golang generate unique id - Yahoo search <https://search.yahoo.com/search?p=golang+generate+unique+id>`_
       | `golang generate unique id - Baidu search <https://www.baidu.com/s?wd=golang+generate+unique+id>`_
       | `golang generate unique id - Yandex search <https://www.yandex.com/search/?text=golang+generate+unique+id>`_
       |
       | `Generating good unique ids in Go <https://blog.kowalczyk.info/article/JyRZ/generating-good-unique-ids-in-go.html>`_

.. [6] | `go kit + gin - Google search <https://www.google.com/search?q=go+kit+%2B+gin>`_
       | `Gin与go-kit如何搭配使用？ - 知乎 <https://www.zhihu.com/question/323548694>`_
       | `examples/stringsvc5 at feature-add-gin-example · xpzouying/examples <https://github.com/xpzouying/examples/tree/feature-add-gin-example/stringsvc5>`_

.. [7] | `run golang on docker - Google search <https://www.google.com/search?q=run+golang+on+docker>`_
       | `run golang on docker - DuckDuckGo search <https://duckduckgo.com/?q=run+golang+on+docker>`_
       | `run golang on docker - Ecosia search <https://www.ecosia.org/search?q=run+golang+on+docker>`_
       | `run golang on docker - Qwant search <https://www.qwant.com/?q=run+golang+on+docker>`_
       | `run golang on docker - Bing search <https://www.bing.com/search?q=run+golang+on+docker>`_
       | `run golang on docker - Yahoo search <https://search.yahoo.com/search?p=run+golang+on+docker>`_
       | `run golang on docker - Baidu search <https://www.baidu.com/s?wd=run+golang+on+docker>`_
       | `run golang on docker - Yandex search <https://www.yandex.com/search/?text=run+golang+on+docker>`_
       |
       | `Build your Go image | Docker Documentation <https://docs.docker.com/language/golang/build-images/>`_

.. [8] | `Overview of Docker Compose | Docker Documentation <https://docs.docker.com/compose/>`_
       | `用 Docker 玩轉 PostgreSQL | My.APOLLO <https://myapollo.com.tw/zh-tw/docker-postgres/>`_
       | `Docker - 第十二章 | 安裝PostgreSQL | J.J.'s Blogs <https://morosedog.gitlab.io/docker-20190505-docker12/>`_
       | `[Docker] 於 Windows 中運行 PostgreSQL Container 來提供本機 DB 開發環境 | 搞搞就懂 - 點部落 <https://www.dotblogs.com.tw/wasichris/2020/11/13/104023>`_
       | `Play PostgreSQL with Docker. Docker 就是你的實驗室，執行記錄 | by 古哥 | pgsql-tw | Medium <https://medium.com/pgsql-tw/play-postgresql-with-docker-4dbc15d9b0d3>`_
       | `Docker筆記 - 進入Container，建立並操作 PostgreSQL Container | by Albert Hg | alberthg-docker-notes | Medium <https://medium.com/alberthg-docker-notes/docker%E7%AD%86%E8%A8%98-%E9%80%B2%E5%85%A5container-%E5%BB%BA%E7%AB%8B%E4%B8%A6%E6%93%8D%E4%BD%9C-postgresql-container-d221ba39aaec>`_
       | `docker compose postgres - Google search <https://www.google.com/search?q=docker+compose+postgres>`_
       | `docker compose postgres - DuckDuckGo search <https://duckduckgo.com/?q=docker+compose+postgres>`_
       | `docker compose postgres - Ecosia search <https://www.ecosia.org/search?q=docker+compose+postgres>`_
       | `docker compose postgres - Qwant search <https://www.qwant.com/?q=docker+compose+postgres>`_
       | `docker compose postgres - Bing search <https://www.bing.com/search?q=docker+compose+postgres>`_
       | `docker compose postgres - Yahoo search <https://search.yahoo.com/search?p=docker+compose+postgres>`_
       | `docker compose postgres - Baidu search <https://www.baidu.com/s?wd=docker+compose+postgres>`_
       | `docker compose postgres - Yandex search <https://www.yandex.com/search/?text=docker+compose+postgres>`_
       | `Creating and filling a Postgres DB with Docker compose | by José David Arévalo | Level Up Coding <https://levelup.gitconnected.com/creating-and-filling-a-postgres-db-with-docker-compose-e1607f6f882f>`_

.. [9] | `run postgres on github actions - Google search <https://www.google.com/search?q=run+postgres+on+github+actions>`_
       | `run postgres on github actions - DuckDuckGo search <https://duckduckgo.com/?q=run+postgres+on+github+actions>`_
       | `run postgres on github actions - Ecosia search <https://www.ecosia.org/search?q=run+postgres+on+github+actions>`_
       | `run postgres on github actions - Qwant search <https://www.qwant.com/?q=run+postgres+on+github+actions>`_
       | `run postgres on github actions - Bing search <https://www.bing.com/search?q=run+postgres+on+github+actions>`_
       | `run postgres on github actions - Yahoo search <https://search.yahoo.com/search?p=run+postgres+on+github+actions>`_
       | `run postgres on github actions - Baidu search <https://www.baidu.com/s?wd=run+postgres+on+github+actions>`_
       | `run postgres on github actions - Yandex search <https://www.yandex.com/search/?text=run+postgres+on+github+actions>`_
       |
       | `Creating PostgreSQL service containers - GitHub Docs <https://docs.github.com/en/actions/using-containerized-services/creating-postgresql-service-containers>`_

.. [10] | `Microservices using Go, Docker, RabbitMQ, Redis, AWS, CI/CD : golang <https://old.reddit.com/r/golang/comments/qy8h7j/microservices_using_go_docker_rabbitmq_redis_aws/>`_
        | `GitHub - ebosas/microservices: A microservices example in Go <https://github.com/ebosas/microservices>`_

.. [11] | `golang short uuid url friendly - Google search <https://www.google.com/search?q=golang+short+uuid+url+friendly>`_
        | `golang short uuid url friendly - DuckDuckGo search <https://duckduckgo.com/?q=golang+short+uuid+url+friendly>`_
        | `golang short uuid url friendly - Ecosia search <https://www.ecosia.org/search?q=golang+short+uuid+url+friendly>`_
        | `golang short uuid url friendly - Qwant search <https://www.qwant.com/?q=golang+short+uuid+url+friendly>`_
        | `golang short uuid url friendly - Bing search <https://www.bing.com/search?q=golang+short+uuid+url+friendly>`_
        | `golang short uuid url friendly - Yahoo search <https://search.yahoo.com/search?p=golang+short+uuid+url+friendly>`_
        | `golang short uuid url friendly - Baidu search <https://www.baidu.com/s?wd=golang+short+uuid+url+friendly>`_
        | `golang short uuid url friendly - Yandex search <https://www.yandex.com/search/?text=golang+short+uuid+url+friendly>`_
        |
        | `GitHub - matoous/go-nanoid: Golang random IDs generator. <https://github.com/matoous/go-nanoid>`_


.. _Go: https://golang.org/
.. _Ubuntu 20.04: https://releases.ubuntu.com/20.04/
.. _Go 1.17: https://golang.org/dl/
.. _UNLICENSE: https://unlicense.org/
.. _Go kit: https://gokit.io/
.. _Docker: https://www.docker.com/
.. _Docker Compose: https://docs.docker.com/compose/
.. _Amazon EKS: https://aws.amazon.com/eks/
.. _Amazon RDS: https://aws.amazon.com/rds/
