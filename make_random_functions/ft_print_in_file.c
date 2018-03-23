/* ************************************************************************** */
/*                                                                            */
/*                                                        :::      ::::::::   */
/*   ft_print_in_file.c                                 :+:      :+:    :+:   */
/*                                                    +:+ +:+         +:+     */
/*   By: tbleuse <marvin@42.fr>                     +#+  +:+       +#+        */
/*                                                +#+#+#+#+#+   +#+           */
/*   Created: 2018/02/28 11:37:56 by tbleuse           #+#    #+#             */
/*   Updated: 2018/03/23 17:54:15 by tbleuse          ###   ########.fr       */
/*                                                                            */
/* ************************************************************************** */

#include <sys/types.h>
#include <sys/stat.h>
#include "../header/make_random.h"

int	ft_print_in_file(int *tab, int len, char *str)
{
	char	*file;
	int	fd;

	file = ft_strjoin(str, "_random_numbers");
	if (!(fd = open(file, O_CREAT | O_WRONLY)))
		return (0);
	if (chmod(file, S_IRWXU | S_IRWXG | S_IROTH) < 0)
		return (0);
	while (--len >= 0)
	{
		ft_putnbr_fd(tab[len], fd);
		ft_putchar_fd('\n', fd);
	}
	close(fd);
	return (1);
}
