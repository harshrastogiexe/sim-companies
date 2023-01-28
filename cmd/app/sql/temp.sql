use SIMCOMPANIES
GO

SELECT * FROM [Resource]

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


CREATE TABLE [Building] (
    id NVARCHAR(10) NOT NULL PRIMARY KEY,
    [name] NVARCHAR(100) NOT NULL,
    [image] NVARCHAR(100) NOT NULL,
    [category] NVARCHAR(100) NOT NULL,
    [cost] INT NOT NULL,
    [robots_needed] FLOAT NOT NULL,
    [realm_available] BIT NOT NULL
)

CREATE TABLE [BuildingProduces] (
    [building_id] NVARCHAR(10) NOT NULL,
    [resource_id] INT NOT NULL,
    FOREIGN KEY (building_id) REFERENCES [Building](id),
    FOREIGN KEY (resource_id) REFERENCES [Resource](id),
    CONSTRAINT UNI_BuildingProduces_buildingIdAndResourcid UNIQUE([building_id], [resource_id])
)

CREATE TABLE [BuildingImages] (
    [id] NVARCHAR(10) NOT NULL ,
    [level] INT NOT NULL,
    [path] NVARCHAR(100) NOT NULL,
    FOREIGN KEY (id) REFERENCES Building(id),
    CONSTRAINT UNI_REF UNIQUE(id, [level])
)


DELETE FROM [Building]
WHERE [Building].[id] IN (
    SELECT [Building].[id] FROM [Building]
    WHERE [Building].[id] NOT IN (
        SELECT DISTINCT [BuildingImages].[id] FROM [BuildingImages]
    )
)

INSERT INTO [BuildingImages]
VALUES
('100', 1, 'path')

SELECT * FROM [BuildingImages]
SELECT * FROM [Building]
-- DELETE FROM Building

SELECT img.* FROM BuildingImages [img], Building[b]
WHERE img.[id] = b.[id] AND b.[id]=5

DROP TABLE [Building]


SELECT * FROM [Resource]
WHERE [Resource].[name] LIKE '%cow%'

EXEC sp_rename [Resourse.RealmAvailable], [realmAvailable], [COLUMN]
