Expose the Service on Amazon EKS
================================

When you create a cluster on Amazon EKS and deploy a service in the cluster, the
service is not accessible from outside internet by default. To allow network
traffic into the service in the cluster is called exposing the service. The
following guides discuss about how to expose the service:

- `Expose Kubernetes services running on Amazon EKS clusters <https://aws.amazon.com/premiumsupport/knowledge-center/eks-kubernetes-services-cluster/>`_
- `Exposing the Service :: Amazon EKS Workshop <https://www.eksworkshop.com/beginner/130_exposing-service/exposing/>`_

There are three different *Kubernetes ServiceType* to expose the service. It
seems that the only possible way to make the service accessible directly from
outside internet is to use a load balancer, i.e., **LoadBalancer** Kubernetes
ServiceType.

To use the load balancer, we need to install *AWS Load Balancer Controller* in
the Kubernetes cluster. The following guides show how to install and use the
*AWS Load Balancer Controller*:

- `AWS Load Balancer Controller - Amazon EKS <https://docs.aws.amazon.com/eks/latest/userguide/aws-load-balancer-controller.html>`_

The guide is full of strange terms for beginners and even now I do not know what
I was doing while the installation. The trick I learned is that you can use the
following command to see if the pods of the controller is running:

.. code-block:: bash

  kubectl -n kube-system get all

If no pod of *AWS Load Balancer Controller* is running, you need to check if
you follow the instructions correctly and try again.

The following link gives a general idea of how AWS Load Balancer Controller
works:

- `How it works - AWS Load Balancer Controller <https://kubernetes-sigs.github.io/aws-load-balancer-controller/v2.3/how-it-works/>`_

There are two kinds of load balancing provisioned by the controller -
*application load balancing* and *network load balancing*:

- `Application load balancing on Amazon EKS - Amazon EKS <https://docs.aws.amazon.com/eks/latest/userguide/alb-ingress.html>`_
- `Network load balancing on Amazon EKS - Amazon EKS <https://docs.aws.amazon.com/eks/latest/userguide/network-load-balancing.html>`_

Here the short URL service choose network load balancing to expose itself
externally. Add the following in the service manifest of short URL service:

- Specify ``spec.type`` as ``LoadBalancer``
- Add following three lines in ``metadata.annotations``

  .. code-block:: yaml

    service.beta.kubernetes.io/aws-load-balancer-type: external
    service.beta.kubernetes.io/aws-load-balancer-nlb-target-type: ip
    service.beta.kubernetes.io/aws-load-balancer-scheme: internet-facing

For example of complete working service manifest, see
`docker-url-shortener-service.yaml <../../resources/eks/docker-url-shortener-service.yaml>`_

Note that the documentation of network load balancing says:

  If you're deploying to Fargate nodes, remove the
  ``service.beta.kubernetes.io/aws-load-balancer-scheme: internet-facing`` line.

This is not true. If the internet-facing line is removed, the service cannot be
accessed from the internet. The documentation is misleading.

In the beginning I think it is easy to expose the service, but it turned out
that this is the time-consuming, hard to understand, and difficult part of the
deployment process.
