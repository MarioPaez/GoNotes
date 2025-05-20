# Sistema concurrente de procesamiento de tareas con reporting

## Objetivo
El objetivo de realizar este mini proyecto es entender bien la concurrencia en Go y hacernos familiares con las goroutines, channels, patrones de concurrencia como fan-in y fan-out, mutexes, waitgroups, workers y una iniciación sobre el diseño de sistemas concurrentes eficientes.

## Idea general
El proyecto consiste en realizar un sistema concurrente de procesamiento de tareas con reporting.

El objetivo es simular un backend el cual recibe muchas tareas a la vez. Queremos procesar estas tareas con múltiples workers (concurrentemente) para así ahorrar tiempo de ejecución total.

Cada tarea tomará un tiempo aleatorio entre 0 y 1s con un 20% de fallo en la tarea. El sistema debe recolectar estadísticas en tiempo real:
1. Cuántas tareas fueron procesadas por cada worker.
2. Cuánto tiempo tardó cada una de las tareas.
3. Cuántas fallaron.

## Módulos

| Módulo         | Responsabilidad                                        |  
| -------------- | ------------------------------------------------------ | 
| **Dispatcher** | Crea workers y les asigna tareas                       |
| **Worker**     | Procesa tareas concurrentemente                        |
| **Reporter**   | Recoge resultados y estadísticas                       |
| **Main loop**  | Orquesta todo, envía tareas, espera a que todo termine |

## Requisitos funcionales

1. Procesar una lista de tareas con múltiples workers (fan-out).
2. Recolectar resultados en un canal único (fan-in).
3. Usar sync.WaitGroup para esperar el procesamiento completo. (Es realmente necesario?)
4. Proteger acceso a estadísticas con sync.Mutex.
5. Medir duración de cada tarea y global.
6. Contar tareas exitosas y fallidas.
7. Mostrar un resumen final (estadísticas globales + por worker).
8. Usar canales correctamente para evitar deadlocks y race conditions.

## Flujo de ejecución

1. Genera `N` tareas con IDs únicos.
2. Lanza `M` workers que reciben tareas por un canal.
3. Cada worker simula procesar la tarea (con tiempo aleatorio entre 0-1s y un 20% de probabilidad de fallo).
4. Los resultados se mandan por otro canal.
5. Una goroutine escucha los resultados y actualiza estadísticas protegidas por `sync.Mutex`.
6. Al terminar, se imprime un resumen completo.
