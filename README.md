# Firebolt Client

This repository contains the necessary code to connect any application to the **Firebolt** platform ([firebolt.io](https://firebolt.io)), an online Data Warehouse. The client is implemented using **Golang** ([golang.org](https://golang.org)) and utilizes the plugin provided by Firebolt for this purpose.

## Components

To work effectively with Firebolt, it's important to thoroughly read the documentation available at the following URL: [Firebolt Documentation](https://docs.firebolt.io/). For this simple example, the following concepts are essential, and understanding how they function within the platform is crucial.

- **Account**: Represents your account. When registering your "organization," you'll provide the necessary details. It's important to remember the name (identifier) you configure.
- **Service Account**: This acts as an API Key and is used for authentication purposes in applications. Identifiers and secrets will be provided. To obtain a secret for a Service Account, it must be created first, followed by regenerating the secret.
- **User**: This is the user that will be used for integration. Once the databases, engines, and Service Accounts are created, the user will need to be linked to them. This is where permissions are configured. To insert data into a database, a new role with the appropriate permissions must be created.
- **Database**: As the name suggests, this is where the data is stored. When inserting data, note that Firebolt does not automatically create tables, so they must be created manually.
    Here's an example:
    ```sql
    CREATE FACT TABLE events (
        id INT,
        type VARCHAR(255),
        user_id INT,
        ts TIMESTAMP,
        metadata VARCHAR(255)
    );
    GRANT INSERT ON TABLE events TO [USER_NAME];
    ```
    > **_Note_**: Firebolt uses specific tags, which users should explore further (e.g., FACT).

- **Engine**: The engine is the search and execution platform. For any operations against databases, it must be running. This can be a drawback for real-time ingestion and querying since running the engine 24/7 can be costly. It's designed for **batch processing** at specific times, making it suitable for data processing and use with **BI** tools.

## Functions

Two types of functions have been implemented:

### CheckConnection
This function checks the connectivity to the platform using the provided credentials.

```go
dsn := fmt.Sprintf("firebolt:///%s?account_name=%s&client_id=%s&client_secret=%s", databaseName, accountName, clientId, clientSecret)

// Open the Firebolt connection
db, err := sql.Open("firebolt", dsn)
if err != nil {
  log.Fatal("Error opening connection:", err)
}

// Ping the database to check the connection
err = db.Ping()
if err != nil {
  log.Fatal("Failed to connect to Firebolt:", err)
}

fmt.Println("Connection successful!")
defer db.Close()
```

> **_Nota_**: Handling the DNS can be tricky, and the documentation lacks clarity in this area.

### Insert
This function inserts an entry into a specific 'database.table'.

```go
// Firebolt connection string
dsn := fmt.Sprintf("firebolt:///%s?account_name=%s&client_id=%s&client_secret=%s&engine=%s", databaseName, accountName, clientId, clientSecret, engine)

// Open the connection
db, err := sql.Open("firebolt", dsn)
if err != nil {
  log.Fatal("Failed to connect to Firebolt:", err)
}
defer db.Close()

// Insert event data
query := `
  INSERT INTO events (id, type, user_id, ts, metadata) 
  VALUES (?, ?, ?, ?, ?)
`

_, err = db.Exec(query, 123, "click", 456, "2024-09-02 12:34:56", "sample metadata")
if err != nil {
  log.Fatal("Failed to insert event:", err)
}

fmt.Println("Event ingested successfully")
```

## Conclusions
My conclusion is that Firebolt works very well as a Data Warehouse and is a highly recommended option due to its pricing and efficiency. However, it may be more suited for technical users rather than business users. While integration is simple, it is not necessarily easy.

Its performance is excellent, very fast, and even the engine startup time is quick compared to competitors. However, if you're looking for a platform for real-time operations, I wouldn't recommend it, as it's primarily designed for batch processes—i.e., ingestion through batch processes and queries via **ETLs** and **BI** visualization tools.

