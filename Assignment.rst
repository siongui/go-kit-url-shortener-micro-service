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
