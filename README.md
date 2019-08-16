This is a sample  code for  dependency injection and mocking. 

1. Doing depedency injection 
2. Mocking the DB call 
3. Ignorning error handling 
4. Ignorning loging
5. There are many libraries for mocking and generating mocks, i have not used it.  




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