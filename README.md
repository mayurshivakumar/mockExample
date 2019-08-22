This is a sample  code for  dependency injection and mocking. 

1. Doing depedency injection 
2. Mocking the DB call for unit testing
3. Generated mock files using mockery. I also have test that do not use it. 
4. Minimal loging. 

Installation

Install docker 
Then pull down mssql server lastest and  get it up and running 

```
sudo docker pull microsoft/mssql-server-linux:2017-latest
docker run -e 'HOMEBREW_NO_ENV_FILTERING=1' -e 'ACCEPT_EULA=Y' -e 'SA_PASSWORD=yourStrong(!)Password' -p 1433:1433 -d microsoft/mssql-server-linux
```

Install azure data studio GUI 

```
brew cask install azure-data-studio
```

If you want to install sql command line unility knock your self out 

```
brew tap microsoft/mssql-release https://github.com/Microsoft/homebrew-mssql-release
brew install --no-sandbox msodbcsql17 mssql-tools
```

DB script to get it up and running

```
CREATE DATABASE SampleDB;
GO

CREATE SCHEMA TestSchema;
GO

CREATE TABLE TestSchema.Employees (
  Id INT IDENTITY(1,1) NOT NULL PRIMARY KEY,
  Name NVARCHAR(50),
  Location NVARCHAR(50)
);
GO

INSERT INTO TestSchema.Employees (Name, Location) VALUES
(N'Jared', N'Australia'),
(N'Nikita', N'India'),
(N'Tom', N'Germany');
GO

SELECT * FROM TestSchema.Employees;
GO
```

Dependency injection
1) Dependency injection is a technique whereby one object supplies the dependencies of another object.
2) Things that are commonly injected to handle. Database connection, loggers, configurations....
3) Can be achieved through constructor, setter methods and interfaces


Advantages
1) Helps in unit testing. Along with mocking.
2) Extending the application becomes easy.
3) Helps create loosely coupled objects.

Disadvantages
1) Can be little complex to set up and learn initially.
2) Usually implemented with reflection or dynamic programming. Does not work well with IDE automation 
