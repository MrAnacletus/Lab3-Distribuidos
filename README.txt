Integrantes:
    Sebastián Sepúlveda Barría      	201873512-4
    Salvador Fuentes Azocar         	201873511-6
    Manuel Valenzuela Villarroel     	201873552-3

Para este trabajo se utilizó GOLANG, gRPC, PROTOBUFS

Las maquinas virtuales tienen asignados los siguientes procesos

	Dist 77: Broker
	Dist 78: Fulcrum 1 - Informante 1
	Dist 79: Fulcrum 2 - Informante 2
	Dist 80: Fulcrum 3 - Leia

Para la correcta ejecución de esta tarea se deben seguir los siguientes pasos:
.- Ejecutar proceso broker con make broker en la maquina 77
.- Ejecutar proceso Fulcrum 1 con make fulcrum, luego ingresar un 1 en la maquina 78
.- Ejecutar proceso Fulcrum 2 con make fulcrum, luego ingresar un 2 en la maquina 79
.- Ejecutar proceso Fulcrum 3 con make fulcrum, luego ingresar un 3 en la maquina 80
.- Ejecutar proceso Informante 1 con make informante en la maquina 78
.- Ejecutar proceso Informante 2 con make informante en la maquina 79
.- Ejecutar proceso Leia con make leia en la maquina 80

Habiendo hecho todo lo anterior EN ORDEN se puede proceder a utilizar la tarea

Asunciones:
	-Se utilizan vectores en memoria, por ende los archivos preexistentes de planetas y de logs de planetas podrian generar problemas,
	 esto debido a que no se recoje esta informacion al comenzar una nueva ejecución.
	-Los nombres de archivos tienen un numero junto a ellos, digase [log]{planeta}{indice}.txt, esto en tanto la tarea se programó en
	 windows utilizando solo una maquina y por conveniencia para nosotros poder probarla se hizo así.
	-Solo se aceptan entradas VALIDAS.
	-El orden y la posición de los servers Fulcrum debe ser la expuesta pues por dentro el orden de los servidores importa.
	-El merge utilizado fue un merge simple en el cual se pegan las lineas de comandos una tras otra siguiendo una jerarquía,
	 esta jerarquía proviene del orden de los fulcrum más quien recibe la advertencia de la necesidad de un merge.
