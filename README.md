# gin-rest-api-boilerplate

An easy to use, extensible boilerplate for Go applications with Gin framework.

### Tools

- [Gin](https://github.com/gin-gonic/gin 'Gin') - Gin Web Framework
- [Gorm](https://gorm.io/ 'Gorm') - The fantastic ORM library for Golang
- [Logrus](https://github.com/sirupsen/logrus 'Logrus') - Logrus is a structured logger for Go (golang), completely API
  compatible with the standard library logger.
- [Viper](https://github.com/spf13/viper 'Viper') - Go configuration with fangs.
- [Health Check](https://github.com/heptiolabs/healthcheck 'Health Check') - Implementing Kubernetes liveness and
  readiness probe handlers

### File Structure

    ..
    ├── cmd                            # Main applications for this project.
        ├── api                        # Project application name  ex. gin-rest-api
    ├── config                         # Configuration management using Viper. It checks for relevant environment variables, fallback to config files and finally to defailt values (hardcoded)
    ├── db                             # Configuration connection database
    ├── internal                       # Private application and library code. This is the code you don't want others importing in their applications or libraries
    |   |   ├── dto                    # Data transfer object for map object in application
    │   │   └── errors                 # Custom error  application
    │   │   └── handler                # Handler http status code and error
    │   │   └── models                 # Model struct type for data model that determines the logical structure of a database.
    │   │   └── repository             # Repository for business logic query database
    │   │   └── routes                 # Router gin
    │   │   └── services               # Service for business logic layer
    │   │                              #
    ├── pkg                            # Library code that's ok to use by external applications
    ├── Makefile                       #
    └── ...

#### Makefile

The `make` file is mainly used as a "shortcut" to commonly used commands and tools such as docker, auto code generation
etc.  
For all available commands, please checkout the [Makefile](Makefile 'Makefile').
