import { Bot, Campaign, PaginatorResponse, User } from 'barker-api';
import { ReactElement } from 'react';
import { useParams } from 'react-router-dom';
import useSWR from 'swr';
import { fetcher } from './fetcher';
import { useQuery } from './useQuery';
interface TypeMap {
    'bot.List': Bot;
    'campaign.List': Campaign;
    'user.List': User;
}

type CreatedEntity<T extends keyof TypeMap> = T extends keyof TypeMap
    ? TypeMap[T]
    : never;

type KeyType = keyof TypeMap;
type EntityType = Bot | User | Campaign;

export interface ListLoaderProps<Entity extends EntityType> {
    render: (props: {
        items?: Entity[];
        error?: any;
        paging?: PaginatorResponse;
    }) => ReactElement;
}
export interface CreateListLoaderOptions {
    key: KeyType;
}

export function createListLoader<K extends KeyType>(key: K) {
    return ({ render }: ListLoaderProps<CreatedEntity<K>>) => {
        const { page = 1, size = 10 } = useQuery();
        const { botID } = useParams();

        // const { data: [items, paging] = [], error } = useSWR<
        //     [CreatedEntity<K>[], PaginatorResponse]
        // >([key, size, page, botID], fetcher);

        const { data: [items, paging] = [], error } = useSWR(
            [key, size, page, botID],
            fetcher
        );

        return render({
            items,
            error,
            paging,
        });
    };
}
