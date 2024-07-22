table "users" {
    schema = schema.public

    column "id" {
        type = bigint
        null = false
    }
    column "username" {
      type = text
      null = false
    }
    primary_key {
        columns = [
            column.id
        ]
    }
}