import { AxiosInstance } from 'axios';
import { BotDao, UserDao, CampaignDao, DeliveryDao } from './dao';
import {
    Bot,
    PaginatorRequest,
    PaginatorResponse,
    User,
    Campaign,
    DeliveryTakeResult,
    Delivery,
    DeliveryState,
    CampaignAggregatedStatistics,
} from './types';
import U from 'url-template';

export class BotDaoImplAxios implements BotDao {
    constructor(private http: AxiosInstance) {}

    public async Create(bot: Bot): Promise<Bot> {
        const {
            data: { data },
        } = await this.http.post('/bot', bot);
        return data;
    }

    public async Update(bot: Bot): Promise<Bot> {
        const {
            data: { data },
        } = await this.http.put(
            U.parse('/bot/{BotID}').expand({ BotID: bot.ID }),
            bot
        );
        return data;
    }

    public async Get(botID: number): Promise<Bot> {
        const {
            data: { data },
        } = await this.http.get(
            U.parse('/bot/{BotID}').expand({ BotID: botID })
        );
        return data;
    }

    public async GetByToken(token: string): Promise<Bot> {
        const {
            data: { data },
        } = await this.http.get(
            U.parse('/bot-by-token/{token}').expand({ Token: token })
        );
        return data;
    }

    public async List(
        pageRequest: PaginatorRequest
    ): Promise<[Bot[], PaginatorResponse]> {
        const {
            data: { data, paging },
        } = await this.http.get('/bot', {
            params: pageRequest,
        });
        return [data, paging];
    }

    public async RRTake(): Promise<Bot> {
        const {
            data: { data },
        } = await this.http.post('/rr/bot');
        return data;
    }
}

export class UserDaoImplAxios implements UserDao {
    constructor(private http: AxiosInstance) {}

    public async Get(botID: number, telegramID: number): Promise<User> {
        const {
            data: { data },
        } = await this.http.get(
            U.parse('/bot/{botID}/user/{telegramID}').expand({
                botID,
                telegramID,
            })
        );
        return data;
    }

    public async Put(user: User): Promise<User> {
        const {
            data: { data },
        } = await this.http.put(
            U.parse('/bot/{botID}/user').expand({ botID: user.BotID })
        );
        return data;
    }

    public async List(
        botID: number,
        pageRequest: PaginatorRequest
    ): Promise<[User[], PaginatorResponse]> {
        const {
            data: { data, paging },
        } = await this.http.get(
            U.parse('/bot/{botID}/user').expand({ botID }),
            {
                params: pageRequest,
            }
        );
        return [data, paging];
    }
}

export class CampaignDaoImplAxios implements CampaignDao {
    constructor(private http: AxiosInstance) {}

    public async Create(campaign: Campaign): Promise<Campaign> {
        const {
            data: { data },
        } = await this.http.post(
            U.parse('/bot/{botID}/campaign').expand({ botID: campaign.BotID }),
            campaign
        );
        return data;
    }

    public async Update(campaign: Campaign): Promise<Campaign> {
        const {
            data: { data },
        } = await this.http.put(
            U.parse('/bot/{botID}/campaign/{campaignID}').expand({
                botID: campaign.BotID,
                campaignID: campaign.ID,
            }),
            campaign
        );
        return data;
    }

    public async Get(botID: number, campaignID: number): Promise<Campaign> {
        const {
            data: { data },
        } = await this.http.get(
            U.parse('/bot/{botID}/campaign/{campaignID}').expand({
                botID,
                campaignID,
            })
        );
        return data;
    }

    public async GetAggregatedStatistics(
        botID: number,
        campaignID: number
    ): Promise<CampaignAggregatedStatistics> {
        const {
            data: { data },
        } = await this.http.get(
            U.parse(
                '/bot/{botID}/campaign/{campaignID}/aggregatedStatistics'
            ).expand({
                botID,
                campaignID,
            })
        );
        return data;
    }

    public async List(
        botID: number,
        pageRequest: PaginatorRequest
    ): Promise<[Campaign[], PaginatorResponse]> {
        const {
            data: { data, paging },
        } = await this.http.get(
            U.parse('/bot/{botID}/campaign').expand({ botID }),
            {
                params: pageRequest,
            }
        );
        return [data, paging];
    }
}

export class DeliveryDaoImplAxios implements DeliveryDao {
    constructor(private http: AxiosInstance) {}

    public async Take(
        botID: number,
        campaignID: number,
        telegramID: number
    ): Promise<DeliveryTakeResult> {
        const url =
            campaignID === 0
                ? '/bot/{botID}/delivery'
                : '/bot/{botID}/campaign/{campaignID}/delivery';
        const {
            data: { data },
        } = await this.http.post(
            U.parse(url).expand({
                botID,
                campaignID,
            }),
            {},
            { params: { TelegramID: telegramID } }
        );
        return data;
    }

    public async SetState(
        delivery: Delivery,
        state: DeliveryState
    ): Promise<void> {
        await this.http.put(
            U.parse(
                '/bot/{botID}/campaign/{campaignID}/delivery/{telegramID}/state/{state}'
            ).expand({
                botID: delivery.BotID,
                campaignID: delivery.CampaignID,
                telegramID: delivery.TelegramID,
                state: DeliveryState[state],
            })
        );
    }

    public async GetState(delivery: Delivery): Promise<DeliveryState> {
        const {
            data: { data },
        } = await this.http.get(
            U.parse(
                '/bot/{botID}/campaign/{campaignID}/delivery/{telegramID}/state'
            ).expand({
                botID: delivery.BotID,
                campaignID: delivery.CampaignID,
                telegramID: delivery.TelegramID,
            })
        );
        return DeliveryState[data as keyof typeof DeliveryState];
    }
}
