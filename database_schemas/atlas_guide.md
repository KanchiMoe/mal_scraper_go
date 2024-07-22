inspect
    if you have a database elsewhere that you want to import

```
atlas schema inspect -u postgres://:@localhost:5432/existing_db?sslmode=disable > schema.hcl
```

## diff 

```
atlas migrate diff <comment> \
  --dir "file://migrations" \
  --to "file://schemas" \
  --dev-url "docker://postgres/16/test?search_path=public"
  ```

## apply

```
atlas migrate apply \
--url "postgres://:@localhost:5432/myanimelist?sslmode=disable"
--dir "file://migrations" 
```

# 