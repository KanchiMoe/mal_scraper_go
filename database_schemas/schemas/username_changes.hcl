table "username_changes" {
    schema = schema.public

    column "uuid" {
        type = uuid
        null = false
    }
    column "id" {
      type = bigint
      null = false
    }
    column "old_username" {
        type = text
        null = false
    }
    column "new_username" {
        type = text
        null = false
    }
    column "timestamp" {
        type = timestamptz
        null = false
    }    

    primary_key {
        columns = [
            column.uuid
        ]
    }
}