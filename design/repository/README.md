## Repository Pattern.

La preocupación principal del patrón Repository es desacoplar la lógica en dos partes
bien diferenciadas: Por un lado, la lógica de persistencia de datos. Y, por otro
lado, la lógica de dominio o negocio, cuya implementación es independiente de cómo, 
cuándo y dónde se persistan los datos.
De esta forma se consigue que el dominio no conozca a la infraestructura, lo que
permite una clara separación de preocupaciones. Mediante este patrón,
el dominio solo debe ocuparse de implementar lógica del negocio.

### ¿Cómo luce el patrón Repository?
Básicamente, se trata de una abstracción del sistema de persistencia, es decir, una
interfaz con sus diferentes implementaciones.
![repositoy_pattern_uml](img/repository_pattern.drawio.png.png)

* <span style="color:blue">**Dominio**</span>
* <span style="color:green">**Infraestructura**</span>

Este patrón permite que se cumpla el primer principio de **SOLID**: **Single
Responsability Principle**. Es decir, que cada *cosa* tenga solo una responsabilidad
bien clara y definida. Un caso de uso acoplado a la infraestructura conlleva a que
el dominio tenga que lidiar con sus propias reglas de negocio y, además, también
cargue con la responsabilidad de operar sobre el sistema de persistencia.

Por otra parte, este patrón también alienta la regla principal de las **Clean
Architectures**: Separar o desacoplar la capa de dominio de la capa de infraestructura.
Asegurando que el dominio no conoce ni lo más mínimo de la infraestructura. Y, de 
forma inversa, la infraestructura conoce cada detalle sobre el dominio, tal que, es
la capa exterior quien se acopla a las necesidades de las capas más interiores.
Esto permite que el dominio únicamente se modifique cuando surjan nuevas reglas de
negocio.
Incluso, dicho patrón se presentó por primera vez en el libro Domain Driven Design
escrito **Eric Evans** y publicado 2004.

Cómo vemos, la propuesta es que tanto **Entidades** como **Agregados**, ambos
objetos accedan a los datos comunicándose con el Repository.
Un punto importante a tener en cuenta es que tanto las **Entidades, Agregados y
Repository** son parte de nuestro **Dominio**(color amarillo). Mientras que son
las implementaciones de la interfaz son quienes se encuentran en la capa más
externa, es decir, infraestructura(color verde).
Esto tiene sentido, puesto que es el dominio quién define el contrato que se debe
cumplir, ya que es el dominio quien conoce qué operaciones son necesarias sobre
el sistema de persistencia. No se debe confundir la definición de la Interfaz con
la implementación de la Interfaz.

![DDD_diagram_repository](img/DDD_Diagram.png)

### ¿Cuáles son los beneficios del repository pattern?
Una vez entendida la solución podemos preguntarnos qué valor aporta este patrón
de diseño. Es que nos permite desacoplar el acceso y actualización de datos
del dominio, pero, ¿de qué sirve dicho desacople?
Un argumento bastante lógico sería pensar en cuántas veces se cambia de base de
datos en un proyecto real. La respuesta es que un cambio de tal magnitud no es algo
frecuente. ¿Entonces, si no está previsto cambiar de base de datos en el futuro,
significa que el patrón no aporta ningún valor? Definitivamente no.
Existe un punto clave que impulsa el uso de repositories en cualquier proyecto:
**Testing**.

Imaginemos que tenemos un **Servicio de dominio** que da de alta nuevos usuarios,
podría llamarse UserRegister por ejemplo.
Ahora bien, este servicio tiene numerosos casos de uso (contraseña inválida, nombre
de usuario ya existe, nombre de usuario no válido, email inexistente,
datos obligatorios vacíos, etc). Podríamos realizar los test de este servicio
utilizando una base de datos real, ya sea copia de producción o una base específica
para testing. Cómo sea, deberíamos ocuparnos de actualizar las entradas de las
tablas antes de correr los tests, de forma que cada caso de uso encuentre los
valores correctos en el sistema de persistencia.

Por otro lado, también deberíamos lidiar con la latencia que implica levantar
una base de datos real cada vez que corremos los tests automáticos.

Analizando este escenario detenidamente, nos encontramos con que tenemos un servicio
de dominio que encapsula el %100 de casos de uso posibles y los tests automáticos
de dicho servicio se ejecutan sobre una base de datos real. Cada vez que agreguemos
un nuevo caso de uso a los tests, el tiempo de ejecución de los tests podría crecer
exponencialmente, ya que cada caso de uso seguramente necesite de una o varias
operaciones CRUD sobre la base de datos.

Este crecimiento en los tiempos de ejecución provoca que introducir una nueva feature
o un simple refactoring se vuelva algo tedioso. Imagine si cada vez que ejecutamos
los tests automáticos tuviéramos que esperar unos cuantos minutos u horas. Sería
un despropósito de recursos.
Pues en proyecto de gran magnitud los test son gigantes y no deberían perder tiempo y recursos
en levantar, limpiar y actualizar la base de datos.

En todo caso, la implementación concreta de cada base de datos debería testearse
en la capa de infraestructura con casos de uso reducidos, pero no en la capa de
dominio.

Este escenario nos da lugar a utilizar un Repository para operar dentro del
servicio de dominio, utilizando una base de datos en memoria, redis, un array o lo
que se nos ocurra y sea útil para correr los test.

Precisamente ese es el poder del Repository, la **cambiabilidad** que ofrece.
Hoy podríamos estar utilizando un sistema de persistencia A, mañana uno B,
pero a su vez utilizando uno C en el servidor de QA y uno D en el servidor de PRD.

### Ejemplo:
Luego de esta explicación teórica, los invito revisar la explicación práctica.
La misma consta de un servicio de dominio que pasa por dos estados diferentes.
El primero sin el patrón repository y el segundo implementando dicho patrón.


