import {
    Bot,
    Campaign,
    Delivery,
    DeliveryState,
    DeliveryTakeResult,
    User,
} from './types';

export interface BotDao {
    Create(bot: Bot): Promise<Bot>;
    Update(bot: Bot): Promise<Bot>;
    Get(botID: number): Promise<Bot>;
    List(): Promise<Bot[]>;
}

export interface CampaignDao {
    Create(campaign: Campaign): Promise<Campaign>;
    Update(campaign: Campaign): Promise<Campaign>;
    Get(botID: number, campaignID: number): Promise<Campaign>;
    List(): Promise<Campaign[]>;
}

export interface UserDao {
    Get(botID: number, telegramID: number): Promise<User>;
    Put(user: User): Promise<User>;
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
