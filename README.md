# nuxeo_mid

Api intermediaria del cliente y el api de nuxeo.

Esta api tiene inicialmente la finalidad de disparar el flujo a los documentos de nuxeo, el cual luego podra ser finalizado si se rechaza o ser continua el flujo y permitir que los documentos queden publicos.

Inicialmente es un requerimiento para el cliente de `CUMPLIDOS` , pero se espera que se pueda usar de forma generica.

Los siguientes diagramas de secuencia evidencian el proceso de disparo de flujo y de aprobacion.

## Especificaciones Técnicas

### Tecnologías Implementadas y Versiones
* [Golang](https://github.com/udistrital/introduccion_oas/blob/master/instalacion_de_herramientas/golang.md)
* [BeeGo](https://github.com/udistrital/introduccion_oas/blob/master/instalacion_de_herramientas/beego.md)
* [Docker](https://docs.docker.com/engine/install/ubuntu/)
* [Docker Compose](https://docs.docker.com/compose/)

### Variables de Entorno
```shell
PORT_NUXEO_MID=[Descripción]
USER_NUXEO=[Descripción]
PASSWORD_NUXEO=[Descripción]
```

**NOTA:** Las variables se pueden ver en el fichero conf/app.conf y están identificadas con _NUXEO...

### Ejecución del Proyecto
```shell
#1. Obtener el repositorio con Go
go get github.com/udistrital/nuxeo_mid

#2. Moverse a la carpeta del repositorio
cd $GOPATH/src/github.com/udistrital/nuxeo_mid

# 3. Moverse a la rama **develop**
git pull origin develop && git checkout develop

# 4. alimentar todas las variables de entorno que utiliza el proyecto.
PORT_NUXEO_MID=8080 USER_NUXEO=XXX PASSWORD_NUXEO=XXX bee run
```

### EndPoints
|  Funcion |Tipo de peticion                  |Parametros| Endpoint |
|----------------|------------------------|---------------------|-------------------|
| **Disparar el flujo a un determinado documento** | **POST** |`DocID`   id del documento| ```workflow/[DocID]```|
| **Eliminar un flujo de un documento** | **DELETE** | `DocID`   id del documento |```workflow/[DocID]```|
| **Aprovar un documento** | **POST** | `DocID`   id del documento |```validacion/[DocID]```|

### Disparo de flujo

![disparo FLUJO](https://user-images.githubusercontent.com/28914781/65219366-08ce8f00-da7e-11e9-8a82-7832c0ee5384.png)

### Aprobación de documentos.

![aprobacion doc](https://user-images.githubusercontent.com/28914781/65219477-3e737800-da7e-11e9-8192-4600d4c8f7ef.png)


## Estado CI

| Develop | Relese 0.0.1 | Master |
| -- | -- | -- |
| [![Build Status](https://hubci.portaloas.udistrital.edu.co/api/badges/udistrital/nuxeo_mid/status.svg?ref=refs/heads/develop)](https://hubci.portaloas.udistrital.edu.co/udistrital/nuxeo_mid) | [![Build Status](https://hubci.portaloas.udistrital.edu.co/api/badges/udistrital/nuxeo_mid/status.svg?ref=refs/heads/release/0.0.1)](https://hubci.portaloas.udistrital.edu.co/udistrital/nuxeo_mid) | [![Build Status](https://hubci.portaloas.udistrital.edu.co/api/badges/udistrital/nuxeo_mid/status.svg)](https://hubci.portaloas.udistrital.edu.co/udistrital/nuxeo_mid) |


## Licencia

This file is part of nuxeo_mid.

nuxeo_mid is free software: you can redistribute it and/or modify it under the terms of the GNU General Public License as published by the Free Software Foundation, either version 3 of the License, or (at your option) any later version.

nuxeo_mid is distributed in the hope that it will be useful, but WITHOUT ANY WARRANTY; without even the implied warranty of MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the GNU General Public License for more details.

You should have received a copy of the GNU General Public License along with nuxeo_mid. If not, see https://www.gnu.org/licenses/.
