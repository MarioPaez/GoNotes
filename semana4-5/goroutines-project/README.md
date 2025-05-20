# Concurrent Task Processing System with Reporting

## Goal
The goal of this mini project is to gain a solid understanding of concurrency in Go and become familiar with goroutines, channels, concurrency patterns like fan-in and fan-out, mutexes, waitgroups, workers, and an introduction to efficient concurrent system design.

## General Idea
The project consists of implementing a concurrent task processing system with reporting.

The goal is to simulate a backend that receives many tasks simultaneously. We want to process these tasks using multiple workers (concurrently) to reduce total execution time.

Each task will take a random amount of time between 0 and 1 second with a 20% chance of failure. The system should collect real-time statistics:
1. How many tasks were processed by each worker.
2. How long each task took.
3. How many tasks failed.

## Modules

| Module         | Responsibility                                        |  
| -------------- | ----------------------------------------------------- | 
| **Dispatcher** | Creates workers and assigns tasks to them            |
| **Worker**     | Processes tasks concurrently                         |
| **Reporter**   | Collects results and statistics                      |
| **Main loop**  | Orchestrates everything, sends tasks, waits for completion |

## Functional Requirements

1. Process a list of tasks using multiple workers (fan-out).
2. Collect results into a single channel (fan-in).
3. Use sync.WaitGroup to wait for complete processing. (Is it really necessary?)
4. Protect access to statistics using sync.Mutex.
5. Measure duration of each task and overall execution.
6. Count successful and failed tasks.
7. Display a final summary (global and per-worker stats).
8. Use channels properly to avoid deadlocks and race conditions.

## Execution Flow

1. Generate `N` tasks with unique IDs.
2. Launch `M` workers that receive tasks via a channel.
3. Each worker simulates processing the task (random time between 0–1s and 20% failure rate).
4. Results are sent through a result channel.
5. A goroutine listens to results and updates statistics protected by `sync.Mutex`.
6. Once everything finishes, print a complete summary.



---
---

_Spanish version below_


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
