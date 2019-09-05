# go-template-repository

## Description

Este es un repositorio plantilla para proyectos go

## Organization

* *main.go* va en el root del proyecto
* *data* es para definir paquetes que acedan a los datos
* *api* es para definir paquetes que expongan los endpoints
* *app* es para definir los tipos de datos de dominio (envio,cliente,contrato,etc)
* *pkg* es para definir paquetes que este proyecto quiera exportar y puedan importarse desde otros. 
* *deployment* es para colocar scripts, docker-compose, helm charts, o yaml con descripciones de objetos de Kubernetes/OCP que sirvan para desplegar local para development o para subir al container orchestrator (OCP, por ejemplo)

Usar lo que se precise, el resto borrarlo.

## Maintainer

chebi - sorfino@andreani.com
Proyecto_APIsClientes@andreani.com