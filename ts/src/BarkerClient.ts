import { AxiosInstance } from 'axios';
import { BotDao, CampaignDao, DeliveryDao, UserDao } from './dao';
import {
    BotDaoImplAxios,
    CampaignDaoImplAxios,
    UserDaoImplAxios,
    DeliveryDaoImplAxios,
} from './dao_impl_axios';

export class BarkerClient {
    public readonly bot: BotDao;
    public readonly user: UserDao;
    public readonly campaign: CampaignDao;
    public readonly delivery: DeliveryDao;

    constructor(private http: AxiosInstance) {
        this.bot = new BotDaoImplAxios(http);
        this.campaign = new CampaignDaoImplAxios(http);
        this.user = new UserDaoImplAxios(http);
        this.delivery = new DeliveryDaoImplAxios(http);
    }
}

export * from './types';
