import { Bot, Campaign, PaginatorResponse, User } from 'barker-api';
import { ReactElement } from 'react';
import { useParams } from 'react-router-dom';
import useSWR from 'swr';
import { fetcher } from './fetcher';
import { useQuery } from './useQuery';

export interface ListLoaderProps<Entity> {
    render: (props: {
        items?: Entity[];
        error?: any;
        paging?: PaginatorResponse;
    }) => ReactElement;
}

export interface CreateListLoaderOptions {
    key: 'bot.List' | 'campaign.List' | 'user.List';
}

export function createListLoader<Entity extends Bot | Campaign | User>(
    options: CreateListLoaderOptions
) {
    return ({ render }: ListLoaderProps<Entity>) => {
        const { page = 1, size = 10 } = useQuery();
        const { botID } = useParams();

        const { data: [items, paging] = [], error } = useSWR<
            [Entity[], PaginatorResponse]
        >([options.key, size, page, botID], fetcher as any);
        return render({
            items,
            error,
            paging,
        });
    };
}
