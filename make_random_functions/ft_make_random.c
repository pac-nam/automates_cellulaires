/* ************************************************************************** */
/*                                                                            */
/*                                                        :::      ::::::::   */
/*   ft_make_random.c                                   :+:      :+:    :+:   */
/*                                                    +:+ +:+         +:+     */
/*   By: tbleuse <marvin@42.fr>                     +#+  +:+       +#+        */
/*                                                +#+#+#+#+#+   +#+           */
/*   Created: 2018/02/28 11:37:56 by tbleuse           #+#    #+#             */
/*   Updated: 2018/03/23 17:54:15 by tbleuse          ###   ########.fr       */
/*                                                                            */
/* ************************************************************************** */

#include <stdio.h>
#include <time.h>
#include "../header/make_random.h"

static int	ft_some_neg(int nb, int len)
{
	if ((nb % 2 == 0 && len % 2 == 1) || (nb % 2 == 1 && len % 2 == 0))
		return (-nb);
	return (nb);
}

static int	ft_is_in_tab(int nb, int *tab, int len)
{
	while (--len >= 0)
		if (nb == tab[len])
			return (1);
	return (0);
}

int		ft_make_random(int nb, char *str)
{
	int	tab[nb];
	int	rand_nb;
	int	len;

	srand(time(NULL));
	len = 0;
	tab[0] = ft_some_neg(rand(), len + 1);
	rand_nb = ft_some_neg(rand(), len);
	while (++len < nb)
	{
		while (ft_is_in_tab(rand_nb, tab, len))
			rand_nb = ft_some_neg(rand(), len);
		tab[len] = rand_nb;
	}
	if (!ft_print_in_file(tab, nb, str))
		ft_putstr("Error\n");
	return (0);
}
