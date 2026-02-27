# Chaos-Heal 🧠🔥  
A Chaos-First Self-Healing Distributed System

## Overview

Chaos-Heal is a distributed system designed with failure as a first-class citizen.

The system:
- Detects node failures via heartbeat monitoring
- Elects a leader using majority consensus
- Reassigns responsibilities automatically
- Recovers from crashes without manual intervention
- Injects chaos to validate resilience

This project simulates real-world distributed infrastructure patterns used in large-scale production systems.

---

## Architecture

Client → Service Nodes → Failure Detector → Leader Election → Recovery Engine → Chaos Injector

Core components:

- **Service Nodes** – Workers responsible for processing tasks or owning shards
- **Heartbeat Monitor** – Detects node liveness
- **Leader Election Module** – Ensures a single recovery authority
- **Recovery Engine** – Reassigns responsibilities upon failure
- **Chaos Engine** – Injects faults (node kill, delay, network drop)

---

## Design Goals

- Expect failure
- Avoid split-brain
- Ensure eventual consistency
- Prefer safety over availability during partitions
- Maintain minimal global state

---

## Tech Stack

- Language: Go
- Communication: gRPC
- Containerization: Docker
- Orchestration: Docker Compose
- Chaos Injection: Custom scripts

---

## Running the System

```bash
docker compose up --build
