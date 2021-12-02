package main

import "fmt"

const SIZE = 10

const GREEN = "\033[32m"
const RED = "\033[31m"
const DEFAULT = "\033[0m"

const AMAZON_PIECE = "o"
const ARROW_PIECE = "."

const PLAYER_1 = "P1"
const PLAYER_2 = "P2"

type Piece struct {
	pieceType string
	player    string
}

type Board [SIZE][SIZE]*Piece

type Coordinate struct {
	x uint8
	y uint8
}

// export initializeBoard
func initializeBoard() Board {
	board := Board{}

	// Put player 1's initial pieces on the board
	board = putPieceOnBoard(board, nil, Coordinate{0, 3}, Piece{AMAZON_PIECE, PLAYER_1})
	board = putPieceOnBoard(board, nil, Coordinate{0, 6}, Piece{AMAZON_PIECE, PLAYER_1})
	board = putPieceOnBoard(board, nil, Coordinate{3, 0}, Piece{AMAZON_PIECE, PLAYER_1})
	board = putPieceOnBoard(board, nil, Coordinate{3, 9}, Piece{AMAZON_PIECE, PLAYER_1})

	// Put player 2's initial pieces on the board
	board = putPieceOnBoard(board, nil, Coordinate{6, 0}, Piece{AMAZON_PIECE, PLAYER_2})
	board = putPieceOnBoard(board, nil, Coordinate{6, 9}, Piece{AMAZON_PIECE, PLAYER_2})
	board = putPieceOnBoard(board, nil, Coordinate{9, 3}, Piece{AMAZON_PIECE, PLAYER_2})
	board = putPieceOnBoard(board, nil, Coordinate{9, 6}, Piece{AMAZON_PIECE, PLAYER_2})

	return board
}

// export putPieceOnBoard
func putPieceOnBoard(board Board, oldCoord *Coordinate, newCoord Coordinate, piece Piece) Board {

	if oldCoord != nil {
		board[oldCoord.x][oldCoord.y] = nil
	}

	board[newCoord.x][newCoord.y] = &piece

	return board
}

// export printBoard
func printBoard(board Board) {

	for i := 0; i < SIZE; i++ {
		for j := 0; j < SIZE; j++ {
			piece := board[i][j]
			if piece == nil {
				fmt.Print(DEFAULT, "[ ]")
			} else {
				var color string
				if piece.player == PLAYER_1 {
					color = RED
				} else {
					color = GREEN
				}

				toPrint := "[" + piece.pieceType + "]"
				fmt.Print(color, toPrint)
			}
		}
		fmt.Println("")
	}
}

func main() {
	board := initializeBoard()

	fmt.Println("Initial board:")
	printBoard(board)

	fmt.Println("")
	board = putPieceOnBoard(board, &Coordinate{0, 3}, Coordinate{8, 3}, Piece{AMAZON_PIECE, PLAYER_1})
	fmt.Println("Updated board:")
	printBoard(board)
}
