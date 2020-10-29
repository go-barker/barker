import {
    Bot,
    Campaign,
    Delivery,
    DeliveryState,
    DeliveryTakeResult,
    User,
    PaginatorResponse,
    PaginatorRequest,
    CampaignAggregatedStatistics,
} from './types';

export interface BotDao {
    Create(bot: Bot): Promise<Bot>;
    Update(bot: Bot): Promise<Bot>;
    Get(botID: number): Promise<Bot>;
    List(pageRequest: PaginatorRequest): Promise<[Bot[], PaginatorResponse]>;
    RRTake(): Promise<Bot>;
}

export interface CampaignDao {
    Create(campaign: Campaign): Promise<Campaign>;
    Update(campaign: Campaign): Promise<Campaign>;
    Get(botID: number, campaignID: number): Promise<Campaign>;
    GetAggregatedStatistics(
        botID: number,
        campaignID: number
    ): Promise<CampaignAggregatedStatistics>;
    List(
        botID: number,
        pageRequest: PaginatorRequest
    ): Promise<[Campaign[], PaginatorResponse]>;
}

export interface UserDao {
    Get(botID: number, telegramID: number): Promise<User>;
    Put(user: User): Promise<User>;
    List(
        botID: number,
        pageRequest: PaginatorRequest
    ): Promise<[User[], PaginatorResponse]>;
}

export interface DeliveryDao {
    Take(
        botID: number,
        campaignID: number,
        telegramID: number
    ): Promise<DeliveryTakeResult>;
    SetState(delivery: Delivery, state: DeliveryState): Promise<void>;
    GetState(delivery: Delivery): Promise<DeliveryState>;
}
