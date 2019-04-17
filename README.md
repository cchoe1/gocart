# Introduction
The goal of this is to act as a small microservice that foregoes the need to connect a shopping cart to an intricate 
e-commerce website that may require heavy refactoring and code rewriting.  This *ideally* will be able to sit on top
of any e-commerce website and provide shopping cart capabilities with minimal refactoring.

The data transfer will happen via a common interface--HTTP requests.  SPAs and traditional models alike will be able
to query the service for information on a given cart.  The internals will be powered by Golang and persistence can be
achieved using any typical database software (by default, Mysql support will be included).  

Configuration will be saved as yaml or json (not yet determined). 

# Roadmap

1. Implement internals
2. Expose functionality via REST API
3. Dockerize
