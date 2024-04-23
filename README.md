# Lab2SD
Benjamín Gutierrez 202004621-2
Sofía Parada Hormazábal 202004671-9



Instrucciones de ejecucuión:

Para ejecutar ambos codigos, se necesitan dos terminales powershell (en windows o vsc) que se encuentren en el directorion principal donde se encuentran las carpetas con sus respectivos archivos go
y el dockerfile.

-Para dockerizar Central se debe escribir en terminal: docker build -t tierra .
-Para correr el docker se debe escribir en terminal: docker run -p 8080:8080 tierra

Una vez el Central ejecutandose, se debe correr el Equipos.go, para esto se debe escribir en otra terminal: go run Equipos/Equipos.go

Si se desea cancelar la ejecucion antes, para Equipos.go se utiliza CTRL + C y para Central.go, se debe eliminar la ejecucuón desde la aplicación Dockert Desktop
