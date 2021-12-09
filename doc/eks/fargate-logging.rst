Short URL Service Logging / Fargate Logging
===========================================

The short URL service uses `Go kit`_ for microservice framwork. Logging is
standard feature for any online service. How do we view logs from the service in
Amazon EKS? Fargate is used as nodes in the cluster of short URL service, Amazon
EKS on Fargate offers a built-in log router based on Fluent Bit.

- `Fargate logging - Amazon EKS <https://docs.aws.amazon.com/eks/latest/userguide/fargate-logging.html>`_

It is quite complicated to set up the logging and pipe the logs from Fargate to
Amazon CloudWatch. The following is the note I wrote down during setup:

- If you create the cluster via ``eksctl``, the ``eksctl`` already creates a
  role named ``eksctl-CLUSTER_NAME-cluster-FargatePodExecutionRole-SOME_ID``
  with ``AmazonEKSFargatePodExecutionRolePolicy``. We can use this existing
  Fargate pod execution role instead of creating a new one.
- Remember to set correct region in ``ConfigMap``, the same region as your
  Kubernetes cluster.

Follow the instructions and if everything is set correctly, you will be able to
view logs in the console of Amazon CloudWatch.

.. _Go kit: https://github.com/go-kit/kit
