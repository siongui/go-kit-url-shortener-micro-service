Deploy Docker Compose to Amazon ECS
+++++++++++++++++++++++++++++++++++

1. Install Docker Engine and `Docker Compose`_
2. `Setting up with Amazon ECS <https://docs.aws.amazon.com/AmazonECS/latest/developerguide/get-set-up-for-amazon-ecs.html>`_
3. `Installing the AWS Command Line Interface <https://docs.aws.amazon.com/cli/latest/userguide/cli-chap-getting-started.html>`_
4. `Install the Docker Compose CLI on Linux <https://docs.docker.com/cloud/ecs-integration/#install-the-docker-compose-cli-on-linux>`_
   (Note that Docker Compose CLI is different from Docker Compose)

5. Build and push Docker image to container registry. See
   `Docker basics for Amazon ECS - Amazon Elastic Container Service <https://docs.aws.amazon.com/AmazonECS/latest/developerguide/docker-basics.html>`_
6. Create docker AWS context and ``docker compose up``. See
   | `Deploying Docker containers on ECS | Docker Documentation <https://docs.docker.com/cloud/ecs-integration/>`_
   | `Deploy applications on Amazon ECS using Docker Compose | Containers <https://aws.amazon.com/blogs/containers/deploy-applications-on-amazon-ecs-using-docker-compose/>`_
   | **IMPORTANT TRICK**: need ``sudo`` before
     ``docker context create ecs myecscontext`` or ``docker context ls``.
     Otherwise cannot create ecs context or view ecs context.



Issue
+++++

.. code-block:: bash

  ARNING services.hostname: unsupported attribute
  WARNING services.build: unsupported attribute
  service (NAME OF SERVICE) doesn't define a Docker image to run: incompatible attribute

Search the message on Google and found

- `docker compose ecs deploy aws tutorial: celery incompatible attribute · Issue #8040 · docker/compose · GitHub <https://github.com/docker/compose/issues/8040>`_
- `docker compose ecs deploy aws tutorial: celery incompatible attribute · Issue #1159 · docker/compose-cli · GitHub <https://github.com/docker/compose-cli/issues/1159>`_


.. [1] | `ecs vs kubernetes - Google search <https://www.google.com/search?q=ecs+vs+kubernetes>`_
       | `ecs vs kubernetes - DuckDuckGo search <https://duckduckgo.com/?q=ecs+vs+kubernetes>`_
       | `ecs vs kubernetes - Ecosia search <https://www.ecosia.org/search?q=ecs+vs+kubernetes>`_
       | `ecs vs kubernetes - Qwant search <https://www.qwant.com/?q=ecs+vs+kubernetes>`_
       | `ecs vs kubernetes - Bing search <https://www.bing.com/search?q=ecs+vs+kubernetes>`_
       | `ecs vs kubernetes - Yahoo search <https://search.yahoo.com/search?p=ecs+vs+kubernetes>`_
       | `ecs vs kubernetes - Baidu search <https://www.baidu.com/s?wd=ecs+vs+kubernetes>`_
       | `ecs vs kubernetes - Yandex search <https://www.yandex.com/search/?text=ecs+vs+kubernetes>`_
       |
       | `Deploying Docker containers on ECS | Docker Documentation <https://docs.docker.com/cloud/ecs-integration/>`_

.. [2] | `Deploying Docker compose on ECS - Google search <https://www.google.com/search?q=Deploying+Docker+compose+on+ECS>`_
       | `Deploying Docker compose on ECS - DuckDuckGo search <https://duckduckgo.com/?q=Deploying+Docker+compose+on+ECS>`_
       | `Deploying Docker compose on ECS - Ecosia search <https://www.ecosia.org/search?q=Deploying+Docker+compose+on+ECS>`_
       | `Deploying Docker compose on ECS - Qwant search <https://www.qwant.com/?q=Deploying+Docker+compose+on+ECS>`_
       | `Deploying Docker compose on ECS - Bing search <https://www.bing.com/search?q=Deploying+Docker+compose+on+ECS>`_
       | `Deploying Docker compose on ECS - Yahoo search <https://search.yahoo.com/search?p=Deploying+Docker+compose+on+ECS>`_
       | `Deploying Docker compose on ECS - Baidu search <https://www.baidu.com/s?wd=Deploying+Docker+compose+on+ECS>`_
       | `Deploying Docker compose on ECS - Yandex search <https://www.yandex.com/search/?text=Deploying+Docker+compose+on+ECS>`_
       |
       | `How to Deploy Containers Directly to AWS ECS using Docker <https://blog.56k.cloud/how-to-deploy-containers-directly-to-aws-ecs-using-docker/>`_
       | `Docker Context | Docker Documentation <https://docs.docker.com/engine/context/working-with-contexts/>`_
       | `使用ECS CLI 部署docker-compose到AWS ECS | by Leo Chang | Medium <https://medium.com/@cchangleo/%E4%BD%BF%E7%94%A8ecs-cli-%E9%83%A8%E7%BD%B2docker-compose%E5%88%B0aws-ecs-3a3a13b2494e>`_
       | `Docker to ECS tutorial: Manual deployment | by Luka Klaric | Sep, 2021 | FAUN Publication <https://faun.pub/deploying-your-first-docker-container-on-aws-ecs-ed19a3599b6c>`_
       | `Setting up with Amazon ECS - Amazon Elastic Container Service <https://docs.aws.amazon.com/AmazonECS/latest/developerguide/get-set-up-for-amazon-ecs.html>`_
       | `Getting started with Amazon ECS using AWS Copilot - Amazon Elastic Container Service <https://docs.aws.amazon.com/AmazonECS/latest/developerguide/getting-started-aws-copilot-cli.html>`_

.. [3] | `docker context create ecs fail - Google search <https://www.google.com/search?q=docker+context+create+ecs+fail>`_
       | `docker context create ecs fail - DuckDuckGo search <https://duckduckgo.com/?q=docker+context+create+ecs+fail>`_
       | `docker context create ecs fail - Ecosia search <https://www.ecosia.org/search?q=docker+context+create+ecs+fail>`_
       | `docker context create ecs fail - Qwant search <https://www.qwant.com/?q=docker+context+create+ecs+fail>`_
       | `docker context create ecs fail - Bing search <https://www.bing.com/search?q=docker+context+create+ecs+fail>`_
       | `docker context create ecs fail - Yahoo search <https://search.yahoo.com/search?p=docker+context+create+ecs+fail>`_
       | `docker context create ecs fail - Baidu search <https://www.baidu.com/s?wd=docker+context+create+ecs+fail>`_
       | `docker context create ecs fail - Yandex search <https://www.yandex.com/search/?text=docker+context+create+ecs+fail>`_
       |
       | `amazon web services - docker context create ecs myecs - requires exactly one argument - Stack Overflow <https://stackoverflow.com/questions/67236401/docker-context-create-ecs-myecs-requires-exactly-one-argument>`_
       | `Install the Docker Compose CLI on Linux - Deploying Docker containers on ECS <https://docs.docker.com/cloud/ecs-integration/#install-the-docker-compose-cli-on-linux>`_

.. [4] | Fail: aws ecr get-login-password | docker login --username AWS --password-stdin aws_account_id.dkr.ecr.region.amazonaws.com
       | `How to fix docker: Got permission denied issue - Stack Overflow <https://stackoverflow.com/questions/48957195/how-to-fix-docker-got-permission-denied-issue>`_

.. [5] | `docker communicate between containers ecs - Google search <https://www.google.com/search?q=docker+communicate+between+containers+ecs>`_
       | `docker communicate between containers ecs - DuckDuckGo search <https://duckduckgo.com/?q=docker+communicate+between+containers+ecs>`_
       | `docker communicate between containers ecs - Ecosia search <https://www.ecosia.org/search?q=docker+communicate+between+containers+ecs>`_
       | `docker communicate between containers ecs - Qwant search <https://www.qwant.com/?q=docker+communicate+between+containers+ecs>`_
       | `docker communicate between containers ecs - Bing search <https://www.bing.com/search?q=docker+communicate+between+containers+ecs>`_
       | `docker communicate between containers ecs - Yahoo search <https://search.yahoo.com/search?p=docker+communicate+between+containers+ecs>`_
       | `docker communicate between containers ecs - Baidu search <https://www.baidu.com/s?wd=docker+communicate+between+containers+ecs>`_
       | `docker communicate between containers ecs - Yandex search <https://www.yandex.com/search/?text=docker+communicate+between+containers+ecs>`_
       | `Question - how to communicate between docker compose services in fargate · Issue #1005 · aws/amazon-ecs-cli · GitHub <https://github.com/aws/amazon-ecs-cli/issues/1005>`_

.. [6] | `go kit service discovery - Google search <https://www.google.com/search?q=go+kit+service+discovery>`_
       | `go kit service discovery - DuckDuckGo search <https://duckduckgo.com/?q=go+kit+service+discovery>`_
       | `go kit service discovery - Ecosia search <https://www.ecosia.org/search?q=go+kit+service+discovery>`_
       | `go kit service discovery - Qwant search <https://www.qwant.com/?q=go+kit+service+discovery>`_
       | `go kit service discovery - Bing search <https://www.bing.com/search?q=go+kit+service+discovery>`_
       | `go kit service discovery - Yahoo search <https://search.yahoo.com/search?p=go+kit+service+discovery>`_
       | `go kit service discovery - Baidu search <https://www.baidu.com/s?wd=go+kit+service+discovery>`_
       | `go kit service discovery - Yandex search <https://www.yandex.com/search/?text=go+kit+service+discovery>`_


.. _Docker Compose: https://docs.docker.com/compose/
