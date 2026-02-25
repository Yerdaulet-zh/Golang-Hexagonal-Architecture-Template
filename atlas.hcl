data "external_schema" "gorm" {
  program = [
    "go",
    "run",
    "-mod=mod",
    "./internal/adapters/repository/postgre/loader/loader.go",
  ]
}

data "composite_schema" "app" {
  schema "public" {
    url = data.external_schema.gorm.url
  }
}

env "local" {
  url = "postgres://admin:password@localhost:5432/myapp?sslmode=disable"

  src = data.composite_schema.app.url

  dev = "docker://postgres/17/dev?search_path=public"

  migration {
    dir = "file://internal/adapters/repository/postgre/migrations"
  }

  format {
    migrate {
      diff = "{{ sql . \"  \" }}"
    }
  }
}