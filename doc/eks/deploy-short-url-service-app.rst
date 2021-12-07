Deploy Short URL Service Application
====================================

In previous tutorials, we learn about how to push Docker image to GitHub
Packages, translate Docker Compose files to Kubernetes resources, and create a
cluster on Amazon EKS. This tutorial will show how to deploy the short URL
service in this repository to the Kubernetes cluster on Amazon EKS, based on the
following user guide on Amazon EKS:

- `Deploy a sample application - Amazon EKS <https://docs.aws.amazon.com/eks/latest/userguide/sample-deployment.html>`_

We will create a cluster from scratch.

.. code-block:: bash

  eksctl create cluster \
  --name url-shorten \
  --region ap-northeast-1 \
  --fargate

A cluster named ``url-shorten`` is created in the region
``ap-northeast-1 (Tokyo)``. Next create a `AWS Fargate profile`_ named
``url-shorten-app`` for short URL service:

.. code-block:: bash

  eksctl create fargateprofile \
      --cluster url-shorten \
      --region ap-northeast-1 \
      --name url-shorten-app \
      --namespace url-shorten-app

Then use the same name to create Kubernetes namespace ``url-shorten-app``

.. code-block:: bash

  kubectl create namespace url-shorten-app

Next, follow the following guide to install *AWS Load Balancer Controller*,
which will help expose the servie to outside internet.

- `AWS Load Balancer Controller - Amazon EKS <https://docs.aws.amazon.com/eks/latest/userguide/aws-load-balancer-controller.html>`_

It is very complicated to install *AWS Load Balancer Controller*.
Run the following command to see if the pods of AWS LBC is running correctly.

.. code-block:: bash

  kubectl -n kube-system get all

If all goes well, you will see something like this:

.. code-block:: txt

  NAME                                                READY   STATUS    RESTARTS   AGE
  pod/aws-load-balancer-controller-6cf4589f69-6pbk6   1/1     Running   0          12d
  pod/aws-load-balancer-controller-6cf4589f69-pt27g   1/1     Running   0          12d
  pod/coredns-9f6f89c76-5clhw                         1/1     Running   0          13d
  pod/coredns-9f6f89c76-l7xp8                         1/1     Running   0          13d

  NAME                                        TYPE        CLUSTER-IP       EXTERNAL-IP   PORT(S)         AGE
  service/aws-load-balancer-webhook-service   ClusterIP   10.100.156.160   <none>        443/TCP         12d
  service/kube-dns                            ClusterIP   10.100.0.10      <none>        53/UDP,53/TCP   13d

  NAME                        DESIRED   CURRENT   READY   UP-TO-DATE   AVAILABLE   NODE SELECTOR   AGE
  daemonset.apps/aws-node     0         0         0       0            0           <none>          13d
  daemonset.apps/kube-proxy   0         0         0       0            0           <none>          13d

  NAME                                           READY   UP-TO-DATE   AVAILABLE   AGE
  deployment.apps/aws-load-balancer-controller   2/2     2            2           12d
  deployment.apps/coredns                        2/2     2            2           13d

  NAME                                                      DESIRED   CURRENT   READY   AGE
  replicaset.apps/aws-load-balancer-controller-6cf4589f69   2         2         2       12d
  replicaset.apps/coredns-76f4967988                        0         0         0       13d
  replicaset.apps/coredns-9f6f89c76                         2         2         2       13d


Apply the deployment manifest of short URL service to the cluster:

.. code-block:: bash

  kubectl apply -f resources/eks/docker-url-shortener-deployment.yaml

Apply the service manifest of short URL service to the cluster:

.. code-block:: bash

  kubectl apply -f resources/eks/docker-url-shortener-service.yaml

View all status of the application ``url-shorten-app`` to see if the short URL
service is running correctly:

.. code-block:: bash

  kubectl -n url-shorten-app get all

If everything goes fine, you will something like this:

.. code-block:: txt

  NAME                                   READY   STATUS    RESTARTS   AGE
  pod/url-shorten-app-65d6795cd5-bctvb   1/1     Running   19         8d

  NAME                      TYPE           CLUSTER-IP      EXTERNAL-IP                                                                          PORT(S)          AGE
  service/url-shorten-app   LoadBalancer   10.100.69.173   k8s-urlshort-urlshort-0e19b7a815-51a8337aa2444e28.elb.ap-northeast-1.amazonaws.com   8080:31478/TCP   12d

  NAME                              READY   UP-TO-DATE   AVAILABLE   AGE
  deployment.apps/url-shorten-app   1/1     1            1           8d

  NAME                                         DESIRED   CURRENT   READY   AGE
  replicaset.apps/url-shorten-app-65d6795cd5   1         1         1       8d

Congratulations! From the *EXTERNAL-IP* field, the service is ready to receive
requests at ``http://k8s-urlshort-urlshort-0e19b7a815-51a8337aa2444e28.elb.ap-northeast-1.amazonaws.com:8080/``


.. _AWS Fargate: https://docs.aws.amazon.com/eks/latest/userguide/fargate.html
.. _AWS Fargate profile: https://docs.aws.amazon.com/eks/latest/userguide/fargate-profile.html
