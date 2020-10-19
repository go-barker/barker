import Axios from 'axios';
import {
    BarkerClient,
    Bot,
    Campaign,
    PaginatorResponse,
    User,
} from 'barker-api';

export const barker = new BarkerClient(Axios);

export function fetcher(
    key: 'bot.List',
    size: number,
    page: number
): Promise<[Bot[], PaginatorResponse]>;
export function fetcher(key: 'bot.Get', botID: number): Promise<Bot>;
export function fetcher(
    key: 'campaign.List',
    size: number,
    page: number,
    botID: number
): Promise<[Campaign[], PaginatorResponse]>;
export function fetcher(
    key: 'user.List',
    size: number,
    page: number,
    botID: number
): Promise<[User[], PaginatorResponse]>;
export function fetcher(
    key: 'bot.List' | 'bot.Get' | 'campaign.List' | 'user.List',
    ...args: (string | number)[]
) {
    switch (key) {
        case 'bot.List': {
            const [size, page] = args as number[];
            return barker.bot.List({ Page: page, Size: size });
        }
        case 'user.List': {
            const [size, page, botID] = args as number[];
            return barker.user.List(botID, { Page: page, Size: size });
        }
        case 'campaign.List': {
            const [size, page, botID] = args as number[];
            return barker.campaign.List(botID, {
                Page: page,
                Size: size,
            });
        }
        case 'bot.Get': {
            const [botID] = args as number[];
            return barker.bot.Get(botID);
        }
        default:
            throw new Error('Bad key');
    }
}
