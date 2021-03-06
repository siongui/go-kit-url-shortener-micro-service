Summary
+++++++


Summary about how to deploy Go microservice in this repo to `Amazon EKS`_ and
`Amazon RDS PostgresSQL`_.

- `Push Docker Image to GitHub Pachages <push-docker-image-to-github-packages.rst>`_
- `Translate Docker Compose File to Kubernetes Resources <docker-compose-to-kubernetes.rst>`_
- `Quick Start with Amazon EKS <quick-start-with-amazon-ecs.rst>`_
- `Deploy Short URL Service Application <deploy-short-url-service-app.rst>`_
- `Expose the Service on Amazon EKS <expose-the-service-on-amazon-eks.rst>`_
- `Access Amazon RDS PostgreSQL in Amazon EKS Cluster <access-amazon-rds-in-amazon-eks.rst>`_
- `Short URL Service Logging / Fargate Logging <fargate-logging.rst>`_


----


Deploy Docker Compose to Amazon EKS
+++++++++++++++++++++++++++++++++++

- `Translate a Docker Compose File to Kubernetes Resources | Kubernetes <https://kubernetes.io/docs/tasks/configure-pod-container/translate-compose-kubernetes/>`_
- `Getting started with Amazon EKS - Amazon EKS <https://docs.aws.amazon.com/eks/latest/userguide/getting-started.html>`_
- `Deploy a sample application - Amazon EKS <https://docs.aws.amazon.com/eks/latest/userguide/sample-deployment.html>`_
- | `Expose Kubernetes services running on Amazon EKS clusters <https://aws.amazon.com/premiumsupport/knowledge-center/eks-kubernetes-services-cluster/>`_
  | `Exposing the Service :: Amazon EKS Workshop <https://www.eksworkshop.com/beginner/130_exposing-service/exposing/>`_
  | `AWS Load Balancer Controller - Amazon EKS <https://docs.aws.amazon.com/eks/latest/userguide/aws-load-balancer-controller.html>`_
  | `How it works - AWS Load Balancer Controller <https://kubernetes-sigs.github.io/aws-load-balancer-controller/v2.3/how-it-works/>`_
  | `Network load balancing on Amazon EKS`_
  | In service yaml, ``spec.type`` set to ``LoadBalancer``
  | ``metadata.annotations`` set:
  | service.beta.kubernetes.io/aws-load-balancer-type: external
  | service.beta.kubernetes.io/aws-load-balancer-nlb-target-type: ip
  | service.beta.kubernetes.io/aws-load-balancer-scheme: internet-facing
- **Important Note**: In `Network load balancing on Amazon EKS`_: ``it says If you're deploying to Fargate nodes, remove the service.beta.kubernetes.io/aws-load-balancer-scheme: internet-facing line.``.
  If you look at the overview of the cluster in the EKS console, the node name
  starts with *fargete*. If you do remove ``internet-facing``, the service can
  not be accessed from outside internet. The documentation is **misleading**.


Connect to Amazon RDS
+++++++++++++++++++++

- Follow instructions in [2]_, choose ``Standard create``.
- Create the PostgreSQL DB instance, choose the same VPC as EKS cluster.
- Remember to specify *Initial database name* for PostgresSQL instance in the
  **Additional configuration** of ``Standard create``. If you do not specify a
  database name, Amazon RDS does not create a database.
- Follow instructions in [3]_ to access the DB instance in the same VPC.
- If the pod cannot connect to the DB instance, check the security group of the
  DB instance. For example, if the DB instance belongs to ``default`` security
  group of the VPC, try to allow all inbound traffic from **Source** of other
  security group in the same VPC.


Useful Command
++++++++++++++

View all status of default ``kube-system``

.. code-block:: bash

  kubectl -n kube-system get all

Create fargate profile ``url-shorten-app``

.. code-block:: bash

  eksctl create fargateprofile \
      --cluster my-cluster \
      --region region-code \
      --name url-shorten-app \
      --namespace url-shorten-app

Create a namespace for the application ``url-shorten-app``

.. code-block:: bash

  kubectl create namespace url-shorten-app

Apply the deployment manifest to the cluster

.. code-block:: bash

  kubectl apply -f resources/eks/docker-url-shortener-deployment.yaml

Apply the service manifest to the cluster

.. code-block:: bash

  kubectl apply -f resources/eks/docker-url-shortener-service.yaml

View all status of the application ``url-shorten-app``

.. code-block:: bash

  kubectl -n url-shorten-app get all

Delete the namespace

.. code-block:: bash

  kubectl delete namespace url-shorten-app


Fargate logging
+++++++++++++++

- `Fargate logging - Amazon EKS <https://docs.aws.amazon.com/eks/latest/userguide/fargate-logging.html>`_
- If you create the cluster via ``eksctl``, the ``eksctl`` already creates a
  role named ``eksctl-CLUSTER_NAME-cluster-FargatePodExecutionRole-SOME_ID``
  with ``AmazonEKSFargatePodExecutionRolePolicy``. We can use this existing
  Fargate pod execution role instead of creating a new one.
- Remember to set correct region in ``ConfigMap``.


References
++++++++++

.. [1] | `Building a bare-metal Kubernetes cluster on Raspberry Pi | Hacker News <https://news.ycombinator.com/item?id=29306616>`_
       | `Building a bare-metal Kubernetes cluster on Raspberry Pi <https://anthonynsimon.com/blog/kubernetes-cluster-raspberry-pi/>`_

.. [2] `Creating a PostgreSQL DB instance and connecting to a database on a PostgreSQL DB instance - Amazon Relational Database Service <https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/CHAP_GettingStarted.CreatingConnecting.PostgreSQL.html>`_

.. [3] | `Scenarios for accessing a DB instance in a VPC - Amazon Relational Database Service <https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/USER_VPC.Scenarios.html>`_
       | `Connecting to a DB instance running the PostgreSQL database engine <https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/USER_ConnectToPostgreSQLInstance.html>`_

.. [4] | `Accessing Amazon RDS From AWS EKS - DEV Community <https://dev.to/bensooraj/accessing-amazon-rds-from-aws-eks-2pc3>`_
       | `Accessing Amazon RDS From AWS EKS - Google search <https://www.google.com/search?q=Accessing+Amazon+RDS+From+AWS+EKS>`_
       | `Accessing Amazon RDS From AWS EKS - DuckDuckGo search <https://duckduckgo.com/?q=Accessing+Amazon+RDS+From+AWS+EKS>`_
       | `Accessing Amazon RDS From AWS EKS - Ecosia search <https://www.ecosia.org/search?q=Accessing+Amazon+RDS+From+AWS+EKS>`_
       | `Accessing Amazon RDS From AWS EKS - Qwant search <https://www.qwant.com/?q=Accessing+Amazon+RDS+From+AWS+EKS>`_
       | `Accessing Amazon RDS From AWS EKS - Bing search <https://www.bing.com/search?q=Accessing+Amazon+RDS+From+AWS+EKS>`_
       | `Accessing Amazon RDS From AWS EKS - Yahoo search <https://search.yahoo.com/search?p=Accessing+Amazon+RDS+From+AWS+EKS>`_
       | `Accessing Amazon RDS From AWS EKS - Baidu search <https://www.baidu.com/s?wd=Accessing+Amazon+RDS+From+AWS+EKS>`_
       | `Accessing Amazon RDS From AWS EKS - Yandex search <https://www.yandex.com/search/?text=Accessing+Amazon+RDS+From+AWS+EKS>`_


.. _Docker Compose: https://docs.docker.com/compose/
.. _Network load balancing on Amazon EKS: https://docs.aws.amazon.com/eks/latest/userguide/network-load-balancing.html
.. _Amazon EKS: https://aws.amazon.com/eks/
.. _Amazon RDS PostgresSQL: https://aws.amazon.com/rds/postgresql/
