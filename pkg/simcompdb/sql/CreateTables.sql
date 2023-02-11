DROP TABLE needed_for;
GO
DROP TABLE does_produce;
GO
DROP TABLE does_sell;
GO
DROP TABLE building_level_image;
GO
DROP TABLE improves_quality_of;
GO
DROP TABLE building_main;
GO
DROP TABLE resource_main;
GO
DROP TABLE resource_base;
GO
DROP TABLE building_base;

USE SIMCOMPANIES
GO

/*
************************************************
------------------Bulding Base------------------
************************************************
*/
SET ANSI_NULLS ON
GO
SET QUOTED_IDENTIFIER ON
GO
CREATE TABLE [dbo].[building_base]
(
    [id] [nvarchar](10) NOT NULL,
    [name] [nvarchar](100) NOT NULL,
    [image] [nvarchar](200) NOT NULL,
    [category] [nvarchar](100) NOT NULL,
    [cost] [int] NOT NULL,
    [robots_needed] [float] NOT NULL,
    [realm_available] [bit] NOT NULL
) ON [PRIMARY]
GO
SET ANSI_PADDING ON
GO
ALTER TABLE [dbo].[building_base] ADD PRIMARY KEY CLUSTERED
(
	[id] ASC
)WITH (PAD_INDEX = OFF, STATISTICS_NORECOMPUTE = OFF, SORT_IN_TEMPDB = OFF, IGNORE_DUP_KEY = OFF, ONLINE = OFF, ALLOW_ROW_LOCKS = ON, ALLOW_PAGE_LOCKS = ON) ON [PRIMARY]
GO

/*
************************************************
 Resource Base
************************************************
*/

SET ANSI_NULLS ON
GO
SET QUOTED_IDENTIFIER ON
GO
CREATE TABLE [dbo].[resource_base]
(
    [id] [bigint] IDENTITY(1,1) NOT NULL,
    [name] [nvarchar](200) NOT NULL,
    [image] [nvarchar](200) NOT NULL,
    [transportation] [float] NOT NULL,
    [retailable] [bit] NOT NULL,
    [research] [bit] NOT NULL,
    [exchange_tradable] [bit] NOT NULL,
    [realm_available] [bit] NOT NULL
) ON [PRIMARY]
GO
ALTER TABLE [dbo].[resource_base] ADD PRIMARY KEY CLUSTERED
(
	[id] ASC
)WITH (PAD_INDEX = OFF, STATISTICS_NORECOMPUTE = OFF, SORT_IN_TEMPDB = OFF, IGNORE_DUP_KEY = OFF, ONLINE = OFF, ALLOW_ROW_LOCKS = ON, ALLOW_PAGE_LOCKS = ON) ON [PRIMARY]
GO

/*
************************************************
 Resource Main
************************************************
*/

SET ANSI_NULLS ON
GO
SET QUOTED_IDENTIFIER ON
GO
CREATE TABLE [dbo].[resource_main]
(
    [resource_id] [bigint] IDENTITY(1,1) NOT NULL,
    [sold_at_building_id] [nvarchar](10) NULL,
    [sold_at_restaurant_id] [nvarchar](10) NULL,
    [produced_at_building_id] [nvarchar](10) NULL,
    [transport_needed] [float] NOT NULL,
    [produced_an_hour] [float] NOT NULL,
    [base_salary] [float] NOT NULL,
    [average_retail_price] [float] NULL,
    [market_saturation] [float] NULL,
    [market_saturation_label] [nvarchar](200) NULL,
    [retail_modeling] [nvarchar](200) NULL,
    [store_base_salary] [float] NULL
) ON [PRIMARY]
GO
ALTER TABLE [dbo].[resource_main] ADD PRIMARY KEY CLUSTERED
(
	[resource_id] ASC
)WITH (PAD_INDEX = OFF, STATISTICS_NORECOMPUTE = OFF, SORT_IN_TEMPDB = OFF, IGNORE_DUP_KEY = OFF, ONLINE = OFF, ALLOW_ROW_LOCKS = ON, ALLOW_PAGE_LOCKS = ON) ON [PRIMARY]
GO
ALTER TABLE [dbo].[resource_main]  WITH CHECK ADD  CONSTRAINT [fk_resource_main_produced_at] FOREIGN KEY([produced_at_building_id])
REFERENCES [dbo].[building_base] ([id])
GO
ALTER TABLE [dbo].[resource_main] CHECK CONSTRAINT [fk_resource_main_produced_at]
GO
ALTER TABLE [dbo].[resource_main]  WITH CHECK ADD  CONSTRAINT [fk_resource_main_resource_base] FOREIGN KEY([resource_id])
REFERENCES [dbo].[resource_base] ([id])
GO
ALTER TABLE [dbo].[resource_main] CHECK CONSTRAINT [fk_resource_main_resource_base]
GO
ALTER TABLE [dbo].[resource_main]  WITH CHECK ADD  CONSTRAINT [fk_resource_main_sold_at] FOREIGN KEY([sold_at_building_id])
REFERENCES [dbo].[building_base] ([id])
GO
ALTER TABLE [dbo].[resource_main] CHECK CONSTRAINT [fk_resource_main_sold_at]
GO
ALTER TABLE [dbo].[resource_main]  WITH CHECK ADD  CONSTRAINT [fk_resource_main_sold_at_restaurant] FOREIGN KEY([sold_at_restaurant_id])
REFERENCES [dbo].[building_base] ([id])
GO
ALTER TABLE [dbo].[resource_main] CHECK CONSTRAINT [fk_resource_main_sold_at_restaurant]
GO


/*
************************************************
---------------Building Level Images------------
************************************************
*/
SET ANSI_NULLS ON
GO
SET QUOTED_IDENTIFIER ON
GO
CREATE TABLE [dbo].[building_level_image]
(
    [building_id] [nvarchar](10) NULL,
    [level] [int] NOT NULL,
    [image] [nvarchar](200) NOT NULL
) ON [PRIMARY]
GO
ALTER TABLE [dbo].[building_level_image]  WITH CHECK ADD  CONSTRAINT [fk_building_base_images] FOREIGN KEY([building_id])
REFERENCES [dbo].[building_base] ([id])
GO
ALTER TABLE [dbo].[building_level_image] CHECK CONSTRAINT [fk_building_base_images]
GO

/*
************************************************
 Building Main
************************************************
*/
SET ANSI_NULLS ON
GO
SET QUOTED_IDENTIFIER ON
GO
CREATE TABLE [dbo].[building_main]
(
    [building_id] [nvarchar](10) NOT NULL,
    [cost_units] [float] NOT NULL,
    [hours] [int] NOT NULL,
    [wages] [float] NOT NULL
) ON [PRIMARY]
GO
SET ANSI_PADDING ON
GO
ALTER TABLE [dbo].[building_main] ADD PRIMARY KEY CLUSTERED
(
	[building_id] ASC
)WITH (PAD_INDEX = OFF, STATISTICS_NORECOMPUTE = OFF, SORT_IN_TEMPDB = OFF, IGNORE_DUP_KEY = OFF, ONLINE = OFF, ALLOW_ROW_LOCKS = ON, ALLOW_PAGE_LOCKS = ON) ON [PRIMARY]
GO
ALTER TABLE [dbo].[building_main]  WITH CHECK ADD  CONSTRAINT [fk_building_main_building_base] FOREIGN KEY([building_id])
REFERENCES [dbo].[building_base] ([id])
GO
ALTER TABLE [dbo].[building_main] CHECK CONSTRAINT [fk_building_main_building_base]
GO

/*
************************************************
 Does Produce
************************************************
*/
SET ANSI_NULLS ON
GO
SET QUOTED_IDENTIFIER ON
GO
CREATE TABLE [dbo].[does_produce]
(
    [building_main_building_base_id] [nvarchar](10) NOT NULL,
    [resource_base_id] [bigint] NOT NULL
) ON [PRIMARY]
GO
SET ANSI_PADDING ON
GO
ALTER TABLE [dbo].[does_produce] ADD PRIMARY KEY CLUSTERED
(
	[building_main_building_base_id] ASC,
	[resource_base_id] ASC
)WITH (PAD_INDEX = OFF, STATISTICS_NORECOMPUTE = OFF, SORT_IN_TEMPDB = OFF, IGNORE_DUP_KEY = OFF, ONLINE = OFF, ALLOW_ROW_LOCKS = ON, ALLOW_PAGE_LOCKS = ON) ON [PRIMARY]
GO
ALTER TABLE [dbo].[does_produce]  WITH CHECK ADD  CONSTRAINT [fk_does_produce_building_main] FOREIGN KEY([building_main_building_base_id])
REFERENCES [dbo].[building_main] ([building_id])
GO
ALTER TABLE [dbo].[does_produce] CHECK CONSTRAINT [fk_does_produce_building_main]
GO
ALTER TABLE [dbo].[does_produce]  WITH CHECK ADD  CONSTRAINT [fk_does_produce_resource_base] FOREIGN KEY([resource_base_id])
REFERENCES [dbo].[resource_base] ([id])
GO
ALTER TABLE [dbo].[does_produce] CHECK CONSTRAINT [fk_does_produce_resource_base]
GO

/*
************************************************
 Does Produce
************************************************
*/
SET ANSI_NULLS ON
GO
SET QUOTED_IDENTIFIER ON
GO
CREATE TABLE [dbo].[does_sell]
(
    [building_main_building_base_id] [nvarchar](10) NOT NULL,
    [resource_base_id] [bigint] NOT NULL
) ON [PRIMARY]
GO
SET ANSI_PADDING ON
GO
ALTER TABLE [dbo].[does_sell] ADD PRIMARY KEY CLUSTERED
(
	[building_main_building_base_id] ASC,
	[resource_base_id] ASC
)WITH (PAD_INDEX = OFF, STATISTICS_NORECOMPUTE = OFF, SORT_IN_TEMPDB = OFF, IGNORE_DUP_KEY = OFF, ONLINE = OFF, ALLOW_ROW_LOCKS = ON, ALLOW_PAGE_LOCKS = ON) ON [PRIMARY]
GO
ALTER TABLE [dbo].[does_sell]  WITH CHECK ADD  CONSTRAINT [fk_does_sell_building_main] FOREIGN KEY([building_main_building_base_id])
REFERENCES [dbo].[building_main] ([building_id])
GO
ALTER TABLE [dbo].[does_sell] CHECK CONSTRAINT [fk_does_sell_building_main]
GO
ALTER TABLE [dbo].[does_sell]  WITH CHECK ADD  CONSTRAINT [fk_does_sell_resource_base] FOREIGN KEY([resource_base_id])
REFERENCES [dbo].[resource_base] ([id])
GO
ALTER TABLE [dbo].[does_sell] CHECK CONSTRAINT [fk_does_sell_resource_base]
GO

/*
************************************************
 Improves Quality Of
************************************************
*/
SET ANSI_NULLS ON
GO
SET QUOTED_IDENTIFIER ON
GO
CREATE TABLE [dbo].[improves_quality_of]
(
    [resource_main_resource_base_id] [bigint] NOT NULL,
    [resource_base_id] [bigint] NOT NULL
) ON [PRIMARY]
GO
ALTER TABLE [dbo].[improves_quality_of] ADD PRIMARY KEY CLUSTERED
(
	[resource_main_resource_base_id] ASC,
	[resource_base_id] ASC
)WITH (PAD_INDEX = OFF, STATISTICS_NORECOMPUTE = OFF, SORT_IN_TEMPDB = OFF, IGNORE_DUP_KEY = OFF, ONLINE = OFF, ALLOW_ROW_LOCKS = ON, ALLOW_PAGE_LOCKS = ON) ON [PRIMARY]
GO
ALTER TABLE [dbo].[improves_quality_of]  WITH CHECK ADD  CONSTRAINT [fk_improves_quality_of_resource_base] FOREIGN KEY([resource_base_id])
REFERENCES [dbo].[resource_base] ([id])
GO
ALTER TABLE [dbo].[improves_quality_of] CHECK CONSTRAINT [fk_improves_quality_of_resource_base]
GO
ALTER TABLE [dbo].[improves_quality_of]  WITH CHECK ADD  CONSTRAINT [fk_improves_quality_of_resource_main] FOREIGN KEY([resource_main_resource_base_id])
REFERENCES [dbo].[resource_main] ([resource_id])
GO
ALTER TABLE [dbo].[improves_quality_of] CHECK CONSTRAINT [fk_improves_quality_of_resource_main]
GO

/*
************************************************
 Needed For
************************************************
*/
SET ANSI_NULLS ON
GO
SET QUOTED_IDENTIFIER ON
GO
CREATE TABLE [dbo].[needed_for]
(
    [resource_main_resource_base_id] [bigint] NOT NULL,
    [resource_base_id] [bigint] NOT NULL
) ON [PRIMARY]
GO
ALTER TABLE [dbo].[needed_for] ADD PRIMARY KEY CLUSTERED
(
	[resource_main_resource_base_id] ASC,
	[resource_base_id] ASC
)WITH (PAD_INDEX = OFF, STATISTICS_NORECOMPUTE = OFF, SORT_IN_TEMPDB = OFF, IGNORE_DUP_KEY = OFF, ONLINE = OFF, ALLOW_ROW_LOCKS = ON, ALLOW_PAGE_LOCKS = ON) ON [PRIMARY]
GO
ALTER TABLE [dbo].[needed_for]  WITH CHECK ADD  CONSTRAINT [fk_needed_for_resource_base] FOREIGN KEY([resource_base_id])
REFERENCES [dbo].[resource_base] ([id])
GO
ALTER TABLE [dbo].[needed_for] CHECK CONSTRAINT [fk_needed_for_resource_base]
GO
ALTER TABLE [dbo].[needed_for]  WITH CHECK ADD  CONSTRAINT [fk_needed_for_resource_main] FOREIGN KEY([resource_main_resource_base_id])
REFERENCES [dbo].[resource_main] ([resource_id])
GO
ALTER TABLE [dbo].[needed_for] CHECK CONSTRAINT [fk_needed_for_resource_main]
GO

/*
************************************************
 Produced Froms
************************************************
*/

SET ANSI_NULLS ON
GO
SET QUOTED_IDENTIFIER ON
GO
CREATE TABLE [dbo].[produced_froms]
(
    [resource_id] [bigint] NULL,
    [amount] [float] NOT NULL
) ON [PRIMARY]
GO
ALTER TABLE [dbo].[produced_froms]  WITH CHECK ADD  CONSTRAINT [fk_resource_base_produced_from] FOREIGN KEY([resource_id])
REFERENCES [dbo].[resource_base] ([id])
GO
ALTER TABLE [dbo].[produced_froms] CHECK CONSTRAINT [fk_resource_base_produced_from]
GO
