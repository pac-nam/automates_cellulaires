# **************************************************************************** #
#                                                                              #
#                                                         :::      ::::::::    #
#    Makefile                                           :+:      :+:    :+:    #
#                                                     +:+ +:+         +:+      #
#    By: tbleuse <marvin@42.fr>                     +#+  +:+       +#+         #
#                                                 +#+#+#+#+#+   +#+            #
#    Created: 2018/02/22 14:30:08 by tbleuse           #+#    #+#              #
#    Updated: 2018/03/23 17:55:11 by tbleuse          ###   ########.fr        #
#                                                                              #
# **************************************************************************** #

NAME = make_random

LIBFT = libft

CC = gcc

FLAGS = -Wall -Wextra -Werror

SRC_NAME =	main_make_random.c	\
		ft_make_random.c	\
		ft_print_in_file.c	\

SRC = $(addprefix make_random_functions/, $(SRC_NAME))

OBJ = $(SRC:.c=.o)

all : $(LIBFT) $(NAME)

$(LIBFT) :
	@make -C $(LIBFT)

$(NAME) : $(OBJ)
	@$(CC) $(FLAGS) $(OBJ) $(LIBFT)/libft.a -o $(NAME)
	@echo "$(NAME) have been compiled"

%.o : %.c
	@$(CC) $(FLAGS) -c $< -o $@

clean :
	@make clean -C $(LIBFT)
	@/bin/rm -f $(OBJ)
	@echo "$(NAME) objects have been deleted"

fclean : clean
	@make fclean -C $(LIBFT)
	@/bin/rm -f $(NAME)
	@/bin/rm -f *_random_numbers
	@echo "$(NAME) have been deleted"

re : fclean all

programs : all clean

.PHONY : all clean fclean re libft
