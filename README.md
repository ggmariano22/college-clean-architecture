![Go](https://img.shields.io/badge/go-%2300ADD8.svg?style=for-the-badge&logo=go&logoColor=white)![Docker](https://img.shields.io/badge/docker-%230db7ed.svg?style=for-the-badge&logo=docker&logoColor=white)

# College management system

### This project implements the concepts of Clean Architecture
![clean architecture image](https://blog.cleancoder.com/uncle-bob/images/2012-08-13-the-clean-architecture/CleanArchitecture.jpg)

---

### Project structure

📦 project
 ┣ 📂 app
 ┃ ┣ 📂 domain
 ┃ ┃ ┣ 📂 entities
 ┃ ┃ ┗ 📂 useCases
 ┃ ┣ 📂 infrastructure
 ┃ ┣ 📂 interfaces
 ┃ ┗ 📂 vendor
 ┣ 📂 cmd
 ┃ ┗ 📜 main.go
 ┣ 📂 config

---
 ### How to run

If it's the first time you are running this project, run the command `make docker-install`.
Otherwise, just run `make docker-up`.

---

### Testing
In progress :construction_worker:

---

### Future features
- Resources
    - Resource to get student by filter
    - Resource to edit a student
    - Resource to delete a student (set inactive date field)
    - Resource to add additional activities
    - Resource to create a teacher
    - Resource to get all teachers
    - Resource to get teacher by filter
    - Resource to edit a teacher
    - Resource to delete a teacher (set inactive date field)
    - Resource to create a class
    - Resource to register students at class
- Business rules
    - Grade system
    - Validate if that received student isn't registered in another course already
    - Include CPF validation
    - Include additional activities
- Project infrastructure
    - Add Air for live reload when make some changes in the code :heavy_check_mark:
    - Split Dockerfile into dev and prod :heavy_check_mark: