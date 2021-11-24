Deploy Docker Compose to Amazon EKS
+++++++++++++++++++++++++++++++++++

- `Translate a Docker Compose File to Kubernetes Resources | Kubernetes <https://kubernetes.io/docs/tasks/configure-pod-container/translate-compose-kubernetes/>`_
- `Getting started with Amazon EKS - Amazon EKS <https://docs.aws.amazon.com/eks/latest/userguide/getting-started.html>`_
- `Deploy a sample application - Amazon EKS <https://docs.aws.amazon.com/eks/latest/userguide/sample-deployment.html>`_
- | `Expose Kubernetes services running on Amazon EKS clusters <https://aws.amazon.com/tw/premiumsupport/knowledge-center/eks-kubernetes-services-cluster/>`_
  | `Exposing the Service :: Amazon EKS Workshop <https://www.eksworkshop.com/beginner/130_exposing-service/exposing/>`_
  | `AWS Load Balancer Controller - Amazon EKS <https://docs.aws.amazon.com/eks/latest/userguide/aws-load-balancer-controller.html>`_
  | `How it works - AWS Load Balancer Controller <https://kubernetes-sigs.github.io/aws-load-balancer-controller/v2.3/how-it-works/>`_
  | `Network load balancing on Amazon EKS <https://docs.aws.amazon.com/eks/latest/userguide/network-load-balancing.html>`_
  | In service yaml, ``spec.type`` set to ``LoadBalancer``
  | ``metadata.annotations`` set:
  | service.beta.kubernetes.io/aws-load-balancer-type: external
  | service.beta.kubernetes.io/aws-load-balancer-nlb-target-type: ip


References
++++++++++

.. [1] | `Building a bare-metal Kubernetes cluster on Raspberry Pi | Hacker News <https://news.ycombinator.com/item?id=29306616>`_
       | `Building a bare-metal Kubernetes cluster on Raspberry Pi <https://anthonynsimon.com/blog/kubernetes-cluster-raspberry-pi/>`_

.. _Docker Compose: https://docs.docker.com/compose/
