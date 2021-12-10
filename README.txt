Integrantes:
    Sebastián Sepúlveda Barría      201873512-4
    Salvador Fuentes Azocar         201873511-6
    Manuel Valenzuela Villarroel     201873552-3

Para este trabajo se utilizó GOLANG, gRPC, PROTOBUFS, rabbitMQ¿?

Las conexiones entre: Lider - Cliente, Lider - Pozo, Lider - NameNode, por medio de gRPC están hechas.


Los 3 juegos pedidos han sido implementados, si el jugador principal muere en alguna etapa, el juego sigue automaticamente.

Para correr los diferentes servicios de GO se deben utilizar los siguientes comandos Makefile:
    En la maquina virtual dist77 se encuentra alojado el servicio Cliente, que se debe ejecutar con: make run Cliente
    En la maquina virtual dist78 se encuentra alojado el servicio Lider, que se debe ejecutar con: make run Lider
    En la maquina virtual dist79 se encuentra alojado el servicio NameNode, que se debe ejecutar con: make run NameNode
    En la maquina virtual dist80 se encuentra alojado el servicio Pozo, que se debe ejecutar con: make run Pozo

El archivo Makefile se encuentra dentro de la carpeta pruebas, que es el entorno de desarrollo que se utilizó para esta entrega.

Recordar que el servidor rabbitMQ debe estar arriba!!!!!!
comando utilizado :
sudo systemctl start rabbitmq-server
opcion 1
