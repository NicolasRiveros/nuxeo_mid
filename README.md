# NUXEO_MID

Api intermediaria del cliente y el api de nuxeo.

Esta api tiene inicialmente la finalidad de disparar el flujo a los documentos de nuxeo, el cual luego podra ser finalizado si se rechaza o ser consitnuar el flujo y permitir que los documentos queden publicos.

Inicialmente es un requerimiento para el cliente de `CUMPLIDOS` , peri se espera que se pueda usar de forma generica.

Los siguientes diagramas de secuencia evidencian el proceso de disparo de flujo y de aprovacion.

---

### Disparo de flujo

![disparo FLUJO](https://user-images.githubusercontent.com/28914781/65219366-08ce8f00-da7e-11e9-8a82-7832c0ee5384.png)

---

### Aprovacion de documentos.

![aprobacion doc](https://user-images.githubusercontent.com/28914781/65219477-3e737800-da7e-11e9-8192-4600d4c8f7ef.png)



---

falta contemplar la accion a seguir al momento de que no se aprobe el documento para su publicacion.