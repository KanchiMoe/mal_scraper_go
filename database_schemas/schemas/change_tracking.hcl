table "change_tracking" {
    schema = schema.public

    column "uuid" {
        type = uuid
        null = false
    }
    column "timestamp" {
        type = uuid
        null = false
    }
    column "category" {
        type = text
        null = false
    }
    column "item_id" {
      type = bigint
      null = false
    }
    column "field" {
        type = text
        null = false
    }
    column "old_value" {
      type = text
      null = false
    }
    column "new_value" {
        type = text
      null = false
      
    }



}