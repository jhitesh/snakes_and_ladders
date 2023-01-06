# Definitions and Descriptions

1. A Snake's **HEAD** can lie in the range [3, SIZE-1] and **TAIL** can lie in the range [2, HEAD-1]
2. A Ladder's **BOTTOM** can lie in the range [2, SIZE-2] and **TOP** can lie in the range [BOTTOM+1, SIZE-1]

## NOTE
1. As soon as someone wins the game, the game is declared and closed
2. Set players from `JSON/players.json` file
3. Set Database configurations from `JSON/database.json` file

# Database Schema
## Players Table

```
CREATE TABLE players (
    ID INT PRIMARY KEY AUTO_INCREMENT,
    name VARCHAR(128) UNIQUE
);
```

## Board Table

```
CREATE TABLE boards (
    ID INT PRIMARY KEY AUTO_INCREMENT,
    size INT NOT NULL,
    num_of_jumpers INT NOT NULL,
    snakes JSON NOT NULL,
    ladders JSON NOT NULL
);
```

## Game Screen Table

```
CREATE TABLE game_screen (
    ID INT PRIMARY KEY AUTO_INCREMENT,
    num_of_dices INT,
    players_in_game JSON,
    board_id INT REFERENCES boards,
    winner_id INT REFERENCES players
);
```

## Game History Table

```
CREATE TABLE game_history (
    game_id INT REFERENCES game_screen,
    turn_no INT,
    player_id INT REFERENCES players,
    dice_roll INT,
    start INT,
    end INT,
    jumped BOOLEAN,
    PRIMARY KEY (game_id, turn_no)
);
```

- Data is stored in such a way that we can recreate a played game entirely