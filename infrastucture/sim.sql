CREATE DATABASE SIMCOMPANIES;
GO

USE SIMCOMPANIES
GO

CREATE TABLE [Resource](
    [id] INT PRIMARY KEY IDENTITY,
    [name] NVARCHAR(100) NOT NULL,
    [image] NVARCHAR(100) NOT NULL,
    [transporation] FLOAT NOT NULL,
    [retailable] BIT NOT NULL,
    [research] BIT NOT NULL,
    [exchangeTradable] BIT NOT NULL,
    [realmAvailable] BIT NOT NULL
)
SET IDENTITY_INSERT [Resource] OFF
GO

DROP TABLE Resource

SELECT * FROM [Resource]
GO

INSERT INTO [Resource] (id, [name], [image], []) VALUES
(1, 'Harsh', 'ksdj', 12.2, 1, 1, 0, 1)

-- Delete rows from table 'TableName'
DELETE FROM [Resource]
GO


SELECT UPPER([name]), CASE
        WHEN exchangeTradable = 1
            THEN 'true'
        ELSE 'false'
    END as [VALUE]
 FROM [Resource]
-- WHERE exchangeTradable = 0


-- Renaming Coloum
-- EXEC sp_rename [Resourse.RealmAvailable], [realmAvailable], [COLUMN]
