-- Create "games" table
CREATE TABLE `games` (`id` integer NOT NULL PRIMARY KEY AUTOINCREMENT, `show` integer NOT NULL, `air_date` datetime NOT NULL, `tape_date` datetime NOT NULL, `season_games` integer NULL, CONSTRAINT `games_seasons_games` FOREIGN KEY (`season_games`) REFERENCES `seasons` (`id`) ON DELETE SET NULL);
-- Create "seasons" table
CREATE TABLE `seasons` (`id` integer NOT NULL PRIMARY KEY AUTOINCREMENT, `number` integer NOT NULL, `start_date` datetime NOT NULL, `end_date` datetime NOT NULL);
