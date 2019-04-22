##
## EPITECH PROJECT, 2018
## 205IQ
## File description:
## Makefile for 205IQ.
##

SRC		=	main.go			\
			params.go		\
			parseArgv.go	\
			utility.go

SRC		:=	$(addprefix src/, $(SRC))

NAME	=	205IQ

all: $(NAME)

$(NAME): $(SRC)
	go build -o $(NAME) $(SRC)

launch_tests:
	go test ./src/

fclean:
	$(RM) $(NAME)

re: fclean all

.PHONY: all launch_tests fclean re
