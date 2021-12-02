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
	pieceType    string
	player       string
	currentCoord *Coordinate
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
	board = putPieceOnBoard(board, Coordinate{0, 3}, Piece{AMAZON_PIECE, PLAYER_1, nil})
	board = putPieceOnBoard(board, Coordinate{0, 6}, Piece{AMAZON_PIECE, PLAYER_1, nil})
	board = putPieceOnBoard(board, Coordinate{3, 0}, Piece{AMAZON_PIECE, PLAYER_1, nil})
	board = putPieceOnBoard(board, Coordinate{3, 9}, Piece{AMAZON_PIECE, PLAYER_1, nil})

	// Put player 2's initial pieces on the board
	board = putPieceOnBoard(board, Coordinate{6, 0}, Piece{AMAZON_PIECE, PLAYER_2, nil})
	board = putPieceOnBoard(board, Coordinate{6, 9}, Piece{AMAZON_PIECE, PLAYER_2, nil})
	board = putPieceOnBoard(board, Coordinate{9, 3}, Piece{AMAZON_PIECE, PLAYER_2, nil})
	board = putPieceOnBoard(board, Coordinate{9, 6}, Piece{AMAZON_PIECE, PLAYER_2, nil})

	return board
}

// export putPieceOnBoard
func putPieceOnBoard(board Board, coord Coordinate, piece Piece) Board {

	if piece.currentCoord != nil {
		board[piece.currentCoord.x][piece.currentCoord.y] = nil
		piece.currentCoord.x = coord.x
		piece.currentCoord.y = coord.y
	}

	board[coord.x][coord.y] = &piece

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
	printBoard(board)
}
