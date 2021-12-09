Access Amazon RDS PostgreSQL in Amazon EKS Cluster
==================================================

It is very common for web applications to use database for persistent storage.
The short URL service also uses database for storing the original and shortened
URLs. Through the setting of environment variables, the short URL service will
choose which database (PostgreSQL or SQLite). To use PostgreSQL, the following
environment variables must be set:

- POSTGRES_USER
- POSTGRES_PASSWORD
- POSTGRES_HOST
- POSTGRES_PORT
- POSTGRES_DB

In Kubernetes service manifest, the setting looks like:

.. code-block:: yaml

    spec:
      containers:
        - image: ghcr.io/siongui/go-kit-url-shortener-micro-service:main
          env:
            - name: POSTGRES_USER
              value: "postgres"
            - name: POSTGRES_PASSWORD
              value: "YOUR_PASSWORD"
            - name: POSTGRES_HOST
              value: "YOUR_SUBDOMAIN.rds.amazonaws.com"
            - name: POSTGRES_PORT
              value: "5432"
            - name: POSTGRES_DB
              value: "YOUR_DATABASE_NAME"
          name: url-shorten-app
          ports:
            - name: tcp
              containerPort: 8080
          resources: {}

If POSTGRES_* are set correctly, the short URL service will connect to the
specified PostgreSQL database for persistent storage.

In previous tutorial, the short URL service is deployed in the Kubernetes
cluster. In this tutorial, we will attach a DB instance of Amazon ADS PostgreSQL
to the cluster and tell the short URL service to use the DB instance through
environment variables.

First we will create a DB instance in the same region as the Kubernetes cluster.
For example, if you create the cluster in ``ap-northeast-1 (Tokyo)``, you also
have to create the DB instance in the same region ``ap-northeast-1 (Tokyo)``.

- `Creating a PostgreSQL DB instance and connecting to a database on a PostgreSQL DB instance - Amazon Relational Database Service <https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/CHAP_GettingStarted.CreatingConnecting.PostgreSQL.html>`_
- Refer to the above guide. Choose **Standard create** instead of
  **Easy create**
- Create the PostgreSQL DB instance
- Choose the same VPC as EKS cluster
- Remember to specify *Initial database name* for PostgresSQL instance in the
  **Additional configuration** of ``Standard create``. If you do not specify a
  database name, Amazon RDS does not create a database.

Remember to write down the user name, password, and database name. You need to
set these in the environment variable. After creation of the DB instance, you
can find host and port of the DB instance in the AWS management console. Refer
to the following guide for more information:

- `Connecting to a DB instance running the PostgreSQL database engine <https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/USER_ConnectToPostgreSQLInstance.html>`_
- `Scenarios for accessing a DB instance in a VPC - Amazon Relational Database Service <https://docs.aws.amazon.com/AmazonRDS/latest/UserGuide/USER_VPC.Scenarios.html>`_

If all POSTGRES_* env are set correctly in the service manifest, restart the pod
of short URL service, go to the Amazon RDS console, now you should be able to
see one connection from the service.

Troubleshooting: If the pod cannot connect to the DB instance, check the
security group of the DB instance. For example, if the DB instance belongs to
``default`` security group of the VPC, try to allow all inbound traffic from
**Source** of other security group in the same VPC.
