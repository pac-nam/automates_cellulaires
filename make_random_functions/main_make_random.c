/* ************************************************************************** */
/*                                                                            */
/*                                                        :::      ::::::::   */
/*   main_make_random.c                                 :+:      :+:    :+:   */
/*                                                    +:+ +:+         +:+     */
/*   By: tbleuse <marvin@42.fr>                     +#+  +:+       +#+        */
/*                                                +#+#+#+#+#+   +#+           */
/*   Created: 2018/02/28 11:37:56 by tbleuse           #+#    #+#             */
/*   Updated: 2018/03/23 17:54:15 by tbleuse          ###   ########.fr       */
/*                                                                            */
/* ************************************************************************** */

#include "../header/make_random.h"

int	ft_verif_av(char *str)
{
	int	i;

	i = -1;
	while (str[++i])
		if (!ft_isdigit(str[i]))
			return (0);
	return (1);
}

int	main(int ac, char **av)
{
	int	nb;

	if (ac != 2 || !ft_verif_av(av[1]))
	{
		ft_putstr("./make_random <integer number>\n");
		return (0);
	}
	nb = ft_atoi(av[1]);
	if (nb < 0)
		ft_putstr("impossible with negative quantity\n");
	if (nb == 0)
		ft_putstr("you are funny you\n");
	ft_make_random(nb, av[1]);
	return (0);
}
