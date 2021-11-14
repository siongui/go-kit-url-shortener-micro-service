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


URL shortener microservice using `Go kit`_.

Development Environment:

  - `Ubuntu 20.04`_
  - `Go 1.17`_


Description
+++++++++++

Use gin_ to design a RESTful API with following endpoints:

- ``POST`` ``/create`` to create short URL for given URL.
- ``GET`` ``/:shorturl`` to return the original URL.

Use JSON for data exchange in the RESTful API. Using search engines for help is
encouraged, and the code/algorithm/help from search engines must be clearly
specified in the code or documentation.

Requirement
-----------

- Design an algorithm to convert given URL to short URL.
  Test code for the algorithm.
- Use gin_ to implement the API endpoints.
- Use curl_ to access the API endpoints.
- `Go module`_ support: ``go.mod`` and ``go.sum``.
- ``go fmt`` Go source code.
- bun_ + SQLite for database. Test code for database operation.
- README to explain:

  * Development environment (Go version, Windows/Linux/Mac, etc.)
  * The algorithm description of URL shortener.
  * How to run the code.
  * The curl_ examples to access the API endpoints.
  * References or anything you think worth telling.

- Commit all the files to public GitHub repository.

Bonus
-----

- Use GitHub Actions for testing.
- Documentation on `pkg.go.dev`_ and a badge link in README.
- Get A+ in `Go report card`_ and a badge link in README.
- Run the code in docker_ container and instructions in README.
- Use bun_ + PostgreSQL
- Use logrus_ for logging
- Use Go micro-service framework, like `Go kit`_ or Kratos_, to implement URL
  shortener as micro-service.
- Anything you think worth the points.


Usage
+++++

To run the micro-service:

.. code-block:: bash

  $ go run instrumenting.go logging.go main.go service.go transport.go

To create short URL:

.. code-block:: bash

  # Algorithm not implemented yet. Return only original URL.
  $ curl -XPOST -d'{"url":"https://github.com/siongui/go-kit-url-shortener-micro-service"}' localhost:8080/create

To see metrics of the the micro-service:

.. code-block:: bash

  $ curl -XGET localhost:8080/metrics


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


.. _Go: https://golang.org/
.. _Ubuntu 20.04: https://releases.ubuntu.com/20.04/
.. _Go 1.17: https://golang.org/dl/
.. _UNLICENSE: https://unlicense.org/
.. _Go kit: https://gokit.io/
.. _gin: https://github.com/gin-gonic/gin
.. _curl: https://curl.se/
.. _Go module: https://golang.org/doc/tutorial/create-module
.. _bun: https://github.com/uptrace/bun
.. _pkg.go.dev: https://pkg.go.dev/
.. _Go report card: https://goreportcard.com/
.. _docker: https://www.docker.com/
.. _logrus: https://github.com/sirupsen/logrus
.. _Kratos: https://go-kratos.dev/
