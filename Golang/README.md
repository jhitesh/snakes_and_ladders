# Definitions and Descritions

1. A Snake's **HEAD** can lie in the range [3, SIZE-1] and **TAIL** can lie in the range [2, HEAD-1]
2. Any Ladder's **BOTTOM** can lie in the range [2, SIZE-2] and **TOP** can lie in the range [BOTTOM+1, SIZE-1]

## NOTE
1. As soon as someone wins the game, the game is declared and closed
2. Set players from `JSON/players.json` file

# Database Schema
## Board Table

```
CREATE TABLE boards (
    ID int primary key,
    size int,
    num_of_jumpers int,
    snakes JSON,
    ladders JSON,
    screen_id int references game_screen
)
```

## Players Table

```
CREATE TABLE players (
    ID int primary key,
    name string unique
)
```

## Game Screen Table

```
CREATE TABLE game_screen (
    ID int primary key,
    num_of_dices int,
    players_in_game JSON,
    board_id int references boards,
    winner_id int references players
)
```

## Game History Table

```
CREATE TABLE game_history (
    game_id int references game_screen,
    turn_no int,
    player_id int references players,
    dice_roll int,
    start int,
    end int,
    jumped boolean,
    primary key (game_id, turn_no)
)
```