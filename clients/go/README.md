Airflow Go client
=================

Go Airflow OpenAPI client generated from openapi spec.


Usage
-----

```go
import (
    "github.com/apache/airflow/clients/go/airflow"
)

func main() {
    cli := airflow.NewAPIClient(&airflow.Configuration{})
    variable, _, err := client.VariableApi.GetVariable(context.TODO(), "foo")
    if err != nil {
        fmt.Println(variable)
    } else {
        fmt.Println(err)
    }
}
```

See [README](./airflow/README.md) for full client API documentation.


Release process
---------------

Go client is versioned using [semantic import
versioning](https://blog.golang.org/versioning-proposal).

To release a new version `1.x.y`, simply push a new tag to Airflow repo named
`clients/go/v1.x.y`.

Major version upgrade requires changing package import path based on semantic
import versioning, which needs to be done manually.
