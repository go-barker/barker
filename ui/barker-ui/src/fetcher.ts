import Axios from 'axios';
import {
    BarkerClient,
    Bot,
    Campaign,
    PaginatorResponse,
    User,
} from 'barker-api';

export const barker = new BarkerClient(Axios);

interface TypeMap {
    'bot.Get': Promise<Bot>;
    'campaign.Get': Promise<Campaign>;
    'bot.List': Promise<[Bot[], PaginatorResponse]>;
    'user.List': Promise<[User[], PaginatorResponse]>;
    'campaign.List': Promise<[Campaign[], PaginatorResponse]>;
}

type ReturnedType<K extends keyof TypeMap> = K extends keyof TypeMap
    ? TypeMap[K]
    : never;

export function fetcher<K extends keyof TypeMap, R extends ReturnedType<K>>(
    key: K,
    ...args: (string | number)[]
): R {
    switch (key) {
        case 'bot.List': {
            const [size, page] = args as number[];
            return <R>barker.bot.List({ Page: page, Size: size });
        }
        case 'user.List': {
            const [size, page, botID] = args as number[];
            return <R>barker.user.List(botID, { Page: page, Size: size });
        }
        case 'campaign.List': {
            const [size, page, botID] = args as number[];
            return <R>barker.campaign.List(botID, {
                Page: page,
                Size: size,
            });
        }
        case 'bot.Get': {
            const [botID] = args as number[];
            return <R>barker.bot.Get(botID);
        }
        case 'campaign.Get': {
            const [botID, campaignID] = args as number[];
            return <R>barker.campaign.Get(botID, campaignID);
        }
        default:
            throw new Error('Bad key');
    }
}
