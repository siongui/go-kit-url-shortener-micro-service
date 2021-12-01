Translate Docker Compose File to Kubernetes Resources
=====================================================

During development on local machines, `Docker Compose`_ provides easy to setup,
intuitive, and multi-container environment for running and testing our
application. Docker Compose simplifies the development process. After
development finished, to deploy the container application on Kubernetes, we need
to translate the Docker Compose YAML file which configure application's services
to Kubernetes resources. The tool Kompose_ can help us to convert. [1]_ [2]_
[3]_

Kompose_ is a conversion tool which translate the services and container
environment defined in ``docker-compose.yml`` to Kubernetes resources files. In
normal cases, each container provisioning one service in the Docker Compose
environment will be translate to two resources YAML files, one for Kubernetes
*deployment* and the other for Kubernetes *service*. It is very easy to use
Kompose_, just run the following command in the same directory as
``docker-compose.yml``:

.. code-block:: bash

  kompose convert

The tool will create all the Kubernetes resources files for us. Let's look at an
example. The following is the Docker Compose file which describes the services
and environment for short URL service:

.. code-block:: yaml

  services:
    docker-url-shortener:
      build:
        context: .
      image: url-shortener:latest
      ports:
        - 8080:8080
      environment:
        - SQLITE_DSN=${SQLITE_DSN:-file::memory:?cache=shared}

After running the conversion, two files are created.

- docker-url-shortener-deployment.yaml
- docker-url-shortener-service.yaml

Kompose cannot 100% convert to what we really need, so we have to take a look at
the two files and modify as we need. Take a look at *deployment* file first:

**docker-url-shortener-deployment.yaml**

.. code-block:: yaml

  apiVersion: apps/v1
  kind: Deployment
  metadata:
    annotations:
      kompose.cmd: kompose convert
      kompose.version: 1.26.0 (40646f47)
    creationTimestamp: null
    labels:
      io.kompose.service: docker-url-shortener
    name: docker-url-shortener
  spec:
    replicas: 1
    selector:
      matchLabels:
        io.kompose.service: docker-url-shortener
    strategy: {}
    template:
      metadata:
        annotations:
          kompose.cmd: kompose convert
          kompose.version: 1.26.0 (40646f47)
        creationTimestamp: null
        labels:
          io.kompose.service: docker-url-shortener
      spec:
        containers:
          - env:
              - name: SQLITE_DSN
                value: file::memory:?cache=shared
            image: url-shortener:latest
            name: docker-url-shortener
            ports:
              - containerPort: 8080
            resources: {}
        restartPolicy: Always
  status: {}

The problem here is the following line:

.. code-block:: yaml

  image: url-shortener:latest

We can not serve Docker images like this on Amazon EKS. We need to supply the
image through container registry. Previous tutorial shows how to push the image
to GitHub Packages, so we specify the URI of the image as follows:

.. code-block:: yaml

  image: ghcr.io/siongui/go-kit-url-shortener-micro-service:main

Then the image can be deployed on Amazon EKS successfully. Next take a look at
*service* resource file:

**docker-url-shortener-service.yaml**

.. code-block:: yaml

  apiVersion: v1
  kind: Service
  metadata:
    annotations:
      kompose.cmd: kompose convert
      kompose.version: 1.26.0 (40646f47)
    creationTimestamp: null
    labels:
      io.kompose.service: docker-url-shortener
    name: docker-url-shortener
  spec:
    ports:
      - name: "8080"
        port: 8080
        targetPort: 8080
    selector:
      io.kompose.service: docker-url-shortener
  status:
    loadBalancer: {}

The above service can be deployed sucessfully on Amazon EKS, but cannot receive
connections from outside internet. To allow connections from internet, we need

- Install ``AWS Load Balancer Controller`` [4]_
- Specify ``spec.type`` as ``LoadBalancer``
- Add following three lines in ``metadata.annotations``

  .. code-block:: yaml

    service.beta.kubernetes.io/aws-load-balancer-type: external
    service.beta.kubernetes.io/aws-load-balancer-nlb-target-type: ip
    service.beta.kubernetes.io/aws-load-balancer-scheme: internet-facing

We will take more about Load Balancer later in the following tutorials. Here we
give you a whole picture of how to quick start instead of diving into the
details. To see complete working Kubernetes resources for short URL service,
see `resources/eks/`_ folder on GitHub.

.. [1] `Translate a Docker Compose File to Kubernetes Resources | Kubernetes <https://kubernetes.io/docs/tasks/configure-pod-container/translate-compose-kubernetes/>`_
.. [2] `Docker Compose to Kubernetes: Step-by-Step Migration <https://loft.sh/blog/docker-compose-to-kubernetes-step-by-step-migration/>`_
.. [3] `How To Migrate a Docker Compose Workflow to Kubernetes | DigitalOcean <https://www.digitalocean.com/community/tutorials/how-to-migrate-a-docker-compose-workflow-to-kubernetes>`_
.. [4] `Network load balancing on Amazon EKS - Amazon EKS <https://docs.aws.amazon.com/eks/latest/userguide/network-load-balancing.html>`_

.. _Docker Compose: https://docs.docker.com/compose/
.. _Kompose: https://kompose.io/
.. _resources/eks/: https://github.com/siongui/go-kit-url-shortener-micro-service/tree/main/resources/eks
