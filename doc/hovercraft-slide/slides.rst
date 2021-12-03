:title: Go Microservice with Go Kit
:data-transition-duration: 1000
:css: font.css


----

Go Microservice with Go Kit
===========================

Dec 3, 2021


----

Overview
========

- Go microservice framework
- Why Choose Go Kit
- Go Kit Features
- How to Learn Go Kit
- Quick Demo

----

Go Microservice Framework
=========================

- `Go kit <https://github.com/go-kit/kit>`_ (21.8K GitHub star)

  * Loosely coupled
  * Easy to add customized features
  * Need good knowledge of Go language features

- `Kratos <https://go-kratos.dev/>`_ (15.6K) & `Go Micro <https://github.com/micro/micro>`_ (10.7K)

  * Opinionated
  * Self-written tool to create project boilerplate
  * Quick prototype, less flexibility


----

Why Choose Go Kit
=================

I choose Go kit because

- Most stars on GitHub. Easy for googling and find answers
- Standard Go toolchain to manage the project. No tools from framework
- Easy to add features, such as HTTP 303 See Other
- Easy to integrate web framework `gin <https://github.com/gin-gonic/gin>`_


----

Go Kit Features
===============

- JSON over HTTP (RESTful API), gRPC, etc.
- Separation of concerns

  * Logging
  * Instrumentation

- Service discovery
- Load balancing/ Circuit Breaker/ Rate Limit


----

Go Kit Advance Topics
=====================

- Threading a context
- Distributed tracing
- Creating a client package/SDK


----

How to Learn Go Kit
===================

- MUST read `Go kit stringsvc tutorial <https://gokit.io/examples/stringsvc.html>`_
- The tutorial covers all important features of Go Kit
- The sample code is short and easy to pick up
- You must run it yourself!!!


----

Quick Demo
==========

- Short URL service using Go kit
- Use only following Go kit features

  * JSON over HTTP
  * Logging
  * Instrumentation

- Demo
- Q&A
