import { AxiosInstance } from 'axios';
import {
    BotDaoImplAxios,
    CampaignDaoImplAxios,
    DeliveryDaoImplAxios,
    UserDaoImplAxios,
} from './dao_impl_axios';

export class BarkerClient {
    public readonly bot: BotDaoImplAxios;
    public readonly user: UserDaoImplAxios;
    public readonly campaign: CampaignDaoImplAxios;
    public readonly delivery: DeliveryDaoImplAxios;

    constructor(private http: AxiosInstance) {
        this.bot = new BotDaoImplAxios(http);
        this.campaign = new CampaignDaoImplAxios(http);
        this.user = new UserDaoImplAxios(http);
        this.delivery = new DeliveryDaoImplAxios(http);
    }
}
