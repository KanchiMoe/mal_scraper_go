table "genres" {
    schema = schema.public

    column "uuid" {
        type = uuid
        null = false
    }
    column "id" {
      type = smallint
      null = false
    }
    column "name" {
        type = text
        null = false
    }
    column "description" {
        type = text
        null = true
    }
    column "count" {
        type = int
        null = false
    }
    column "last_interacted" {
        type = timestamptz
        null = false
    }
    primary_key {
        columns = [
            column.uuid
        ]
    }
}

