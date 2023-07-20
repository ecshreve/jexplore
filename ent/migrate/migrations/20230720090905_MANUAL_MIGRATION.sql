pragma foreign_keys=off;

DELETE FROM `games` WHERE `season_id` = 39;
DELETE FROM `seasons` WHERE `id` = 39;

pragma foreign_keys=on;
