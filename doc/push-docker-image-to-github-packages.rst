This tutorial assumes that readers already have basic knowledge about building
Go Docker images [1]_ and GitHub workflow [2]_.

If we deploy our application to cloud platform, such as Amazon EKS, we cannot
build the Docker image on the fly and then deploy the image directly to the
cloud service. That's why we need *container registry*, which we pre-build our
Docker images and store the images to the registry. There are many choices of
registry to push our Docker images. For exmaple, Docker Hub or Amazon Elastic
Container Registry are popular choices to store, share and deploy the images.

*GitHub Packages* is a service by GitHub to provide container registry. Since
most developers use GitHub for daily coding work, we will choose GitHub Packages
to host Docker images. This tutorial is a quick start to use `GitHub Actions`_
to automatically build and push Docker images to GitHub Packages every time the
source code is pushed to GitHub repository.

The following is a typical CI workflow under ``.github/workflows/`` folder for
testing Go source code.

.. code-block:: yaml

  name: ci

  on:
    push:
      branches:
        - main

  jobs:

    test:
      runs-on: ubuntu-latest
      steps:
      - uses: actions/checkout@v2.3.1
        with:
          persist-credentials: false
      - uses: actions/setup-go@v2
        with:
          go-version: '1.17.2'
      - name: Test
        run: |
          make test

The above CI workflow check out the source code, setup Go toolchain, and run
test on the Go code. Now we want to build and push Docker images to container
registry after successful test. We can add the following actions to the end of
our CI yaml file (assume the Dockerfile is already located at root of the GitHub
repository).

.. code-block:: yaml

      - name: Extract metadata (tags, labels) for Docker
        id: meta
        uses: docker/metadata-action@v3
        with:
          images: ${{ env.REGISTRY }}/${{ env.IMAGE_NAME }}
      - name: Log in to the Container registry
        uses: docker/login-action@v1
        with:
          registry: ${{ env.REGISTRY }}
          username: ${{ github.actor }}
          password: ${{ secrets.GITHUB_TOKEN }}
      - name: Build and push Docker image
        uses: docker/build-push-action@v2
        with:
          context: .
          push: true
          tags: ${{ steps.meta.outputs.tags }}
          labels: ${{ steps.meta.outputs.labels }}

The actions here do:

- Extract metadata from your repo, such as the name of the Docker image and tag.
- Login to the registry on your behalf. In this case, login to GitHub Actions.
- Finally build and push the Docker image to the registry, GitHub Pachages in
  this case.

If the repo name is ``siongui/go-kit-url-shortener-micro-service`` and the
default branch is ``main``, the final Docker image will be at
``ghcr.io/siongui/go-kit-url-shortener-micro-service:main``. You can use this
URI to deploy the image on cloud services.

For complete CI example file, please refer to ci.yml_. For more in-depth guide,
please see refer to [3]_, [4]_ and [5]_.

.. [1] `Build your Go image | Docker Documentation <https://docs.docker.com/language/golang/build-images/>`_
.. [2] `Quickstart for GitHub Actions - GitHub Docs <https://docs.github.com/en/actions/quickstart>`_
.. [3] `Publishing Docker images - GitHub Docs <https://docs.github.com/en/actions/publishing-packages/publishing-docker-images>`_
.. [4] `docker/metadata-action - GitHub <https://github.com/docker/metadata-action>`_
.. [5] `docker/build-push-action - GitHub <https://github.com/docker/build-push-action>`_

.. _GitHub Actions: https://github.com/features/actions
.. _ci.yml: https://github.com/siongui/go-kit-url-shortener-micro-service/blob/main/.github/workflows/ci.yml
