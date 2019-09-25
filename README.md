# NUXEO_MID

Api intermediaria del cliente y el api de nuxeo.

Esta api tiene inicialmente la finalidad de disparar el flujo a los documentos de nuxeo, el cual luego podra ser finalizado si se rechaza o ser consitnuar el flujo y permitir que los documentos queden publicos.

Inicialmente es un requerimiento para el cliente de `CUMPLIDOS` , peri se espera que se pueda usar de forma generica.

Los siguientes diagramas de secuencia evidencian el proceso de disparo de flujo y de aprovacion.


# Instalación
Para instalar el proyecto de debe relizar lo siguientes pasos:

Ejecutar desde la terminal 'go get repositorio':
```shell 
go get github.com/udistrital/nuxeo_mid
```

# Ejecución del proyecto


- Ejecutar: 
```shell 
bee run
```
- O si se quiere ejecutar el swager:

```shell 
bee run -downdoc=true -gendoc=true
```

### EndPoints
|  Funcion |Tipo de peticion                  |Parametros| Endpoint |
|----------------|------------------------|---------------------|-------------------|
| **Disparar el flujo a un determinado documento** | **POST** |`DocID`   id del documento| ```workflow/[DocID]```|
| **Eliminar un flujo de un documento** | **DELETE** | `DocID`   id del documento |```workflow/[DocID]```|
| **Aprovar un documento** | **POST** | `DocID`   id del documento |```validacion/[DocID]```|


---

### Disparo de flujo

![disparo FLUJO](https://user-images.githubusercontent.com/28914781/65219366-08ce8f00-da7e-11e9-8a82-7832c0ee5384.png)

---

### Aprovacion de documentos.

![aprobacion doc](https://user-images.githubusercontent.com/28914781/65219477-3e737800-da7e-11e9-8192-4600d4c8f7ef.png)