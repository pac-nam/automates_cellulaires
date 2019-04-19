/* ************************************************************************** */
/*                                                                            */
/*                                                        :::      ::::::::   */
/*   make_random.                                       :+:      :+:    :+:   */
/*                                                    +:+ +:+         +:+     */
/*   By: tbleuse <marvin@42.fr>                     +#+  +:+       +#+        */
/*                                                +#+#+#+#+#+   +#+           */
/*   Created: 2018/02/28 11:44:16 by tbleuse           #+#    #+#             */
/*   Updated: 2018/03/23 14:17:16 by tbleuse          ###   ########.fr       */
/*                                                                            */
/* ************************************************************************** */

#ifndef MAKE_RANDOM_H
# define MAKE_RANDOM_H

# include <fcntl.h>
# include <unistd.h>

int	ft_make_random(int nb, char *str);
int	ft_print_in_file(int *tab, int len, char *str);

# include "../libft/header/libft.h"

#endif
