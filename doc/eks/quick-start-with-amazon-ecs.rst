Quick Start with Amazon EKS
===========================

Although the name is *quick start*, I think for beginners to set up the account
and all the command line tools is not *quick* at all. There are a lot of special
terminology, services, concepts that are specific to Amazon Web Services. It is
not easy to understand the reason behind each step in the beginning. So be
patient and try to ask help from someone with experiences.

The first step is to install the AWS command line interface:

- `Getting started with the AWS CLI <https://docs.aws.amazon.com/cli/latest/userguide/cli-chap-getting-started.html>`_

To be able to use Amazon EKS, in the
*prerequisites to use the AWS CLI version 2*,

- Step 1: Sign up to AWS
- Step 2: Create an IAM user account

is necessary. Then go to the *Install/Update* page to install AWS CLI on your
platform (Linux/macOS/Windows). After installation, go to *Quick setup* page,
run

.. code-block:: bash

  aws configure

The AWS CLI will prompts you for

- Access key ID
- Secret access key
- AWS Region
- Output format

Follow the instructions in the page to finish the setup of AWS CLI.

Next we will use a command line tool called *eksctl* to create Kubernetes
clusters on Amazon EKS. The following guides show you how to do it.

- `Getting started with Amazon EKS <https://docs.aws.amazon.com/eks/latest/userguide/getting-started.html>`_
- `Getting started with Amazon EKS – eksctl <https://docs.aws.amazon.com/eks/latest/userguide/getting-started-eksctl.html>`_

In summary, what you need to do in above guides:

- Install *kubectl*
- Install *eksctl*
- Create Amazon EKS cluster and nodes
- View resources of the cluster
- Delete cluster and nodes

After finish the guide, you should have a basic understanding about what a
Kubernetes cluster looks like.
