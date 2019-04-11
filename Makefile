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

fclean:
	$(RM) $(NAME)

re: fclean all

.PHONY: all fclean re
